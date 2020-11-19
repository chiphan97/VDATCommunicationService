import {
  Component,
  ElementRef, EventEmitter, 
  Input,
  OnInit, Output, OnChanges,
  ViewChild,
  SimpleChanges
} from '@angular/core';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {Group} from '../../../model/group.model';
import {User} from '../../../model/user.model';
import { NzUploadFile, NzUploadModule } from 'ng-zorro-antd/upload';
import {GenericMessage, TextMessage, FileMessage} from '../../../model/generic-message.model';
import {ChatService} from '../../../service/ws/chat.service';
import * as _ from 'lodash';

@Component({
  selector: 'messenger-reply-thread-right',
  templateUrl: './messenger-reply-thread-right.component.html',
  styleUrls: ['./messenger-reply-thread-right.component.sass']
})
export class MessengerReplyThreadRightComponent implements OnInit {
  @Input() groupSelected: Group;
  @Input() currentUser: User;
  @Input() memberOfGroup: Array<User>;

  @Input() parentMessage: GenericMessage;
  @Output() parentMessageChange = new EventEmitter<GenericMessage>();

  //@Output() loadMore = new EventEmitter();

  @ViewChild('messagesContainer') private messagesContainer: ElementRef;
  @ViewChild('textInput') private inputElement: ElementRef;

  public formGroup: FormGroup;
  public loading: boolean;
  public isScrollHeight = true;

  public patientUnknown: User = new User('45', 'Anonymous', 'Patient', '', null, 'patient', 'username', null, null, null);
  public submitting = false;

  public fileList: NzUploadFile[] = [];
  public previewImage: string | undefined = '';
  public previewVisible = false;
  public messageToReply: GenericMessage;

  public openUploadFile: boolean;
  public openReplyToMessage: boolean;

  constructor(private chatService: ChatService) { }

  ngOnInit(): void {
    this.messageToReply = this.parentMessage;
    this.formGroup = this.createFormGroup();
    this.focusInputField();

    if (!this.parentMessageChange) {
      console.log('EEM is undefined');
    }
  }

  ngOnChanges( changes: SimpleChanges): void {
    if (changes.parentMessage.currentValue !== changes.parentMessage.previousValue) {
      this.messageToReply = this.parentMessage;
    }
  }

  private createFormGroup(): FormGroup {
    return new FormGroup({
      message: new FormControl('', [Validators.required])
    });
  }

  public onSubmit(): void {
    this.onSubmitText();

    this.clearReplyToMessage();
  }

  private onSubmitText(): void {
    if (this.formGroup.valid) {
      const rawValue = this.formGroup.getRawValue();
      const message = _.get(rawValue, 'message', '');
      this.chatService.replyToMessage(message, this.groupSelected.id, this.messageToReply.id);    
      this.formGroup.patchValue({message: ''});
    }  
  }

  attachIconClicked = () => {
    this.openUploadFile = !this.openUploadFile;
  }

  handlePreviewFile = async (file: NzUploadFile) => {
    if (!file.url && !file.preview) {
      file.preview = await getBase64(file.originFileObj!);
    }
    this.previewImage = file.url || file.preview;
    this.previewVisible = true;
  }

  closeReplyThread = (): void => {
    //this.parentMessage = null;
    this.openReplyToMessage = false;
    this.parentMessageChange.emit(null);
  }

  clearReplyToMessage(): void{
    this.openReplyToMessage = false;
  }

  onReplyToMessage(event) {
    if (!this.openReplyToMessage) {
      this.openReplyToMessage = true;
    }
    this.focusInputField();
    this.messageToReply = event;
  }

  focusInputField(): void {
    setTimeout(()=>{
      this.inputElement.nativeElement.focus();
    },0);  
  }
}

function getBase64(file: File): Promise<string | ArrayBuffer | null> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
}
