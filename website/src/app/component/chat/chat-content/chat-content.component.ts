import {AfterViewChecked, Component, ElementRef, Input, OnChanges, OnInit, SimpleChanges, ViewChild} from '@angular/core';
import {Group} from '../../../model/group.model';
import {Message} from '../../../model/message.model';
import {MessageDto} from '../../../model/messageDto.model'
import {formatDistance} from 'date-fns';
import {StorageService} from '../../../service/common/storage.service';
import {User} from '../../../model/user.model';
import {ChatService} from '../../../service/ws/chat.service';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import * as _ from 'lodash';
import { GroupService } from 'src/app/service/collector/group.service';

@Component({
  selector: 'app-chat-content',
  templateUrl: './chat-content.component.html',
  styleUrls: ['./chat-content.component.sass']
})
export class ChatContentComponent implements OnInit, AfterViewChecked, OnChanges {

  @Input() groupSelected: Group;
  @ViewChild('message-content') private myScrollContainer: ElementRef;

  public messages: Array<Message>;
  public formGroup: FormGroup;
  public currentUser: User;

  /*mock data for users*/
  public patient1User: User = new User ('34', 'Hoang', 'Hong', 'Hoang Thi Hong', null, 'patient', 'username', null, null, null);
  public patientUnknown: User = new User ('45', 'Anonymous', 'Patient', 'Hoang Thi Hong', null, 'patient', 'username', null, null, null);
  public data = [{
    author: 'Chi Phan',
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
  public historyMessages: Message[] = [
    {
      id: 3,
      groupId: 22,
      sender: this.patient1User,
      content:
      'I have an enquiry about my current health condition...',
      createdAt: new Date(),
      children: [{
        id: 3,
        groupId: 22,
        sender: this.patient1User,
        content:
        'also a sub comment',
        createdAt: new Date(),
        children: []
      }]
    },
    {
      id: 3,
      groupId: 22,
      sender: this.patient1User,
      content:
      'I do not feel well',
      createdAt: new Date(),
      children: []
    }
  ];

  public user = {
    author: 'Han Solo',
    avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png'
  }; 

  submitting = false;
  /* end of mock data for users*/
  constructor(private storageService: StorageService,
              private chatService: ChatService,
              private groupService: GroupService) {
    this.currentUser = this.storageService.userInfo;
    this.formGroup = this.createFormGroup();
    this.messages = new Array<Message>();
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes) {
      if (changes.groupSelected && this.groupSelected ) {
        this.chatService.initWebSocket(this.currentUser.socketId);

        //get list of group members
        this.groupService.getAllMemberOfGroup(this.groupSelected.id).subscribe((users: Array<User>) => {
          this.groupSelected.members = users;
        })

        //get list of group messages
        this.chatService.getEventListener()
          .subscribe((messageDto: MessageDto) => {
            const message: Message = {
              id: -1,
              content: _.get(messageDto.payload.data, 'body', ''),
              createdAt: new Date(),
              sender: this.getUserById(this.groupSelected, messageDto.senderId),
              groupId: _.get(messageDto.payload.data, 'groupId', -1),
              children: null
            }

            if (!message.sender){
              message.sender = this.patientUnknown;
            }

            if (message.sender.userId !== this.currentUser.userId){
              this.historyMessages = [...this.historyMessages, message];
            }
          });
      }
    }
  }

  ngOnInit(): void {
    this.scrollToBottom();
  }

  ngAfterViewChecked() {
    this.scrollToBottom();
  }

  public onSubmit(): void {
    if (this.formGroup.valid) {
      const rawValue = this.formGroup.getRawValue();
      const message = _.get(rawValue, 'message', '');
      //todo: parse to Message
      this.chatService.sendMessage(message, this.groupSelected.id, this.currentUser.socketId);
      this.formGroup.patchValue({message: ''});

      this.mockupUISendMessage(message, this.groupSelected.id);
    }    
  }

  private mockupUISendMessage(message: string, groupId: number): void {
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
    }, 500);

    setTimeout(() => {
    this.historyMessages = [
      ...this.historyMessages,
      {
        id: 1,
        groupId: groupId,
        sender: this.currentUser,
        content: content,
        createdAt: new Date(),
        children: [],
      }
    ].map(e => {
      return {
        ...e,
      };
    });
  }, 500);

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

  public getFirstname(user: User): string {
    return user.firstName;
  }

  private getUserById(group: Group, senderId: string): User {
    return group.members.find(user => user.userId === senderId);
  }
}
