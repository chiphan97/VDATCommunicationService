import {EventEmitter, Injectable} from '@angular/core';

import * as socketIo from 'socket.io-client';
import {Observable} from 'rxjs';
import {Event} from '../const/event';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class SocketService {
  private socket: WebSocket;
  private listener: EventEmitter<any> = new EventEmitter();

  public constructor() {
  }

  public initWebSocket(accessToken: string) {
    this.socket = new WebSocket(`${environment.wsUrl}/test?token=${accessToken}`);
    this.socket.onopen = event => {
      this.listener.emit({type: 'open', data: event});
    };
    this.socket.onclose = event => {
      this.listener.emit({type: 'close', data: event});
    };
    this.socket.onmessage = event => {
      const message = JSON.parse(event.data);
      this.listener.emit({type: 'message', data: message});
    };
  }

  public send(data: string) {
    this.socket.send(data);
  }

  public close() {
    this.socket.close();
  }

  public getEventListener() {
    return this.listener;
  }
}
