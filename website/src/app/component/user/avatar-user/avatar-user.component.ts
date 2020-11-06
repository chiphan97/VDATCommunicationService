import {Component, Input, OnInit} from '@angular/core';
import {User} from '../../../model/user.model';

@Component({
  selector: 'app-avatar-user',
  templateUrl: './avatar-user.component.html',
  styleUrls: ['./avatar-user.component.sass']
})
export class AvatarUserComponent implements OnInit {

  @Input() user: User;

  constructor() { }

  ngOnInit(): void {
  }

  get avatarUrl(): string {
    if (!!!this.user) {
      return;
    }

    if (!!this.user.avatar) {
      return this.user.avatar;
    } else {

    }
  }
}
