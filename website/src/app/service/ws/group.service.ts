import {EventEmitter, Injectable} from '@angular/core';
import {environment} from '../../../environments/environment';
import {Group} from '../../model/group.model';
import {GroupPayload} from '../../model/payload/group.payload';
import {MessagePayload} from '../../model/payload/message.payload';
import {WsEvent} from '../../const/ws.event';
import {group} from '@angular/animations';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class GroupService {

  private readonly WS_ENDPOINT = `${environment.service.wsUrl}${environment.service.endpoint.groups}`;
  private socket: WebSocket = null;
  private listener: EventEmitter<MessagePayload> = new EventEmitter();

  constructor() {
  }

  public connect(accessToken: string): Observable<boolean> {
    return new Observable<boolean>(observer => {
      this.socket = new WebSocket(`${this.WS_ENDPOINT}?token=${accessToken}`);
      this.socket.onopen = () => observer.next(true);
      this.socket.onclose = () => observer.next(false);
      this.socket.onmessage = event => {
        const message = JSON.parse(event.data);
        this.listener.emit(message);
      };
    });
  }

  public send(messagePayload: MessagePayload) {
    while (!this.socket.readyState) {
    }
    this.socket.send(JSON.stringify(messagePayload));
    console.log(messagePayload);
  }

  public close() {
    this.socket.close();
  }

  public getEventListener() {
    return this.listener;
  }

  public createGroup(groupPayload: GroupPayload): void {
    const message: MessagePayload = {
      type: WsEvent.CREATE_GROUP,
      data: groupPayload
    };

    this.send(message);
  }

  public getAllGroup(): Observable<Array<Group>> {
    return new Observable<Array<Group>>(observer => {
      this.send({
        type: WsEvent.LIST_ALL_GROUP,
        data: null
      });

      this.listener
        .subscribe(message => {
          if (message.type === WsEvent.LIST_ALL_GROUP) {
            observer.next(message.data);
          }
        });
    });
  }
}
