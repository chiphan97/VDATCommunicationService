import {AfterContentChecked, Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges} from '@angular/core';
import {Group} from '../../../model/group.model';
import {NzModalService} from 'ng-zorro-antd';
import {GroupType} from '../../../const/group-type.const';
import {CreateNewGroupComponent} from '../../group/create-new-group/create-new-group.component';
import {KeycloakService} from '../../../service/auth/keycloak.service';
import {GroupService} from '../../../service/collector/group.service';
import {User} from '../../../model/user.model';
import {Role} from '../../../const/role.const';

@Component({
  selector: 'app-chat-sidebar-left',
  templateUrl: './chat-sidebar-left.component.html',
  styleUrls: ['./chat-sidebar-left.component.sass']
})
export class ChatSidebarLeftComponent implements OnInit, OnChanges {

  @Input() changed: boolean;
  @Input() currentUser: User;
  @Input() groupSelected: Group;
  @Output() groupSelectedChange = new EventEmitter<Group>();

  public loading = false;
  public groups: Array<Group>;

  constructor(private modalService: NzModalService,
              private groupService: GroupService,
              private keycloakService: KeycloakService) {
    this.groups = new Array<Group>();
    this.groupSelected = null;
  }

  public isGroup = (type) => type === GroupType.MANY;
  public isGroupPublic = (isPrivate) => isPrivate === false;
  public isDoctor = (role) => role === Role.DOCTOR;
  public isSelected = (groupId: number): boolean => this.groupSelected && this.groupSelected.id === groupId;
  public isOwner = (owner: string): boolean => this.currentUser && owner === this.currentUser.userId;

  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes.changed) {
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

  onDeleteGroup(group: Group): void {
    console.log(group);
  }
}
