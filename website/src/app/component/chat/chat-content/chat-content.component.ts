import {AfterViewChecked, Component, ElementRef, Input, OnChanges, OnInit, OnDestroy, SimpleChanges, ViewChild} from '@angular/core';
import {Group} from '../../../model/group.model';
import {Message} from '../../../model/message.model';
import {StorageService} from '../../../service/common/storage.service';
import {WebSocketService} from '../../../service/web-socket/web-socket.service'
import {User} from '../../../model/user.model';
import {ChatService} from '../../../service/ws/chat.service';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {NgForm} from '@angular/forms';
import { addDays, formatDistance } from 'date-fns';
import * as _ from 'lodash';
import { ChildrenOutletContexts } from '@angular/router';
import { getLocaleDateTimeFormat } from '@angular/common';

@Component({
  selector: 'app-chat-content',
  templateUrl: './chat-content.component.html',
  styleUrls: ['./chat-content.component.sass']
})
export class ChatContentComponent implements OnInit, AfterViewChecked, OnChanges, OnDestroy {

  @Input() groupSelected: Group;
  @ViewChild('message-content') private myScrollContainer: ElementRef;

  public currentUser: User;
  public messages: Array<Message>;
  public formGroup: FormGroup;
  public data = [{
    author: 'Han Solo',
    avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
    content:
      'We supply a series of design principles, practical patterns and high quality design resources' +
      '(Sketch and Axure), to help people create their product prototypes beautifully and efficiently.',
    children: [
      {
        author: 'Han Solo',
        avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        content:
          'We supply a series of design principles, practical patterns and high quality design resources' +
          '(Sketch and Axure), to help people create their product prototypes beautifully and efficiently.',
      }
    ]
  }];
  public historyMessages: Message[] = [{
    id: 1,
    groupId: 44,
    sender: this.currentUser,
    content:
      'We supply a series of design principles, practical patterns and high quality design resources' +
      '(Sketch and Axure), to help people create their product prototypes beautifully and efficiently.',
    createdAt: new Date(),
    children:[]
  }];

  submitting = false;
  public user = {
    author: 'Han Solo',
    avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png'
  };
  public inputValue = '';

  constructor(private storageService: StorageService,
              private chatService: ChatService,
              public websocketService: WebSocketService) {
    this.currentUser = this.storageService.userInfo;
    this.formGroup = this.createFormGroup();
    this.messages = new Array<Message>();
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes) {
      if (changes.groupSelected && this.groupSelected) {
        this.chatService.initWebSocket(this.groupSelected.id);

        this.chatService.getEventListener()
          .subscribe(message => this.messages.push(message));
      }
    }
  }

  ngOnInit(): void {
    this.scrollToBottom();
  }

  ngOnDestroy(): void{
    this.websocketService.closeWebsocket();
  }

  ngAfterViewChecked() {
    this.scrollToBottom();
  }

  sendMessage(sendForm: NgForm){
    console.log(sendForm.value);
    const toSendMessage = new Message(sendForm);
    this.websocketService.sendMessage(toSendMessage);
    sendForm.controls.message.reset();
  }

  public onSubmit(): void {
    if (this.formGroup.valid) {
      const rawValue = this.formGroup.getRawValue();
      const message = _.get(rawValue, 'message', '');
      this.chatService.sendMessage(message);
      this.formGroup.patchValue({message: ''});

      this.mockupUISendMessage(message);
    }
  }

  private mockupUISendMessage(message): void {
    this.inputValue = '';
    this.submitting = true;
    const content = message;
    
    setTimeout(() => {
      this.submitting = false;
      this.data = [
        ...this.data,
        {
          ...this.user,
          content,
          displayTime: formatDistance(new Date(), new Date()),
          children: []
        }
      ].map(e => {
        return {
          ...e,
        };
      });
    }, 800);
  }

  formatDistanceTime(date: Date = new Date()): string {
    return formatDistance(date, new Date());
  }

  scrollToBottom(): void {
    try {
      this.myScrollContainer.nativeElement.scrollTop = this.myScrollContainer.nativeElement.scrollHeight;
    } catch (err) { }
  }

  private createFormGroup(): FormGroup {
    return new FormGroup({
      message: new FormControl('', [Validators.required])
    });
  }
}
