import { Component, OnInit } from '@angular/core';
import {formatDistance} from 'date-fns';
import {Message} from '../../model/message.model';
import {NzDrawerService} from 'ng-zorro-antd';
import {MessengerDrawerComponent} from '../messenger-drawer/messenger-drawer.component';

@Component({
  selector: 'app-messenger-content',
  templateUrl: './messenger-content.component.html',
  styleUrls: ['./messenger-content.component.sass']
})
export class MessengerContentComponent implements OnInit {

  currentUser = 'Me';
  messages: Array<Message>;

  constructor() {
    this.messages = this.fakeDate();
  }

  ngOnInit(): void {
  }

  formatDistanceTime(date: Date = new Date()): string {
    return formatDistance(date, new Date());
  }

  fakeDate(): Array<Message> {
    const message: Message = {
      id: 1,
      groupId: 1,
      sender: {
        subject: '1',
        fullName: 'Nguyễn Chí Cường',
        avatarUrl: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png'
      },
      createdAt: new Date(),
      content: 'Hello world !!!'
    };

    const messages = new Array<Message>();
    for (let i = 0; i < 20; i++) {
      messages.push(message);
    }

    return messages;
  }
}
