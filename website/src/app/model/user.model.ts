import * as _ from 'lodash';
import {UserStatus} from '../const/user-status.enum';
import * as CryptoJS from 'crypto-js';

export class User {
  userId: string;
  firstName: string;
  lastName: string;
  fullName: string;
  avatar: string;
  role: string;
  username: string;
  status: UserStatus;
  hostName: string;
  socketId: string;

  color: string;
  isOnline: boolean;

  constructor(userId: string, firstName: string, lastName: string, fullName: string,
              avatar: string, role: string, username: string, status: UserStatus = UserStatus.OFFLINE,
              hostName: string = '', socketId: string = '') {
    this.userId = userId;
    this.firstName = firstName;
    this.lastName = lastName;
    this.fullName = fullName;
    this.role = role;
    this.username = username;
    this.status = status;
    this.hostName = hostName;
    this.socketId = socketId;
    this.isOnline = status === UserStatus.ONLINE;
    this.avatar = avatar;

    // if (!!avatar) {
    //   this.avatar = avatar;
    // } else if (!!userId) {
    //   const email = `${userId.trim()}@vdatlab.com`;
    //   const hash = CryptoJS.MD5(email.toLowerCase());
    //   this.avatar = `https://www.gravatar.com/${hash}`;
    // } else {
    //   this.avatar = '';
    // }
  }

  public static fromJson(data: any): User {
    return new User(
      _.get(data, 'id', '').trim(),
      _.get(data, 'first', '').trim(),
      _.get(data, 'last', '').trim(),
      _.get(data, 'fullName', '').trim(),
      _.get(data, 'avatar', '').trim(),
      _.get(data, 'role', '').trim(),
      _.get(data, 'userName', '').trim(),
      _.get(data, 'status', UserStatus.OFFLINE),
      _.get(data, 'hostName', '').trim(),
      _.get(data, 'socketId', '').trim(),
    );
  }
}
