import * as _ from 'lodash';

export class User {
  userId: string;
  firstName: string;
  lastName: string;
  fullName: string;
  avatar: string;

  constructor(userId: string, firstName: string, lastName: string, fullName: string, avatar: string) {
    this.userId = userId;
    this.firstName = firstName;
    this.lastName = lastName;
    this.fullName = fullName;
    this.avatar = avatar;
  }

  public static fromJson(data: any): User {
    return new User(
      _.get(data, 'userId', ''),
      _.get(data, 'firstName', ''),
      _.get(data, 'lastName', ''),
      _.get(data, 'fullName', ''),
      _.get(data, 'avatar', ''));
  }
}
