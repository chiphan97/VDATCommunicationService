import {AfterContentChecked, Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges} from '@angular/core';
import {Group} from '../../../model/group.model';
import {NzMessageService, NzModalService} from 'ng-zorro-antd';
import {GroupType} from '../../../const/group-type.const';
import {CreateNewGroupComponent} from '../../group/create-new-group/create-new-group.component';
import {KeycloakService} from '../../../service/auth/keycloak.service';
import {GroupService} from '../../../service/collector/group.service';
import {User} from '../../../model/user.model';
import {Role} from '../../../const/role.const';
import {ActivatedRoute} from '@angular/router';
import * as _ from 'lodash';

@Component({
  selector: 'app-chat-sidebar-left',
  templateUrl: './chat-sidebar-left.component.html',
  styleUrls: ['./chat-sidebar-left.component.sass']
})
export class ChatSidebarLeftComponent implements OnInit, OnChanges {

  @Input() changed: boolean;
  @Input() currentUser: User;
  @Input() groupSelected: Group;
  @Input() isMember: boolean;

  @Input() refreshGroup: boolean;
  @Output() groupSelectedChange = new EventEmitter<Group>();

  public loading = false;
  public groups: Array<Group>;
  private currentGroupId: number;

  constructor(private route: ActivatedRoute,
              private modalService: NzModalService,
              private messageService: NzMessageService,
              private groupService: GroupService,
              private keycloakService: KeycloakService) {
    this.route.params
      .subscribe(params => {
        this.currentGroupId = _.get(params, 'groupId', null);
      });

    this.groups = new Array<Group>();
    this.groupSelected = null;
  }

  public isGroup = (type) => type === GroupType.MANY;
  public isGroupPublic = (isPrivate) => isPrivate === false;
  public isDoctor = (role) => role === Role.DOCTOR;
  public isSelected = (groupId: number): boolean => this.groupSelected && this.groupSelected.id === groupId;
  public isOwner = (owner: string): boolean => this.currentUser && owner === this.currentUser.userId;

  ngOnInit(): void {
    this.fetchingData();
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes.refreshGroup) {
      this.fetchingData();
    }
  }

  private fetchingData() {
    this.loading = true;
    this.groupService.getAllGroup()
      .subscribe(groups => {
          this.groups = groups;

          if (groups.length > 0 && !this.groupSelected) {
            this.groupSelectedChange.emit(groups[0]);
          }
        }, error => this.groups = [],
        () => this.loading = false);
  }

  showModalCreateGroup(): void {
    const modalCreate = this.modalService.create({
      nzTitle: 'Tạo nhóm mới',
      nzContent: CreateNewGroupComponent,
      nzWidth: '40vw'
    });

    modalCreate.afterClose
      .subscribe(value => {
        if (value === 'created') {
          this.fetchingData();
        }
      });
  }

  onSelectGroup(group: Group): void {
    this.groupSelectedChange.emit(group);
  }

  onConfirmDelete(group: Group) {
    this.modalService.confirm({
      nzTitle: 'Cảnh báo',
      nzContent: 'Bạn có muốn xóa cuộc hội thoại này không ?',
      nzAutofocus: 'cancel',
      nzOkType: 'danger',
      nzOkText: 'Đồng ý',
      nzCancelText: 'Hủy',
      nzOnOk: () => this.deleteGroup(group.id)
    });
  }

  private deleteGroup(groupId: number): void {
    const messId = this.messageService.loading('Đang xóa cuộc hội thoại của bạn ...',
      {nzDuration: 0}).messageId;

    this.groupService.deleteGroup(groupId)
      .subscribe(result => {
          this.messageService.remove(messId);

          if (result) {
            this.messageService.success('Đã xóa cuộc hôi thoại.');
            this.fetchingData();
          } else {
            this.messageService.error('Không thể xóa cuộc hội thoại vào lúc này. Vui lòng thử lại sau');
          }
        }, error => {
          this.messageService.remove(messId);
          this.messageService.error(error);
        },
        () => this.messageService.remove(messId));
  }
}
