import {Component, Input, Output, EventEmitter, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import {User} from '../../../model/user.model';
import {Message} from '../../../model/message.model';
import {formatDistance} from 'date-fns';
import { HttpClient } from '@angular/common/http';
import { map, catchError } from 'rxjs/operators';
import { NzUploadFile, NzUploadModule } from 'ng-zorro-antd/upload';

@Component({
  selector: 'app-messenger-message',
  templateUrl: './messenger-message.component.html',
  styleUrls: ['./messenger-message.component.sass']
})
export class MessengerMessageComponent implements OnInit {

  @Input() currentUser: User;
  @Input() message: Message;
  @Input() showChildren: boolean;

  @Output() onReply : EventEmitter<any> = new EventEmitter();

  previewImage: string | undefined = '';
  previewVisible = false;

  constructor(private http: HttpClient) {
  }

  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges): void {

  }

  onMessageReply(): void {
    this.onReply.emit(this.message);
  }

  handlePreviewSentFile = async (file: NzUploadFile) => {
    return this.http
      .post<{ thumbnail: string }>(`https://next.json-generator.com/api/json/get/4ytyBoLK8`, {
        method: 'POST',
        body: file
      })
      .pipe(map(res => res.thumbnail));
  }

  public isOwner = (): boolean => this.currentUser && this.message
    && this.currentUser.userId === this.message.sender.userId

  public formatDistanceTime = (date: Date): string => formatDistance(date, new Date());
}
