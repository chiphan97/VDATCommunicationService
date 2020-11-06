import {Component, Input, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import {User} from '../../../model/user.model';
import {Message} from '../../../model/message.model';
import {formatDistance} from "date-fns";

@Component({
  selector: 'app-messenger-message',
  templateUrl: './messenger-message.component.html',
  styleUrls: ['./messenger-message.component.sass']
})
export class MessengerMessageComponent implements OnInit {

  @Input() currentUser: User;
  @Input() message: Message;

  public numLikes = 0;
  public numDislikes = 0;

  constructor() {
  }

  ngOnInit(): void {
  }

  public isOwner = (): boolean => this.currentUser && this.message
    && this.currentUser.userId === this.message.sender.userId

  public formatDistanceTime = (date: Date): string => formatDistance(date, new Date());

  onLike(): void {
    this.numLikes++;
  }

  onDislike(): void {
    this.numDislikes++;
  }

}
