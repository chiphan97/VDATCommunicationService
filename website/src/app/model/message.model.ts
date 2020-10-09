import {User} from './user.model';

export class Message {
  id: number;
  groupId: number;
  sender: User;
  content: string;
  createdAt: Date;
  children?: Array<Message>;

  public constructor(messageForm: any){
    messageForm.value.id,
    messageForm.value.groupId,
    messageForm.value.sender,
    messageForm.value.content,
    messageForm.value.createdAt,
    messageForm.value.children
  }
}
