import { Injectable, EventEmitter } from '@angular/core';
import {environment} from '../../../environments/environment';
import {KeycloakService} from '../auth/keycloak.service';
import {Message} from '../../model/message.model'
@Injectable({
  providedIn: 'root'
})
export class WebSocketService {

  webSocket: WebSocket;
  chatMessages: Message[] = [];
  private listener: EventEmitter<Message> = new EventEmitter();
  private readonly WS_ENDPOINT = `${environment.service.wsUrl}${environment.service.endpoint.chat}`;

  constructor(private keycloakService: KeycloakService) { }

  public openWebsocket(groupId: number){
    const accessToken = this.keycloakService.accessToken;
    this.webSocket = new WebSocket(`${this.WS_ENDPOINT}/${groupId.toString()}?token=${accessToken}`);

    this.webSocket.onopen = (event) => {
      console.log('open', event);
    }
    this.webSocket.onmessage = (event) => {
      const chatMessage = JSON.parse(event.data);
      this.chatMessages.push(chatMessage);
      this.listener.emit(chatMessage);
    }

    this.webSocket.onclose = (event) => {
      console.log('Close ', event);
    }
  }

  public sendMessage(chatMessage: Message){
    this.webSocket.send(JSON.stringify(chatMessage));
  }

  public closeWebsocket(){
    this.webSocket.close();
  }

  public getMessageHistory(): Message[]{
    return [];
  }
}
