import {EventEmitter, Injectable} from '@angular/core';
import {environment} from '../../../environments/environment';
import {KeycloakService} from '../auth/keycloak.service';
import {MessagePayload} from '../../model/payload/message.payload';
import {MessageDto} from '../../model/messageDto.model';
import {WsEvent} from '../../const/ws.event';
import * as _ from 'lodash';
import {Observable} from 'rxjs';
import {StorageService} from '../common/storage.service';
import {User} from '../../model/user.model';


@Injectable({
  providedIn: 'root'
})
export class ChatService {
  private readonly WS_ENDPOINT = `${environment.service.wsUrl}${environment.service.endpoint.message}`;

  private socket: WebSocket;
  private currentUser: User;
  private messageListener: EventEmitter<MessageDto> = new EventEmitter();
  private chatHistoryListener: EventEmitter<Array<MessageDto>> = new EventEmitter();

  constructor(private keycloakService: KeycloakService,
              private storageService: StorageService) {
    this.currentUser = this.storageService.userInfo;

    if (!!!this.currentUser.socketId) {
      throw new Error('Cannot connect to websocket');
    }
  }

  /**
   * Open Websocket
   */
  public open(): Observable<boolean> {
    return new Observable<boolean>(observer => {
      const accessToken = this.keycloakService.accessToken;
      const socketId = this.currentUser.socketId;
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
    });
  }

  /**
   * Close Websocket
   */
  public close() {
    if (this.socket && !(this.socket.CLOSED || this.socket.CLOSING)) {
      this.socket.close();
    }
  }

  /**
   * Listenning
   */
  public listener(): Observable<MessageDto> {
    return new Observable<MessageDto>(observer => {
      this.getSocket()
        .subscribe(socket => {
          if (socket) {
            socket.onmessage = event => {
              const rawData = event.data;
              const data = '[' + rawData.replace(/\n/g, ',') + ']';

              const messages = JSON.parse(data);
              messages.forEach(message => {
                const messageDto = MessageDto.fromJson(message);
                observer.next(messageDto);
              });
            };
          }
        });
    });
  }

  /**
   * Get socket instance
   */
  public getSocket(): Observable<WebSocket> {
    return new Observable<WebSocket>(observer => {
      if (!!this.socket) {
        observer.next(this.socket);
        observer.complete();
      } else {
        this.open()
          .subscribe(ready => {
            if (ready) {
              observer.next(this.socket);
              observer.complete();
            } else {
              observer.error('Cannot connect to websocket');
              observer.complete();
            }
          });
      }
    });
  }

  /**
   * Subscribe group
   * @param groupId group id
   */
  public getChatHistory(groupId: number, lastMessageId?: number): Observable<boolean> {
    const socketId = this.currentUser.socketId;

    return new Observable<boolean>(observer => {
      this.getSocket()
        .subscribe(socket => {
          if (socket) {
            const message: MessagePayload = {
              data: {
                groupId,
                body: '',
                socketId
              },
              type: WsEvent.SUBCRIBE_GROUP,
            };
            socket.send(JSON.stringify(message));

            observer.next(true);
            observer.complete();
          } else {
            observer.next(false);
            observer.complete();
          }
        });
    });
  }

  /**
   * Load old messages
   * @param groupId group id
   * @param lastMessageId last message id
   */
  public getMessagesHistory(groupId: number, lastMessageId: number = -1): Observable<boolean> {
    return new Observable<boolean>(observer => {
      this.getSocket()
        .subscribe(socket => {
          if (socket) {
            const message: MessagePayload = {
              data: {
                groupId,
                socketId: this.currentUser.socketId,
                idContinueOldMess: lastMessageId
              },
              type: WsEvent.LOAD_OLD_MESSAGE,
            };
            socket.send(JSON.stringify(message));

            observer.next(true);
            observer.complete();
          } else {
            observer.next(false);
            observer.complete();
          }
        });
    });
  }

  public getChatEventListener(): EventEmitter<any> {
    return this.messageListener;
  }

  public getChatHistoryListener(): EventEmitter<any> {
    return this.chatHistoryListener;
  }

  public sendMessage(message: string, groupId: number): void {
    const socketId = this.currentUser.socketId;

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
}
