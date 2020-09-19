import {AfterViewChecked, Component, ElementRef, Input, OnChanges, OnInit, SimpleChanges, ViewChild} from '@angular/core';
import {Group} from '../../../model/group.model';
import {Message} from '../../../model/message.model';
import {formatDistance} from 'date-fns';
import {StorageService} from '../../../service/common/storage.service';
import {User} from '../../../model/user.model';
import {ChatService} from '../../../service/ws/chat.service';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import * as _ from 'lodash';

@Component({
  selector: 'app-chat-content',
  templateUrl: './chat-content.component.html',
  styleUrls: ['./chat-content.component.sass']
})
export class ChatContentComponent implements OnInit, AfterViewChecked, OnChanges {

  @Input() groupSelected: Group;
  @ViewChild('message-content') private myScrollContainer: ElementRef;

  public currentUser: User;
  public messages: Array<Message>;
  public formGroup: FormGroup;

  constructor(private storageService: StorageService,
              private chatService: ChatService) {
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

  ngAfterViewChecked() {
    this.scrollToBottom();
  }

  public onSubmit(): void {
    if (this.formGroup.valid) {
      const rawValue = this.formGroup.getRawValue();
      const message = _.get(rawValue, 'message', '');
      this.chatService.sendMessage(message);
      this.formGroup.patchValue({message: ''});
    }
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
