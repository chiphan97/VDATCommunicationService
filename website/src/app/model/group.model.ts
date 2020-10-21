import {Message} from './message.model';
import {User} from './user.model';
import {GroupType} from '../const/group-type.const';
import * as _ from 'lodash';

export class Group {
  id: number;
  nameGroup: string;
  type: GroupType;
  isPrivate: boolean;
  thumbnail: string;
  description: string;
  owner: string;
  lastMessage: Message;

  historyMessages: Array<Message>;
  members: Array<User> = [];
  isOwner: boolean;
  isGroup: boolean;

  constructor(id: number, nameGroup: string, type: GroupType,
              isPrivate: boolean, thumbnail: string, description: string, owner: string, lastMessage: Message) {
    this.id = id;
    this.nameGroup = nameGroup;
    this.type = type;
    this.isPrivate = isPrivate;
    this.thumbnail = thumbnail;
    this.lastMessage = lastMessage;
    this.owner = owner;
    this.description = description;
    this.isGroup = this.type === GroupType.MANY;
  }

  public static fromJson(data: any) {
    return new Group(
      _.get(data, 'id', -1),
      _.get(data, 'nameGroup', ''),
      _.get(data, 'type', ''),
      _.get(data, 'private', true),
      _.get(data, 'thumbnail', ''),
      _.get(data, 'description', ''),
      _.get(data, 'owner', ''),
      null
    );
  }
}
