import {User} from './user.model';
import * as _ from 'lodash';
import {Group} from './group.model';

export class Message {
  id: number;
  group: Group;
  sender: User;
  content: string;
  createdAt: Date;
  children?: Array<Message>;

  constructor(id: number, group: Group, sender: User, content: string,
              createdAt: Date, children: Array<Message> = new Array<Message>()) {
    this.id = id;
    this.group = group;
    this.sender = sender;
    this.content = content;
    this.createdAt = createdAt;
    this.children = children;
  }

  public static fromJson(obj: any): Message {
    return new Message(
      _.get(obj, 'id', -1),
      _.get(obj, 'groupId', -1),
      _.get(obj, 'sender', null),
      _.get(obj, 'content', ''),
      _.get(obj, 'createdAt', new Date()),
      _.get(obj, 'children', [])
    );
  }
}
