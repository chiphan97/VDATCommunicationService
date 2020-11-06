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
import * as _ from 'lodash';
import {formatDistance} from 'date-fns';
import {NzMessageService} from 'ng-zorro-antd';

@Component({
  selector: 'app-messenger-content',
  templateUrl: './messenger-content.component.html',
  styleUrls: ['./messenger-content.component.sass']
})
export class MessengerContentComponent implements OnInit {

  @Input() groupSelected: Group;
  @Input() isMember: boolean;
  @Input() currentUser: User;
  @Input() memberOfGroup: Array<User>;
  @Input() messages: Array<Message>;

  @ViewChild('message-content') private myScrollContainer: ElementRef;

  public formGroup: FormGroup;

  public patientUnknown: User = new User('45', 'Anonymous', 'Patient', '', null, 'patient', 'username', null, null, null);
  public submitting = false;

  constructor(private storageService: StorageService,
              private chatService: ChatService,
              private messageService: NzMessageService,
              private groupService: GroupService) {
    this.formGroup = this.createFormGroup();
  }

  ngOnInit(): void {
  }

  public onSubmit(): void {
    if (this.formGroup.valid) {
      const rawValue = this.formGroup.getRawValue();
      const message = _.get(rawValue, 'message', '');
      this.chatService.sendMessage(message, this.groupSelected.id);
      this.formGroup.patchValue({message: ''});

      this.mockupUISendMessage(message, this.groupSelected.id);
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
