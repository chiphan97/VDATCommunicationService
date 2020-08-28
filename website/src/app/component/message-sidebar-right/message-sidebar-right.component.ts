import { Component, OnInit } from '@angular/core';
import {Group} from '../../model/group.model';
import {NzModalService} from 'ng-zorro-antd';
import {NzResizeEvent} from 'ng-zorro-antd/resizable';

@Component({
  selector: 'app-message-sidebar-right',
  templateUrl: './message-sidebar-right.component.html',
  styleUrls: ['./message-sidebar-right.component.sass']
})
export class MessageSidebarRightComponent implements OnInit {

  width = 256;
  id = -1;
  memberCollapse = true;

  loading = false;
  data = [
    {
      title: 'Ant Design Title 1'
    },
    {
      title: 'Ant Design Title 2'
    },
    {
      title: 'Ant Design Title 3'
    },
    {
      title: 'Ant Design Title 4'
    }
  ];

  constructor(private modal: NzModalService) {
  }

  ngOnInit(): void {
  }

  onResize({width}: NzResizeEvent): void {
    cancelAnimationFrame(this.id);
    this.id = requestAnimationFrame(() => {
      this.width = width;
    });
  }

  onConfirmDelete() {
    this.modal.confirm({
      nzTitle: 'Cảnh báo',
      nzContent: 'Bạn có muốn xóa cuộc hội thoại này không ?',
      nzAutofocus: 'cancel',
      nzOkType: 'danger',
      nzOkText: 'Đồng ý',
      nzCancelText: 'Hủy',
      nzOnOk: () => console.log('OK')
    });
  }
}
