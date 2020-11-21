import {
  AfterContentChecked,
  ChangeDetectorRef,
  Component,
  ElementRef,
  EventEmitter,
  HostListener,
  Input,
  OnInit,
  Output,
  ViewChild
} from '@angular/core';
import {Group} from '../../../model/group.model';
import {Message} from '../../../model/message.model';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {User} from '../../../model/user.model';
import {StorageService} from '../../../service/common/storage.service';
import {ChatService} from '../../../service/ws/chat.service';
import {GroupService} from '../../../service/collector/group.service';
import * as _ from 'lodash';
import * as Sentry from '@sentry/angular';
import {NzMessageService} from 'ng-zorro-antd';

@Component({
  selector: 'app-messenger-content',
  templateUrl: './messenger-content.component.html',
  styleUrls: ['./messenger-content.component.sass']
})
export class MessengerContentComponent implements OnInit, AfterContentChecked {

  @Input() groupSelected: Group;
  @Input() isMember: boolean;
  @Input() currentUser: User;
  @Input() memberOfGroup: Array<User>;
  @Input() messages: Array<Message>;

  @Output() loadMore = new EventEmitter();

  @ViewChild('messagesContainer') private messagesContainer: ElementRef;

  private DEFAULT_SCROLL_OFFSET_TOP = 200;

  public formGroup: FormGroup;
  public loading: boolean;
  public isScrollHeight = true;
  private oldScrollTop: number;

  public patientUnknown: User = new User('45', 'Anonymous', 'Patient', '', null, 'patient', 'username', null, null, null);
  public submitting = false;

  constructor(private storageService: StorageService,
              private chatService: ChatService,
              private changeDetectorRef: ChangeDetectorRef,
              private messageService: NzMessageService,
              private groupService: GroupService) {
    this.formGroup = this.createFormGroup();
  }

  ngOnInit(): void {
  }

  ngAfterContentChecked() {
    this.changeDetectorRef.detectChanges();
  }

  @HostListener('scroll', ['$event'])
  public onMessageContainerScroll(event: any) {
    const scrollTop = parseInt(event.target.scrollTop, 0);

    console.log(this.oldScrollTop, scrollTop);

    if (!!!this.oldScrollTop) {
      console.log('init');
      this.oldScrollTop = scrollTop;
    } else if (this.oldScrollTop <= scrollTop) {
      console.log('oldScrollTop <= scrollTop');
      this.isScrollHeight = true;
    } else {
      console.log('oldScrollTop > scrollTop');
      this.oldScrollTop = scrollTop;
      this.isScrollHeight = false;
    }

    if (scrollTop <= this.DEFAULT_SCROLL_OFFSET_TOP) {
      const lastMessage: Message = this.messages[0];
      this.chatService.getMessagesHistory(this.groupSelected.id, lastMessage.id)
        .subscribe(() => {
          console.log('đang load thêm tin nhắn');
        });
      this.loadMore.emit();
    }

    console.log('isScrollHeight: ', this.isScrollHeight);
  }

  public onSubmit(): void {
    if (this.formGroup.valid) {
      const rawValue = this.formGroup.getRawValue();
      const message = _.get(rawValue, 'message', '');
      this.chatService.sendMessage(message, this.groupSelected.id);
      this.formGroup.patchValue({message: ''});
      this.scrollToBottom();
    }
  }

  private scrollToBottom() {
    try {
      this.messagesContainer.nativeElement.scrollTop = this.messagesContainer.nativeElement.scrollHeight;
    } catch (err) {
      Sentry.captureException(err);
    }
  }

  private mockupUISendMessage(message: string, groupId: number): void {
    this.submitting = true;
  }

  private createFormGroup(): FormGroup {
    return new FormGroup({
      message: new FormControl('', [Validators.required])
    });
  }
}
