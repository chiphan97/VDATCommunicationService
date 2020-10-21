import {Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges} from '@angular/core';
import {Group} from '../../../model/group.model';
import {User} from '../../../model/user.model';
import {NzMessageService, NzModalService} from 'ng-zorro-antd';
import {GroupService} from '../../../service/collector/group.service';
import {StorageService} from '../../../service/common/storage.service';
import {GenerateColorService} from '../../../service/common/generate-color.service';
import {Router} from '@angular/router';
import * as _ from 'lodash';
import {GroupPayload} from '../../../model/payload/group.payload';
import {GroupType} from '../../../const/group-type.const';
import {AddMemberGroupComponent} from '../../../component/group/add-member-group/add-member-group.component';

@Component({
  selector: 'app-messenger-sidebar-right',
  templateUrl: './messenger-sidebar-right.component.html',
  styleUrls: ['./messenger-sidebar-right.component.sass']
})
export class MessengerSidebarRightComponent implements OnInit, OnChanges {

  @Input() groupSelected: Group;
  @Output() groupSelectedChange = new EventEmitter<Group>();

  @Input() memberOfGroup: Array<User>;
  @Output() memberOfGroupChange = new EventEmitter<Array<User>>();

  @Input() currentUser: User;
  @Input() isMember: boolean;

  public loading: boolean;
  public memberCollapse = true;
  public optionsCollapse = false;
  public groupClone: Group;

  constructor(private modal: NzModalService,
              private messageService: NzMessageService,
              private groupService: GroupService,
              private storageService: StorageService,
              private generateColorService: GenerateColorService,
              private router: Router) { }

  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes.groupSelected) {
      this.memberCollapse = this.groupSelected.isGroup;
      this.optionsCollapse = !this.groupSelected.isGroup;
      this.groupClone = _.cloneDeep(this.groupSelected);
    }
  }

  // region Event
  public onConfirmDelete() {
    this.modal.confirm({
      nzTitle: 'Cảnh báo',
      nzContent: `Bạn có muốn ${this.groupSelected.isOwner ? 'xóa' : 'rời khỏi'} cuộc hội thoại này không ?`,
      nzAutofocus: 'cancel',
      nzOkType: 'danger',
      nzOkText: 'Đồng ý',
      nzCancelText: 'Hủy',
      nzOnOk: () => this.groupSelected.isOwner ? this.deleteGroup(this.groupSelected.id) : this.outGroup(this.groupSelected.id)
    });
  }

  public onCreateMessenger(userId: string): void {
    const groupPayload: GroupPayload = {
      type: GroupType.ONE,
      private: true,
      users: [userId],
      description: null,
      nameGroup: null
    };

    this.groupService.createGroup(groupPayload)
      .subscribe(group => {
        if (group && group.id) {
          this.router.navigate(['messages', group.id]);
        }
      });
  }

  public onDeleteUser(userId: string) {
    this.loading = true;

    this.groupService.deleteMemberOfGroup(this.groupSelected.id, userId)
      .subscribe(result => {
          if (result) {
            this.messageService.success('Đã xóa thành viên ra khỏi cuộc hội thoại.');

            const members = this.memberOfGroup.filter(member => member.userId !== userId);
            this.memberOfGroupChange.emit(members);
          } else {
            this.messageService.error('Không thể xóa thành viên vào lúc này. Vui lòng thử lại sau');
          }
        }, error => {
          this.messageService.error(error);
          this.loading = false;
        },
        () => this.loading = false);
  }

  public onChangeGroupName(): void {
    if (this.groupSelected.nameGroup === this.groupClone.nameGroup) {
      return;
    }

    this.groupService.updateNameGroup(this.groupSelected.id, this.groupSelected.nameGroup)
      .subscribe(group => {
        if (group) {
          this.groupSelectedChange.emit(this.groupSelected);
          this.messageService.success('Cập nhật thông tin nhóm thành công');
        } else {
          this.messageService.error('Không thể cập nhật thông tin nhóm vào lúc này. Vui lòng thử lại sau');
        }
      }, error => {
        this.messageService.error(error);
      });
  }

  public onOpenModalAddMember(): void {
    const modal = this.modal.create<AddMemberGroupComponent, Array<User>>({
      nzTitle: 'Thêm thành viên',
      nzContent: AddMemberGroupComponent,
      nzWidth: '40vw',
      nzComponentParams: {
        groupId: this.groupSelected.id,
        usersSelected: _.cloneDeep(this.memberOfGroup)
      }
    });

    modal.afterClose.subscribe(members => {
      if (members) {
        this.memberOfGroupChange.emit(members);
      }
    });
  }
  // endregion

  public isCurrentUser = (userId: string) => this.currentUser.userId === userId;

  public checkOwner = (userId: string): boolean => _.get(this.currentUser, 'userId', '') === userId;

  public toggleLoading(loading?: boolean) {
    if (!!loading) {
      this.loading = loading;
    } else {
      this.loading = !this.loading;
    }
  }

  private deleteGroup(groupId: number): void {
    const messId = this.messageService.loading('Đang xóa cuộc hội thoại của bạn ...',
      {nzDuration: 0}).messageId;

    this.groupService.deleteGroup(groupId)
      .subscribe(result => {
          this.messageService.remove(messId);

          if (result) {
            this.groupSelectedChange.emit(null);
            this.messageService.success('Đã xóa cuộc hôi thoại.');
          } else {
            this.messageService.error('Không thể xóa cuộc hội thoại vào lúc này. Vui lòng thử lại sau');
          }
        }, error => {
          this.messageService.remove(messId);
          this.messageService.error(error);
        },
        () => this.messageService.remove(messId));
  }

  private outGroup(groupId: number): void {
    const messId = this.messageService.loading('Đang rời khỏi cuộc hội thoại của này ...',
      {nzDuration: 0}).messageId;

    this.groupService.memberOutGroup(groupId)
      .subscribe(result => {
          this.messageService.remove(messId);

          if (result) {
            this.groupSelectedChange.emit(null);
            this.messageService.success('Đã rời khỏi cuộc hôi thoại.');
          } else {
            this.messageService.error('Không thể rời khỏi cuộc hội thoại vào lúc này. Vui lòng thử lại sau');
          }
        }, error => {
          this.messageService.remove(messId);
          this.messageService.error(error);
        },
        () => this.messageService.remove(messId));
  }
}
