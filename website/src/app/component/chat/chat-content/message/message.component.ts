import { Component, OnInit, Input } from '@angular/core';
import { User } from '../../../../model/user.model';
import { Message } from '@angular/compiler/src/i18n/i18n_ast';

@Component({
  selector: 'app-message',
  templateUrl: './message.component.html',
  styleUrls: ['./message.component.sass']
})
export class MessageComponent implements OnInit {

  @Input() currentUser: User;
  @Input() messageInput: Message;
  public likes = 0;
  public dislikes = 0;

  public myContext = {$implicit: 'World', message: this.messageInput};

  constructor() { }

  ngOnInit(): void {

  }

  public getFirstname(user: User): string {
    return user.firstName;
  }

  like(): void {
    this.likes = 1;
    this.dislikes = 0;
  }

  dislike(): void {
    this.likes = 0;
    this.dislikes = 1;
  }

}
