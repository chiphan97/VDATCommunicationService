import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {Group} from '../../../model/group.model';
import {CreateNewGroupComponent} from '../../../component/group/create-new-group/create-new-group.component';
import {NzModalService} from 'ng-zorro-antd';
import * as _ from 'lodash';
import {FormControl, FormGroup} from '@angular/forms';
import {SettingModalComponent} from '../../../component/setting/setting-modal/setting-modal.component';
import {Setting} from '../../../model/setting.model';

@Component({
  selector: 'app-messenger-sidebar-left',
  templateUrl: './messenger-sidebar-left.component.html',
  styleUrls: ['./messenger-sidebar-left.component.sass']
})
export class MessengerSidebarLeftComponent implements OnInit {

  @Input() groups: Array<Group>;
  @Output() groupsChange = new EventEmitter<Array<Group>>();

  @Input() currentUserIsDoctor: boolean;

  @Input() groupSelected: Group;
  @Output() groupSelectedChange = new EventEmitter<Group>();

  public loading: boolean;
  public formSearch: FormGroup;

  constructor(private modalService: NzModalService) {
    this.formSearch = this.createFormSearch();
  }

  ngOnInit(): void {
  }

  // region Event
  public onSelectGroup(group: Group) {
    if (group && group.id) {
      this.groupSelectedChange.emit(group);
    }
  }

  public onShowModalCreateGroup(): void {
    const modalCreate = this.modalService.create<CreateNewGroupComponent, Group>({
      nzTitle: 'Tạo nhóm mới',
      nzContent: CreateNewGroupComponent,
      nzWidth: '40vw'
    });

    modalCreate.afterClose
      .subscribe(group => {
        if (group) {
          const cloneGroups = _.cloneDeep(this.groups);
          cloneGroups.push(group);
          this.groups = _.sortBy(cloneGroups, 'id');

          this.groupsChange.emit(this.groups);
          this.groupSelectedChange.emit(group);
        }
      });
  }

  public onOpenSetting() {
    this.modalService.create<SettingModalComponent>({
      nzTitle: 'Cài đặt',
      nzContent: SettingModalComponent,
      nzWidth: '40vw'
    });
  }
  // endregion

  public isGroupSelected = (groupId: number) => this.groupSelected && groupId === this.groupSelected.id;

  public toggleLoading(loading?: boolean): void {
    if (!!loading) {
      this.loading = loading;
    } else {
      this.loading = !this.loading;
    }
  }

  private createFormSearch(): FormGroup {
    return new FormGroup({
      keyword: new FormControl('')
    });
  }
}
