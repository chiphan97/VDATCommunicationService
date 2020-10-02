import {Component, Input, OnInit} from '@angular/core';
import {NzMessageService, NzModalRef} from 'ng-zorro-antd';
import {User} from '../../../model/user.model';
import {UserService} from '../../../service/collector/user.service';
import {GroupService} from '../../../service/collector/group.service';

@Component({
  selector: 'app-add-member-group',
  templateUrl: './add-member-group.component.html',
  styleUrls: ['./add-member-group.component.sass']
})
export class AddMemberGroupComponent implements OnInit {

  @Input() groupId: number;

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
      .subscribe(res => {
        console.log(res);
      });
  }

  onClose() {
    this.modalService.destroy('destroy');
  }

}
