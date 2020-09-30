import {EventEmitter, Injectable} from '@angular/core';
import {environment} from '../../../environments/environment';
import {KeycloakService} from '../auth/keycloak.service';
import {MessagePayload} from '../../model/payload/message.payload';
import {Message} from '../../model/message.model';
import * as _ from 'lodash';

@Injectable({
  providedIn: 'root'
})
export class ChatService {
  private socket: WebSocket;
  private listener: EventEmitter<Message> = new EventEmitter();
  private readonly WS_ENDPOINT = `${environment.service.wsUrl}${environment.service.endpoint.chat}`;

  constructor(private keycloakService: KeycloakService) { }

  public initWebSocket(groupId: number) {
    const accessToken = this.keycloakService.accessToken;

    this.socket = new WebSocket(`${this.WS_ENDPOINT}/${groupId.toString()}?token=${accessToken}`);
    this.socket.onopen = event => {
      console.log('opened ws');
    };
    this.socket.onclose = event => {
      console.log('closed ws');
    };
    this.socket.onmessage = event => {
      const message = JSON.parse(event.data);
      this.listener.emit({
        id: -1,
        content: _.get(message, 'body', ''),
        createdAt: new Date(),
        sender: null,
        groupId: _.get(message, 'group_id', -1),
        children: null
      });
    };
  }

  public close() {
    this.socket.close();
  }

  public getEventListener(): EventEmitter<any> {
    return this.listener;
  }

  public sendMessage(message: string): void {
    const payload: MessagePayload = {
      data: message,
      type: null,
      groupId: null
    };

    this.socket.send(JSON.stringify(payload));
  }
}
