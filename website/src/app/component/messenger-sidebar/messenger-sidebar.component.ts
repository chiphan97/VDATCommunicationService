import { Component, OnInit } from '@angular/core';
import {Group} from '../../model/group.model';

@Component({
  selector: 'app-messenger-sidebar',
  templateUrl: './messenger-sidebar.component.html',
  styleUrls: ['./messenger-sidebar.component.sass']
})
export class MessengerSidebarComponent implements OnInit {
  loading = false;

  public groups: Array<Group>;

  constructor() {
    this.groups = this.fakeData();
  }

  ngOnInit(): void {
  }

  fakeData(): Array<Group> {
    const group = {
      title: 'Hồ Quốc Vững',
      thumbnail: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
      lastMessage: {
        id: 1,
        groupId: 1,
        sender: {
          subject: '1',
          fullName: 'Nguyễn Chí Cường',
          avatarUrl: ''
        },
        createdAt: new Date(),
        content: 'Hello world !!!'
      }
    };

    const groups = new Array<Group>();
    for (let i = 1; i < 30; i++) {
      groups.push(group);
    }

    return groups;
  }
}
