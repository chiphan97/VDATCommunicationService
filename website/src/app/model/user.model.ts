import * as _ from 'lodash';

export class User {
  userId: string;
  firstName: string;
  lastName: string;
  fullName: string;
  avatar: string;
  role: string;
  username: string;

  constructor(userId: string, firstName: string, lastName: string, fullName: string, avatar: string, role: string, username: string) {
    this.userId = userId;
    this.firstName = firstName;
    this.lastName = lastName;
    this.fullName = fullName;
    this.avatar = avatar;
    this.role = role;
    this.username = username;
  }

  public static fromJson(data: any): User {
    return new User(
      _.get(data, 'id', ''),
      _.get(data, 'first', ''),
      _.get(data, 'last', ''),
      _.get(data, 'fullName', ''),
      _.get(data, 'avatar', ''),
      _.get(data, 'role', ''),
      _.get(data, 'userName', ''),
    );
  }
}
