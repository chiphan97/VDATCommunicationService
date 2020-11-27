import {User} from './user.model';
import {NzUploadFile} from 'ng-zorro-antd/upload';
import {Group} from './group.model';
import * as _ from 'lodash';

export class GenericMessage {
  id: number;
  group: Group;
  sender: User;
  messageType: MessageType;
  content?: any;
  parentID: number;
  createdAt: Date;
  children?: Array<GenericMessage>;

  constructor(id: number, group: Group, sender: User, parentID: number, createdAt: Date, children: Array<GenericMessage>) {
    this.id = id;
    this.group = group;
    this.sender = sender;
    this.parentID = parentID;
    this.createdAt = createdAt;
    this.children = children;
  }
}

export class FileMessage extends GenericMessage {
  constructor(id: number, group: Group, sender: User, content: NzUploadFile[], parentID: number, createdAt: Date, children: Array<GenericMessage>) {
    super(id, group, sender, parentID, createdAt, children);
    this.messageType = 'FILE_MESSAGE';
    this.content = content;
  }
}

export class TextMessage extends GenericMessage {
  constructor(id: number, group: Group, sender: User, content: string, parentID: number, createdAt: Date, children: Array<GenericMessage>) {
    super(id, group, sender, parentID, createdAt, children);
    this.messageType = 'TEXT_MESSAGE';
    this.content = content;
  }

  public static fromJson(obj: any): TextMessage {
    return new TextMessage(
      _.get(obj, 'id', -1),
      _.get(obj, 'groupId', -1),
      _.get(obj, 'sender', null),
      _.get(obj, 'content', ''),
      _.get(obj, 'parentID', -1),
      _.get(obj, 'createdAt', new Date()),
      _.get(obj, 'children', [])
    );
  }
}

export type MessageType = 'TEXT_MESSAGE' | 'FILE_MESSAGE';