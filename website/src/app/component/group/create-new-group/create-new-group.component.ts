import { Component, OnInit } from '@angular/core';
import {User} from '../../../model/user.model';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {UserService} from '../../../service/user.service';
import {GroupService} from '../../../service/group.service';
import {NzMessageService, NzModalRef} from 'ng-zorro-antd';
import {GroupPayload} from '../../../model/payload/group.payload';
import {GroupType} from '../../../const/group-type.const';

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

  public formGroup: FormGroup;

  constructor(private userService: UserService,
              private groupService: GroupService,
              private modalService: NzModalRef,
              private messageService: NzMessageService) {
    this.formGroup = this.createFormGroup();
    this.suggestions = new Array<User>();
  }

  ngOnInit(): void {
  }

  onSearchChange(value: string): void {
    if (value) {
      this.loading = true;
      this.userService.findUserByKeyword(value)
        .subscribe(users => {
          this.suggestions = users;
        }, error => {
          console.log(error);
        }, () => {
          this.loading = false;
        });
    }
  }

  private createFormGroup() {
    return new FormGroup({
      nameGroup: new FormControl('', Validators.required),
      users: new FormControl(null, [Validators.required]),
      private: new FormControl()
    });
  }

  onSubmit(): void {
    if (this.formGroup.valid) {
      this.formGroup.disable();
      this.loading = true;

      const groupPayload: GroupPayload = this.formGroup.getRawValue();
      groupPayload.type = GroupType.MANY;

      this.groupService.createGroup(groupPayload)
        .subscribe(group => {
          if (group) {
            this.messageService.success('Tạo cuộc hội thoại thành công');
            this.modalService.destroy('created');
          } else {
            this.messageService.error('Không thể tạo cuộc hội thoại vào lúc này. Vui lòng thử lại sau');
            this.formGroup.enable();
          }
        }, error => {
          this.formGroup.enable();
          this.messageService.error(error);
        }, () => this.loading = false);
    }
  }

  onClose() {
    this.modalService.destroy('destroy');
  }
}
