import {AfterViewChecked, Component, ElementRef, Input, OnInit, ViewChild} from '@angular/core';
import {Group} from '../../../model/group.model';
import {Message} from '../../../model/message.model';
import {formatDistance} from 'date-fns';

@Component({
  selector: 'app-chat-content',
  templateUrl: './chat-content.component.html',
  styleUrls: ['./chat-content.component.sass']
})
export class ChatContentComponent implements OnInit, AfterViewChecked {

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
