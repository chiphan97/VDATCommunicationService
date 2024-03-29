import {EventEmitter, Injectable} from '@angular/core';
import {environment} from '../../environments/environment';
import * as _ from 'lodash';
import {UserOnline} from '../model/user-online.model';

@Injectable({
  providedIn: 'root'
})
export class UserOnlineService {
  private socket: WebSocket;
  private listener: EventEmitter<any> = new EventEmitter();
  private readonly usersOnline: Array<UserOnline>;
  private readonly WS_ENDPOINT = environment.service.wsUrl;

  public constructor() {
    this.usersOnline = new Array<UserOnline>();
  }

  public initWebSocket() {
    this.socket = new WebSocket(`${this.WS_ENDPOINT}/user-online?token=${localStorage.getItem('TOKEN')}`);
    this.socket.onopen = event => {
      this.listener.emit({type: 'open', data: event});
    };
    this.socket.onclose = event => {
      this.listener.complete();
      this.socket.close();
    };
    this.socket.onmessage = event => {
      const message = JSON.parse(event.data);
      const body = _.get(message, 'body', {});
      const users = _.map(body, item => UserOnline.fromJson(item));

      const newUsers = _.differenceBy(users, this.usersOnline, 'userId');

      if (newUsers.length > 0) {
        this.usersOnline.push(...newUsers);
      }

      this.listener.emit({type: 'online', data: users});
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

  public getUsersOnline(): Array<UserOnline> {
    return this.usersOnline;
  }
}
