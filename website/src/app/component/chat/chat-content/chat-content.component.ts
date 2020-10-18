import { AfterViewChecked, Component, ElementRef, Input, OnChanges, OnInit, SimpleChanges, ViewChild } from '@angular/core';
import { Group } from '../../../model/group.model';
import { Message } from '../../../model/message.model';
import { MessageDto } from '../../../model/messageDto.model'
import { formatDistance } from 'date-fns';
import { StorageService } from '../../../service/common/storage.service';
import { User } from '../../../model/user.model';
import { ChatService } from '../../../service/ws/chat.service';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import * as _ from 'lodash';
import { GroupService } from 'src/app/service/collector/group.service';
import { WsEvent } from 'src/app/const/ws.event';

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
  public patientUnknown: User = new User('45', 'Anonymous', 'Patient', '', null, 'patient', 'username', null, null, null);
  public historyMessages: Message[] = [];

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
      if (changes.groupSelected && this.groupSelected) {
        this.chatService.initWebSocket(this.currentUser.socketId);

        //get list of group members
        this.groupService.getAllMemberOfGroup(this.groupSelected.id).subscribe((users: Array<User>) => {
          this.groupSelected.members = users;
          //get history chat
          this.chatService.sendGroupChatHistoryRequest(this.groupSelected.id, this.currentUser.socketId);
        })

        this.chatService.getChatEventListener()
          .subscribe((messageDto: MessageDto) => {
            if (messageDto.payload.type.trim() === WsEvent.SEND_TEXT) {
              const message: Message = this.getMessage(this.groupSelected, messageDto);

              if (message.sender.userId !== this.currentUser.userId) {
                this.historyMessages = [...this.historyMessages, message];
              }
            }
            else {
              console.log('not yet supported type ');
            }
          });

        //todo: seperate 2 methods, return boolean 
        this.chatService.getChatHistoryListener().subscribe((messageDtos: Array<MessageDto>) => {
          const pastMessages: Array<Message> = messageDtos.map((messageDto: MessageDto) => {
            return this.getMessage(this.groupSelected, messageDto);
          })
          this.historyMessages = pastMessages;
          this.scrollToBottom();
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
      this.chatService.sendMessage(message, this.groupSelected.id, this.currentUser.socketId);
      this.formGroup.patchValue({ message: '' });

      this.mockupUISendMessage(message, this.groupSelected.id);
    }
  }

  private mockupUISendMessage(message: string, groupId: number): void {
    this.submitting = true;

    setTimeout(() => {
      this.historyMessages = [
        ...this.historyMessages,
        {
          id: 1,
          groupId: groupId,
          sender: this.currentUser,
          content: message,
          createdAt: new Date(),
          children: [],
        }
      ].map(e => {
        return {
          ...e,
        };
      });
      this.scrollToBottom();
    }, 200);
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

  private getMessage(group: Group, messageDto: MessageDto): Message {
    const message: Message = {
      id: -1,
      content: _.get(messageDto.payload.data, 'body', ''),
      createdAt: new Date(),
      sender: this.groupService.getUserById(group, messageDto.senderId),
      groupId: _.get(messageDto.payload.data, 'groupId', -1),
      children: null
    }
    if (!message.sender) {
      message.sender = this.patientUnknown;
    }
    return message;
  }
}
