import {EventEmitter, Injectable} from '@angular/core';
import {environment} from '../../../environments/environment';
import {KeycloakService} from '../auth/keycloak.service';
import {MessagePayload} from '../../model/payload/message.payload';
import {MessageDto} from '../../model/messageDto.model';
import {WsEvent} from '../../const/ws.event';
import * as _ from 'lodash';


@Injectable({
  providedIn: 'root'
})
export class ChatService {
  private socket: WebSocket;
  private messageListener: EventEmitter<MessageDto> = new EventEmitter();
  private chatHistoryListener: EventEmitter<Array<MessageDto>> = new EventEmitter();
  private readonly WS_ENDPOINT = `${environment.service.wsUrl}/${environment.service.endpoint.message}`;

  constructor(private keycloakService: KeycloakService) { 
  }

  public initWebSocket(socketId: string) {
    const accessToken = this.keycloakService.accessToken;
    this.socket = new WebSocket(`${this.WS_ENDPOINT}/${socketId}?token=${accessToken}`);
 
    this.socket.onopen = event => {
      console.log('opened ws');
    };
    this.socket.onclose = event => {
      console.log('closed ws');
    };
    this.socket.onmessage = event => {
      console.log('on message, event:');
      console.log(event);
      const mssgData = JSON.parse(event.data.trim());

      if (mssgData.Client){
        const messageDto: MessageDto  = {
          payload: mssgData,
          senderId: mssgData.Client
        }   
        this.messageListener.emit(messageDto);
      }
      else {
        console.log(mssgData.split(/\r\n|\r|\n/));
        const messageDtos: Array<MessageDto> = mssgData.split(/\r\n|\r|\n/).map(message => {         
          return {
            payload: message,
            senderId: message.Client
          }   
        });
        this.chatHistoryListener.emit(messageDtos);
      }         
    };

    this.socket.onerror = event => {
      console.error(event);
    }
  }

  public close() {
    this.socket.close();
  }

  public getChatEventListener(): EventEmitter<any> {
    return this.messageListener;
  }

  public getChatHistoryListener(): EventEmitter<any> {
    return this.chatHistoryListener;
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

  public sendGroupChatHistoryRequest(groupId: number, socketId: string): void {
    const payload: MessagePayload = {
      data: {
        'groupId' : groupId,
        'body' : '',
        "socketId": socketId
      },
      type: WsEvent.SUBCRIBE_GROUP,
    }
    this.socket.send(JSON.stringify(payload));
  }
}