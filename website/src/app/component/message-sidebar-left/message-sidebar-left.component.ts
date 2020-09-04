import {Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges} from '@angular/core';
import {Group} from '../../model/group.model';
import {NzModalService} from 'ng-zorro-antd';
import {CreateNewGroupComponent} from '../create-new-group/create-new-group.component';
import {GroupType} from '../../const/group-type.const';
import {GroupService} from '../../service/group.service';
import {ApiService} from '../../service/api.service';
import {environment} from '../../../environments/environment';

@Component({
  selector: 'app-message-sidebar-left',
  templateUrl: './message-sidebar-left.component.html',
  styleUrls: ['./message-sidebar-left.component.sass']
})
export class MessageSidebarLeftComponent implements OnInit, OnChanges {

  @Input() changed: boolean;
  @Input() groupSelected: Group;
  @Output() groupSelectedChange = new EventEmitter<Group>();

  loading = false;

  public groups: Array<Group>;

  constructor(private modalService: NzModalService,
              private groupService: GroupService) {
    this.groups = new Array<Group>();
    this.groupSelected = null;
  }

  isGroup = (type) => type === GroupType.MANY;
  isGroupPublic = (isPrivate) => isPrivate === false;

  ngOnInit(): void {
    this.fetchingData();
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

          if (groups.length > 0) {
            this.groupSelectedChange.emit(groups[0]);
          }
        }, error => this.groups = [],
        () => this.loading = false);
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
