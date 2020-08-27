import {User} from './user.model';

export class Message {
  id: number;
  groupId: number;
  sender: User;
  content: string;
  createdAt: Date;
  children?: Array<Message>;
}
