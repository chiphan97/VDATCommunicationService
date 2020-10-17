import {EventEmitter, Injectable} from '@angular/core';
import {environment} from '../../../environments/environment';
import {KeycloakService} from '../auth/keycloak.service';
import {MessagePayload} from '../../model/payload/message.payload';
import {MessageDto} from '../../model/messageDto.model';
import {WsEvent} from '../../const/ws.event';
import * as _ from 'lodash';
import { Message } from '@angular/compiler/src/i18n/i18n_ast';

@Injectable({
  providedIn: 'root'
})
export class ChatService {
  private socket: WebSocket;
  private listener: EventEmitter<MessageDto> = new EventEmitter();
  private readonly WS_ENDPOINT = `${environment.service.wsUrl}/${environment.service.endpoint.message}`;

  constructor(private keycloakService: KeycloakService) { 
  }

  public initWebSocket(socketId: string) {
    const accessToken = this.keycloakService.accessToken;
    this.socket = new WebSocket(`${this.WS_ENDPOINT}/${socketId}?token=${accessToken}`);
 
    //this.socket = new WebSocket(`${this.WS_ENDPOINT}/${groupId.toString()}?token=${accessToken}`);
    //this.socket = new WebSocket(`${this.WS_ENDPOINT}/${groupId.toString()}`);
    this.socket.onopen = event => {
      console.log('opened ws');
    };
    this.socket.onclose = event => {
      console.log('closed ws');
    };
    this.socket.onmessage = event => {
      console.log('on message');
      const message = JSON.parse(event.data);
      const messageDto : MessageDto = {
        payload: message,
        senderId: message.Client
      }
        
      this.listener.emit(messageDto);
      // this.listener.emit({
      //   id: -1,
      //   content: _.get(message.data, 'body', ''),
      //   createdAt: new Date(),
      //   sender: null,
      //   groupId: _.get(message.data, 'groupId', -1),
      //   children: null
      // });
    };

    this.socket.onerror = event => {
      console.error(event);
    }
  }

  public close() {
    this.socket.close();
  }

  public getEventListener(): EventEmitter<any> {
    return this.listener;
  }

  public sendMessage(message: string, groupId: number, socketId: string): void {
    const payload: MessagePayload = {
      data: {
        'groupId' : groupId,
        'body' : message,
        "socketId": socketId
      },
      type: WsEvent.SEND_TEXT,
      groupId: groupId
    };

    this.socket.send(JSON.stringify(payload));
  }
}