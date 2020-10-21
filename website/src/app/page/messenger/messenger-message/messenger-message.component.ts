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
  @Input() message: Message;

  public user: User;
  public numLikes = 0;
  public numDislikes = 0;

  constructor() {
  }

  ngOnInit(): void {
    this.user = this.currentUser;
  }

  public isOwner = (user: User): boolean => this.currentUser && this.currentUser.userId === user.userId;

  onLike(): void {
    this.numLikes++;
  }

  onDislike(): void {
    this.numDislikes++;
  }

}
