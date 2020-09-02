import {Message} from './message.model';
import {GroupType} from '../const/group-type.const';

export class Group {
  nameGroup: string;
  type: GroupType;
  private: boolean;
  users: Array<string>;
  thumbnail: string;
  lastMessage: Message;
}
