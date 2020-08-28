import {Component, OnInit} from '@angular/core';
import {NzResizeEvent} from 'ng-zorro-antd/resizable';
import {NzModalService} from 'ng-zorro-antd';
import {Group} from '../../model/group.model';
import {UserOnlineService} from '../../service/user-online.service';

@Component({
  selector: 'app-messenger-drawer',
  templateUrl: './messenger-drawer.component.html',
  styleUrls: ['./messenger-drawer.component.sass']
})
export class MessengerDrawerComponent implements OnInit {
  loading = false;

  public groups: Array<Group>;

  constructor(private userOnlineService: UserOnlineService) {
    this.userOnlineService.initWebSocket('eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJGeDdtOW1WNmY3OTZwNDY4RXR3TWFzQi1URlY0bnNoQ011bHc1cDBCWXpvIn0.eyJleHAiOjE1OTg2OTc3NzAsImlhdCI6MTU5ODYxMTM3MCwiYXV0aF90aW1lIjoxNTk4NjExMzY4LCJqdGkiOiI3M2ZjMmEzYi1hMzIyLTRiZGQtOWYyMy00ZjM1NTFhOTc5ZmYiLCJpc3MiOiJodHRwczovL2FjY291bnRzLnZkYXRsYWIuY29tL2F1dGgvcmVhbG1zL3ZkYXRsYWIuY29tIiwic3ViIjoiYjkwMTgzNzktODM5NC00MjA1LTkxMDQtMmQ4NWQ2OTk0M2RiIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiaW9oLmFwcHMudmRhdGxhYi5jb20iLCJzZXNzaW9uX3N0YXRlIjoiMzY4YTc3MmQtMmYyOS00MDFhLWI5OTUtN2NiMWJiMTkyMjQ3IiwiYWNyIjoiMSIsImFsbG93ZWQtb3JpZ2lucyI6WyIqIl0sInNjb3BlIjoib3BlbmlkIGFkZHJlc3MgcGhvbmUgcHJvZmlsZSBlbWFpbCBvZmZsaW5lX2FjY2VzcyIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJhZGRyZXNzIjp7fSwibmFtZSI6IkPGsOG7nW5nIE5ndXnhu4VuIiwicHJlZmVycmVkX3VzZXJuYW1lIjoibmd1eWVuY2hpY3VvbmcyNDAyQGdtYWlsLmNvbSIsImdpdmVuX25hbWUiOiJDxrDhu51uZyIsImZhbWlseV9uYW1lIjoiTmd1eeG7hW4iLCJlbWFpbCI6Im5ndXllbmNoaWN1b25nMjQwMkBnbWFpbC5jb20ifQ.iFTt1d5jKFSWvheaaMl-AhkpgOq6pvMlGDmQhtbECAMbgn26LwSeNRscZ37V93KK46fR2VL6EXyiKBvgFh-sCVlW9L7Hm2PbCcqjgZT_54XkjZ6JU4pPOtteGbCllJG8jOgVIY5D9LOfwUim0o_QZnV8rMPA0MKjxRWbbFmSBW-vLDnMCpfgBGyTyfAcpUJy-C6uT5-rexwjvskvWzGpZ1kp3a7yDj3-zPk5kF-4OZ1Mt_FzfRY2gLjzr53VDp4GAHRVnAQ7nUiJJB9dkxxtPuEZPQbL8QNFpm0xkBxpKgU_GCtEC8n9yP-YDKGvnHdDfB_HeRj-HQW7mpCck4-ung');

    this.userOnlineService.getEventListener()
      .subscribe(value => {
        console.log(value);
      });

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
