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
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {User} from '../../../model/user.model';
import {GenericMessage, TextMessage, FileMessage} from '../../../model/generic-message.model';
import {StorageService} from '../../../service/common/storage.service';
import {ChatService} from '../../../service/ws/chat.service';
import {GroupService} from '../../../service/collector/group.service';
import { MinioService} from '../../../service/upload/minio.service';
import * as _ from 'lodash';
import {NzMessageService} from 'ng-zorro-antd';
import { NzUploadFile, NzUploadModule } from 'ng-zorro-antd/upload';

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
  @Input() messages: Array<GenericMessage>;

  @Output() loadMore = new EventEmitter();

  @ViewChild('messagesContainer') private messagesContainer: ElementRef;
  @ViewChild('textInput') private inputElement: ElementRef;

  private DEFAULT_SCROLL_OFFSET_TOP = 200;

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

  constructor(private storageService: StorageService,
              private chatService: ChatService,
              private changeDetectorRef: ChangeDetectorRef,
              private messageService: NzMessageService,
              private groupService: GroupService,
              private minioService: MinioService) {
    this.formGroup = this.createFormGroup();
  }

  ngOnInit(): void {
    this.focusInputField();
  }

  ngAfterContentChecked() {
    this.changeDetectorRef.detectChanges();
  }

  @HostListener('scroll', ['$event'])
  public onMessageContainerScroll(event: any) {
    this.isScrollHeight = false;
    const offsetTop = parseInt(event.target.scrollTop, 0);

    if (offsetTop <= this.DEFAULT_SCROLL_OFFSET_TOP) {
      const lastMessage: Message = this.messages[0];
      this.chatService.getMessagesHistory(this.groupSelected.id, lastMessage.id)
        .subscribe(() => {
          console.log('đang load thêm tin nhắn');
        });
      this.loadMore.emit();
    }
  }

  private createFormGroup(): FormGroup {
    return new FormGroup({
      message: new FormControl('', [Validators.required])
    });
  }

  public onSubmit(): void {
    this.onSubmitFile();
    this.onSubmitText();

    this.clearReplyToMessage();
  }

  private onSubmitText(): void {
    if (this.formGroup.valid) {
      const rawValue = this.formGroup.getRawValue();
      const message = _.get(rawValue, 'message', '');
      this.chatService.sendMessage(message, this.groupSelected.id);
      this.formGroup.patchValue({message: ''});

      this.mockupUISendMessage(message, this.groupSelected);
    }  
  }

  private mockupUISendMessage(message: string, group: Group): void {
    this.submitting = true;
    // setTimeout(() => {
    //   this.messages = [
    //    ...this.messages, new TextMessage(1, group, this.currentUser, message, new Date(), [])
    //   ].map(e => {
    //     return {
    //      ...e,
    //     };
    //   });
    // }, 200);
  //  this.submitting = false; 
  }

  private onSubmitFile(): void {
    this.createFileMessage(this.fileList, this.groupSelected);    
    this.openUploadFile = false;
    this.fileList = [];
  }

  private createFileMessage(uploadFiles: NzUploadFile[], group: Group){
    this.submitting = true;

    const fileMessage = new FileMessage(1, group, this.currentUser, uploadFiles, new Date(), []);
    setTimeout(() => {
       this.messages = [
        ...this.messages, fileMessage
       ].map(e => {
         return {
          ...e,
         };
       });
     }, 200);
    this.submitting = false; 
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

  onReplyToMessage(event) {
    if (!this.openReplyToMessage) {
      this.openReplyToMessage = true;
    }
    this.focusInputField();
    this.messageToReply = event;
    //place cursor inside input
    console.log('replying to : ' + event.content);
  }

  clearReplyToMessage(): void{
    this.messageToReply = null;
    this.openReplyToMessage = false;
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
