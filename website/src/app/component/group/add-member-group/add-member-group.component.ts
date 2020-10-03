import {Component, Input, OnInit} from '@angular/core';
import {NzMessageService, NzModalRef} from 'ng-zorro-antd';
import {User} from '../../../model/user.model';
import {GroupService} from '../../../service/collector/group.service';

@Component({
  selector: 'app-add-member-group',
  templateUrl: './add-member-group.component.html',
  styleUrls: ['./add-member-group.component.sass']
})
export class AddMemberGroupComponent implements OnInit {

  @Input() groupId: number;
  @Input() members: Array<User>;

  public usersSelected: Array<User>;

  constructor(private modalService: NzModalRef,
              private messageService: NzMessageService,
              private groupService: GroupService) {
    this.usersSelected = new Array<User>();
  }

  ngOnInit(): void {
  }

  public onAddMembers(): void {
    if (this.usersSelected.length <= 0) {
      this.messageService.warning('Vui lòng chọn thành viên cần thêm');
      return;
    }

    this.groupService.addMemberOfGroup(this.groupId, this.usersSelected)
      .subscribe(result => {
        if (result) {
          this.messageService.success('Thêm thành viên thành công');
          this.modalService.destroy(true);
        } else {
          this.messageService.error('Đã có lỗi xảy ra');
        }
      }, error => this.messageService.error(error));
  }

  onClose() {
    this.modalService.destroy(false);
  }

}
