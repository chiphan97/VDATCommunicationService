import * as _ from 'lodash';

export class UserOnline {
  hostName: string;
  socketId: string;
  userId: string;
  fullName: string;
  logAt: Date;

  constructor(hostName: string, socketId: string, userId: string, fullName: string, logAt: Date) {
    this.hostName = hostName;
    this.socketId = socketId;
    this.userId = userId;
    this.fullName = fullName;
    this.logAt = logAt;
  }

  public static fromJson(data: any): UserOnline {
    return new UserOnline(
      _.get(data, 'host_name', ''),
      _.get(data, 'socket_id', ''),
      _.get(data, 'id', ''),
      _.get(data, 'username', ''),
      new Date(_.get(data, 'log_at', '')));
  }
}
