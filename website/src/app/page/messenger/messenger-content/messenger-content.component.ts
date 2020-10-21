import {
  AfterContentChecked,
  AfterViewChecked,
  Component,
  ElementRef,
  Input,
  OnChanges,
  OnInit,
  SimpleChanges,
  ViewChild
} from '@angular/core';
import {Group} from '../../../model/group.model';
import {Message} from '../../../model/message.model';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {User} from '../../../model/user.model';
import {StorageService} from '../../../service/common/storage.service';
import {ChatService} from '../../../service/ws/chat.service';
import {GroupService} from '../../../service/collector/group.service';
import {MessageDto} from '../../../model/messageDto.model';
import {WsEvent} from '../../../const/ws.event';
import * as _ from 'lodash';
import {formatDistance} from 'date-fns';
import {NzMessageService} from 'ng-zorro-antd';

@Component({
  selector: 'app-messenger-content',
  templateUrl: './messenger-content.component.html',
  styleUrls: ['./messenger-content.component.sass']
})
export class MessengerContentComponent implements OnInit, OnChanges {

  @Input() groupSelected: Group;
  @Input() isMember: boolean;
  @Input() currentUser: User;
  @Input() memberOfGroup: Array<User>;

  @ViewChild('message-content') private myScrollContainer: ElementRef;

  public messages: Array<Message>;
  public formGroup: FormGroup;
  public patientUnknown: User = new User('45', 'Anonymous', 'Patient', '', null, 'patient', 'username', null, null, null);
  public historyMessages: Message[] = [];

  submitting = false;

  /* end of mock data for users*/
  constructor(private storageService: StorageService,
              private chatService: ChatService,
              private messageService: NzMessageService,
              private groupService: GroupService) {
    this.formGroup = this.createFormGroup();
    this.messages = new Array<Message>();
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes) {
      // close old web socket
      this.chatService.close();

      // init new websocket
      if (this.groupSelected && this.currentUser) {
        this.chatService.initWebSocket(this.currentUser.socketId)
          .subscribe(ready => {
            if (!ready) {
              this.messageService.error('Không thể kết nối đến server !');
            } else {
              // get history chat
              this.chatService.sendGroupChatHistoryRequest(this.groupSelected.id, this.currentUser.socketId);

              this.chatService.getChatEventListener()
                .subscribe((messageDto: MessageDto) => {
                  if (messageDto.payload.type.trim() === WsEvent.SEND_TEXT) {
                    const message: Message = this.getMessage(this.groupSelected, messageDto);

                    if (message.sender.userId !== this.currentUser.userId) {
                      this.historyMessages = [...this.historyMessages, message];
                    }
                  } else {
                    console.log('not yet supported type ');
                  }
                });

              // todo: seperate 2 methods, return boolean
              this.chatService.getChatHistoryListener().subscribe((messageDtos: Array<MessageDto>) => {
                const pastMessages: Array<Message> = messageDtos.map((messageDto: MessageDto) => {
                  console.log(messageDto);
                  return this.getMessage(this.groupSelected, messageDto);
                });
                this.historyMessages = pastMessages;
              });
            }
          });
      }
    }
  }

  ngOnInit(): void {
  }

  public onSubmit(): void {
    if (this.formGroup.valid) {
      const rawValue = this.formGroup.getRawValue();
      const message = _.get(rawValue, 'message', '');
      this.chatService.sendMessage(message, this.groupSelected.id, this.currentUser.socketId);
      this.formGroup.patchValue({message: ''});

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
          groupId,
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
    }, 200);
  }

  formatDistanceTime(date: Date = new Date()): string {
    return formatDistance(date, new Date());
  }

  private createFormGroup(): FormGroup {
    return new FormGroup({
      message: new FormControl('', [Validators.required])
    });
  }

  private getMessage(group: Group, messageDto: MessageDto): Message {
    if (!this.memberOfGroup) {
      return null;
    }

    const sender: User = this.memberOfGroup.find(member => member.userId === messageDto.senderId);

    const message: Message = {
      id: -1,
      content: _.get(messageDto.payload.data, 'body', ''),
      createdAt: new Date(),
      sender,
      groupId: _.get(messageDto.payload.data, 'groupId', -1),
      children: null
    };

    if (!message.sender) {
      message.sender = this.patientUnknown;
    }

    return message;
  }
}
