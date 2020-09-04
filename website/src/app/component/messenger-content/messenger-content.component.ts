import {AfterViewChecked, Component, ElementRef, Input, OnInit, ViewChild} from '@angular/core';
import {formatDistance} from 'date-fns';
import {Message} from '../../model/message.model';
import {Group} from '../../model/group.model';

@Component({
  selector: 'app-messenger-content',
  templateUrl: './messenger-content.component.html',
  styleUrls: ['./messenger-content.component.sass']
})
export class MessengerContentComponent implements OnInit, AfterViewChecked {

  @Input() groupSelected: Group;

  @ViewChild('message-content') private myScrollContainer: ElementRef;

  currentUser = 'Me';
  messages: Array<Message>;

  constructor() {
    this.messages = this.fakeDate();
  }

  ngOnInit(): void {
    this.scrollToBottom();
  }

  ngAfterViewChecked() {
    this.scrollToBottom();
  }

  formatDistanceTime(date: Date = new Date()): string {
    return formatDistance(date, new Date());
  }

  scrollToBottom(): void {
    try {
      this.myScrollContainer.nativeElement.scrollTop = this.myScrollContainer.nativeElement.scrollHeight;
    } catch (err) { }
  }

  fakeDate(): Array<Message> {
    const message: Message = {
      id: 1,
      groupId: 1,
      sender: {
        avatar: '',
        userId: '1',
        fullName: 'Nguyễn Chí Cường',
        lastName: '',
        firstName: ''
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
