import { Component, OnInit, Input } from '@angular/core';
import { User } from '../../../../model/user.model';
import { Message } from '../../../../model/message.model';

@Component({
  selector: 'app-message',
  templateUrl: './message.component.html',
  styleUrls: ['./message.component.sass']
})
export class MessageComponent implements OnInit {

  @Input() currentUser: User;
  @Input() messageInput: Message;

  public user: User;
  public likes = 0;
  public dislikes = 0;

  public myContext = {message: this.messageInput};

  constructor() { }

  ngOnInit(): void {
    this.user = this.currentUser;
    console.log('this messageInput');
    console.log(this.messageInput);
    console.log('this user '+ this.user.userId);
    console.log('sender: '+ this.messageInput.sender.userId);
    console.log(this.messageInput.sender.userId.trim() == this.user.userId.trim());
  }

  public getFirstname(user: User): string {
    return user.firstName;
  }

  like(): void {
    this.likes = 1;
    this.dislikes = 0;
  }

  dislike(): void {
    this.likes = 0;
    this.dislikes = 1;
  }

}
