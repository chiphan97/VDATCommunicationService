import {Component, OnInit} from '@angular/core';
import {Group} from '../../model/group.model';
import {NzModalService} from 'ng-zorro-antd';
import {CreateNewGroupComponent} from '../create-new-group/create-new-group.component';
import {GroupType} from '../../const/group-type.const';

@Component({
  selector: 'app-message-sidebar-left',
  templateUrl: './message-sidebar-left.component.html',
  styleUrls: ['./message-sidebar-left.component.sass']
})
export class MessageSidebarLeftComponent implements OnInit {
  loading = false;

  public groups: Array<Group>;

  constructor(private modalService: NzModalService) {
    this.groups = this.fakeData();
  }

  ngOnInit(): void {
  }

  showModalCreateGroup(): void {
    this.modalService.create({
      nzTitle: 'Tạo nhóm mới',
      nzContent: CreateNewGroupComponent,
      nzOkText: 'Tạo nhóm',
      nzCancelText: 'Hủy'
    });
  }

  fakeData(): Array<Group> {
    const group: Group = {
      nameGroup: 'Hồ Quốc Vững',
      private: true,
      type: GroupType.ONE,
      users: [],
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
