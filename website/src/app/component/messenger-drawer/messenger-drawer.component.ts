import {Component, OnInit} from '@angular/core';
import {NzResizeEvent} from 'ng-zorro-antd/resizable';
import {NzModalService} from 'ng-zorro-antd';

@Component({
  selector: 'app-messenger-drawer',
  templateUrl: './messenger-drawer.component.html',
  styleUrls: ['./messenger-drawer.component.sass']
})
export class MessengerDrawerComponent implements OnInit {

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
