import {AfterContentChecked, Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges} from '@angular/core';
import {Group} from '../../../model/group.model';
import {NzModalService} from 'ng-zorro-antd';
import {GroupType} from '../../../const/group-type.const';
import {CreateNewGroupComponent} from '../../group/create-new-group/create-new-group.component';
import {GroupService} from '../../../service/ws/group.service';
import {StorageService} from '../../../service/common/storage.service';
import {KeycloakService} from '../../../service/auth/keycloak.service';

@Component({
  selector: 'app-chat-sidebar-left',
  templateUrl: './chat-sidebar-left.component.html',
  styleUrls: ['./chat-sidebar-left.component.sass']
})
export class ChatSidebarLeftComponent implements OnInit, OnChanges {

  @Input() changed: boolean;
  @Input() groupSelected: Group;
  @Output() groupSelectedChange = new EventEmitter<Group>();

  loading = false;

  public groups: Array<Group>;

  constructor(private modalService: NzModalService,
              private groupService: GroupService,
              private keycloakService: KeycloakService) {
    this.groups = new Array<Group>();
    this.groupSelected = null;
    this.groupService.connect(this.keycloakService.accessToken)
      .subscribe(connected => {
        if (connected) {
          this.fetchingData();
        }
      });
  }

  isGroup = (type) => type === GroupType.MANY;
  isGroupPublic = (isPrivate) => isPrivate === false;

  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes.changed) {
      this.fetchingData();
    }
  }

  private fetchingData() {
    this.loading = true;
    // this.groupService.getAllGroup().subscribe();
    // this.groupService.getAllGroup()
    //   .subscribe(groups => {
    //       this.groups = groups;
    //
    //       if (groups.length > 0 && !this.groupSelected) {
    //         this.groupSelectedChange.emit(groups[0]);
    //       }
    //     }, error => this.groups = [],
    //     () => this.loading = false);
  }

  showModalCreateGroup(): void {
    const modalCreate = this.modalService.create({
      nzTitle: 'Tạo nhóm mới',
      nzContent: CreateNewGroupComponent
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
}
