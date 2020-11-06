import {
  AfterContentChecked,
  AfterViewChecked, AfterViewInit, ChangeDetectorRef,
  Component, DoCheck,
  ElementRef, EventEmitter, HostListener,
  Input,
  OnChanges,
  OnInit, Output,
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
import * as _ from 'lodash';
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

  private DEFAULT_SCROLL_OFFSET_TOP = 150;

  public formGroup: FormGroup;
  public loading: boolean;

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
    const offsetTop = parseInt(event.target.scrollTop, 0);

    if (offsetTop <= this.DEFAULT_SCROLL_OFFSET_TOP) {
      this.loadMore.emit();
    }
  }

  public onSubmit(): void {
    if (this.formGroup.valid) {
      const rawValue = this.formGroup.getRawValue();
      const message = _.get(rawValue, 'message', '');
      this.chatService.sendMessage(message, this.groupSelected.id);
      this.formGroup.patchValue({message: ''});
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
