import {Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges} from '@angular/core';
import {Group} from '../../../model/group.model';
import {User} from '../../../model/user.model';
import {NzMessageService, NzModalService} from 'ng-zorro-antd';
import {GroupService} from '../../../service/collector/group.service';
import {StorageService} from '../../../service/common/storage.service';
import {GroupType} from '../../../const/group-type.const';
import * as _ from 'lodash';
import {AddMemberGroupComponent} from '../../group/add-member-group/add-member-group.component';
import {GenerateColorService} from '../../../service/common/generate-color.service';
import {UserStatus} from '../../../const/user-status.enum';
import {GroupPayload} from '../../../model/payload/group.payload';
import {Router} from '@angular/router';

@Component({
  selector: 'app-chat-sidebar-right',
  templateUrl: './chat-sidebar-right.component.html',
  styleUrls: ['./chat-sidebar-right.component.sass']
})
export class ChatSidebarRightComponent implements OnInit, OnChanges {

  @Input() groupSelected: Group;
  @Output() changeGroup = new EventEmitter<boolean>();

  @Input() refreshGroup: boolean;
  @Output() refreshGroupChange = new EventEmitter<boolean>();

  @Input() isMember: boolean;
  @Output() isMemberChange = new EventEmitter<boolean>();

  @Input() currentUser: User;

  public members: Array<User>;
  public isOwner: boolean;
  public colors: { [userId: string]: string } = {};

  width = 256;
  id = -1;
  memberCollapse = true;
  optionsCollapse = true;

  public loading = false;
  public groupClone: Group;

  constructor(private modal: NzModalService,
              private messageService: NzMessageService,
              private groupService: GroupService,
              private storageService: StorageService,
              private generateColorService: GenerateColorService,
              private router: Router) {
    this.members = new Array<User>();
  }

  isGroup = () => this.groupSelected.type === GroupType.MANY;
  isOnline = (user: User) => user.status === UserStatus.ONLINE;
  isCurrentUser = (userId: string) => this.currentUser.userId === userId;

  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes.groupSelected) {
      this.groupClone = _.cloneDeep(this.groupSelected);
      this.fetchingData();

      this.isOwner = this.currentUser.userId === this.groupSelected.owner;
    }
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
          this.refreshGroupChange.emit(true);
          this.router.navigate(['messages', group.id]);
        }
      });
  }

  public onOpenModalAddMember(): void {
    const modal = this.modal.create<AddMemberGroupComponent, boolean>({
      nzTitle: 'Thêm thành viên',
      nzContent: AddMemberGroupComponent,
      nzWidth: '40vw',
      nzComponentParams: {
        groupId: this.groupSelected.id,
        usersSelected: _.cloneDeep(this.members)
      }
    });

    modal.afterClose.subscribe(result => {
      if (result) {
        this.fetchingData();
      }
    });
  }

  public onConfirmDelete() {
    this.modal.confirm({
      nzTitle: 'Cảnh báo',
      nzContent: `Bạn có muốn ${this.isOwner ? 'xóa' : 'rời khỏi'} cuộc hội thoại này không ?`,
      nzAutofocus: 'cancel',
      nzOkType: 'danger',
      nzOkText: 'Đồng ý',
      nzCancelText: 'Hủy',
      nzOnOk: () => this.isOwner ? this.deleteGroup(this.groupSelected.id) : this.outGroup(this.groupSelected.id)
    });
  }

  public onDeleteUser(userId: string) {
    this.loading = true;

    this.groupService.deleteMemberOfGroup(this.groupSelected.id, userId)
      .subscribe(result => {
          if (result) {
            this.messageService.success('Đã xóa thành viên ra khỏi cuộc hội thoại.');
            this.fetchingData();
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
          this.changeGroup.emit(true);
          this.messageService.success('Cập nhật thông tin nhóm thành công');
        } else {
          this.messageService.error('Không thể cập nhật thông tin nhóm vào lúc này. Vui lòng thử lại sau');
        }
      }, error => {
        this.messageService.error(error);
      });
  }

  public getColor(userId: string): string {
    return this.colors[userId];
  }

  public checkOwner(userId: string): boolean {
    const userInfo = this.storageService.userInfo;
    return _.get(userInfo, 'userId', '') === userId;
  }

  private fetchingData() {
    if (!this.groupSelected || !this.groupSelected.id) {
      return;
    }

    if (this.groupSelected.type === GroupType.MANY) {
      this.loading = true;
      this.groupService.getAllMemberOfGroup(this.groupSelected.id)
        .subscribe(members => {
          this.members = members;
          this.loading = false;
          this.generateColorForUserAvatar();

          this.isMember = !!members.find(member => member.userId === this.currentUser.userId);
          this.isMemberChange.emit(this.isMember);
        }, error => this.members = []);
    } else {
      this.isMemberChange.emit(true);
    }
  }

  private generateColorForUserAvatar(): void {
    this.members.forEach(member => {
      const color = this.generateColorService.generate();
      const userId = member.userId;
      this.colors[userId] = color;
    });
  }

  private deleteGroup(groupId: number): void {
    const messId = this.messageService.loading('Đang xóa cuộc hội thoại của bạn ...',
      {nzDuration: 0}).messageId;

    this.groupService.deleteGroup(groupId)
      .subscribe(result => {
          this.messageService.remove(messId);

          if (result) {
            this.refreshGroupChange.emit(true);
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
            this.refreshGroupChange.emit(true);
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
