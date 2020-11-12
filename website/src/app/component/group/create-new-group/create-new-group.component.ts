import {Component, OnInit} from '@angular/core';
import {User} from '../../../model/user.model';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {UserService} from '../../../service/collector/user.service';
import {NzMessageService, NzModalRef} from 'ng-zorro-antd';
import {GroupPayload} from '../../../model/payload/group.payload';
import {GroupType} from '../../../const/group-type.const';
import {GroupService} from '../../../service/collector/group.service';

@Component({
  selector: 'app-create-new-group',
  templateUrl: './create-new-group.component.html',
  styleUrls: ['./create-new-group.component.sass']
})
export class CreateNewGroupComponent implements OnInit {

  selectedValue = null;
  isGroupPrivate: boolean;
  loading: boolean;
  suggestions: Array<User>;

  public usersSelected: Array<User>;
  public formGroup: FormGroup;

  constructor(private userService: UserService,
              private groupService: GroupService,
              private modalService: NzModalRef,
              private messageService: NzMessageService) {
    this.formGroup = this.createFormGroup();
    this.suggestions = new Array<User>();
    this.usersSelected = new Array<User>();
  }

  ngOnInit(): void {
  }

  private createFormGroup() {
    return new FormGroup({
      nameGroup: new FormControl('', Validators.required),
      description: new FormControl(''),
      private: new FormControl(false),
    });
  }

  onSubmit(): void {
    if (this.formGroup.valid) {
      this.formGroup.disable();
      this.loading = true;

      const groupPayload: GroupPayload = this.formGroup.getRawValue();
      groupPayload.type = GroupType.MANY;
      groupPayload.private = !groupPayload.private;

      if (this.usersSelected && this.usersSelected.length > 0) {
        const usersFilter = this.usersSelected.filter(user => !!user.userId);
        groupPayload.users = usersFilter.map(user => user.userId);
      }

      this.groupService.createGroup(groupPayload)
        .subscribe(group => {
          if (group) {
            this.messageService.success('Tạo cuộc hội thoại thành công');
            this.modalService.destroy(group);
          } else {
            this.messageService.error('Không thể tạo cuộc hội thoại vào lúc này. Vui lòng thử lại sau');
            this.formGroup.enable();
          }
        }, error => {
          this.formGroup.enable();
          this.messageService.error(error);
          this.loading = false;
        }, () => {
          this.loading = false;
          this.formGroup.enable();
        });
    }
  }

  onClose() {
    this.modalService.destroy(null);
  }
}
