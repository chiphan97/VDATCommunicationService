import * as _ from 'lodash';

export class UserOnline {
  hostName: string;
  socketId: string;
  userId: string;
  fullName: string;

  constructor(hostName: string, socketId: string, userId: string, fullName: string) {
    this.hostName = hostName;
    this.socketId = socketId;
    this.userId = userId;
    this.fullName = fullName;
  }

  public static fromJson(data: any): UserOnline {
    const obj = JSON.parse(data);
    console.log(obj);
    return new UserOnline('', '', '', '');
  }
}
