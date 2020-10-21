import {EventEmitter, Injectable} from '@angular/core';
import {environment} from '../../../environments/environment';
import {KeycloakService} from '../auth/keycloak.service';
import {MessagePayload} from '../../model/payload/message.payload';
import {MessageDto} from '../../model/messageDto.model';
import {WsEvent} from '../../const/ws.event';
import * as _ from 'lodash';
import {Observable} from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class ChatService {
  private socket: WebSocket;
  private messageListener: EventEmitter<MessageDto> = new EventEmitter();
  private chatHistoryListener: EventEmitter<Array<MessageDto>> = new EventEmitter();
  private readonly WS_ENDPOINT = `${environment.service.wsUrl}${environment.service.endpoint.message}`;

  constructor(private keycloakService: KeycloakService) {
  }

  public initWebSocket(socketId: string): Observable<boolean> {
    return new Observable<boolean>(observer => {
      const accessToken = this.keycloakService.accessToken;
      this.socket = new WebSocket(`${this.WS_ENDPOINT}/${socketId}?token=${accessToken}`);

      this.socket.onopen = () => observer.next(true);

      this.socket.onclose = () => {
        observer.next(false);
        observer.complete();
      };

      this.socket.onerror = () => {
        observer.next(false);
        observer.complete();
      };

      this.socket.onmessage = event => {
        const msgData = JSON.parse(event.data.trim());

        if (msgData.Client) {
          const messageDto: MessageDto = {
            payload: msgData,
            senderId: msgData.Client
          };
          this.messageListener.emit(messageDto);
        } else {
          const messageDtos: Array<MessageDto> = msgData.historys.map(message => {
            return {
              payload: message,
              senderId: message.data.Sender
            };
          });
          this.chatHistoryListener.emit(messageDtos);
        }
      };
    });
  }

  // disconnect websocket
  public close() {
    if (this.socket && !(this.socket.CLOSED || this.socket.CLOSING)) {
      this.socket.close();
    }
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
        groupId,
        body: message,
        socketId
      },
      type: WsEvent.SEND_TEXT,
      groupId
    };

    this.socket.send(JSON.stringify(payload));
  }

  public sendGroupChatHistoryRequest(groupId: number, socketId: string): void {
    const payload: MessagePayload = {
      data: {
        groupId,
        body: '',
        socketId
      },
      type: WsEvent.SUBCRIBE_GROUP,
    };
    this.socket.send(JSON.stringify(payload));
  }
}
