import {Component, Input, OnInit} from '@angular/core';
import {User} from '../../../model/user.model';
import {Message} from '../../../model/message.model';

@Component({
  selector: 'app-messenger-message',
  templateUrl: './messenger-message.component.html',
  styleUrls: ['./messenger-message.component.sass']
})
export class MessengerMessageComponent implements OnInit {

  @Input() currentUser: User;
  @Input() messageInput: Message;

  public user: User;
  public likes = 0;
  public dislikes = 0;

  public myContext = {message: this.messageInput};

  constructor() { }

  ngOnInit(): void {
    this.user = this.currentUser;
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
