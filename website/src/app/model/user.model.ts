import * as _ from 'lodash';

export class User {
  userId: string;
  firstName: string;
  lastName: string;
  fullName: string;

  constructor(userId: string, firstName: string, lastName: string, fullName: string) {
    this.userId = userId;
    this.firstName = firstName;
    this.lastName = lastName;
    this.fullName = fullName;
  }

  public static fromJson(data: any): User {
    return new User(
      _.get(data, 'id', ''),
      _.get(data, 'first', ''),
      _.get(data, 'last', ''),
      _.get(data, 'username', ''));
  }
}
