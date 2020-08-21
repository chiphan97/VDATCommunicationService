import {Component, OnDestroy, OnInit} from '@angular/core';
import {SocketService} from '../../service/socket.service';
import {ActivatedRoute} from '@angular/router';
import {MessageModel} from '../../model/message.model';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.sass']
})
export class ChatComponent implements OnInit, OnDestroy {

  public messages: Array<MessageModel>;
  public messageContent: string;
  private accessToken: string;
  private toSubject: string;

  constructor(private socket: SocketService,
              private route: ActivatedRoute) {
    this.route.queryParams.subscribe(params => {
      this.accessToken = params.token;

      if (this.accessToken) {
        this.socket.initWebSocket(this.accessToken);
        this.toSubject = prompt('Subject receiver', '893a4692-63bb-4919-80d9-aece678c0422');
      } else {
        alert('Unauthenticated');
      }
    });

    this.messages = [];
    this.messageContent = '';
  }

  public ngOnInit() {
    this.socket.getEventListener().subscribe(event => {
      if (event.type === 'message') {
        this.messages.push(event.data);
      }
      if (event.type === 'close') {
        this.messages.push(
          {
            from: 'System',
            to: '',
            body: 'The socket connection has been closed'
          });
      }
      if (event.type === 'open') {
        this.messages.push({
          from: 'System',
          to: '',
          body: 'The socket connection has been established'
        });
      }
    });
  }

  public ngOnDestroy() {
    this.socket.close();
  }

  public send() {
    if (this.messageContent) {
      const message: MessageModel = {
        from: null,
        to: this.toSubject,
        body: this.messageContent
      };
      this.messages.push(message);

      this.socket.send(JSON.stringify(message));
      this.messageContent = '';
    }
  }
}
