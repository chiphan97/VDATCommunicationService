import {AfterViewInit, Component, OnInit, ViewChild} from '@angular/core';
import {GroupService} from '../../service/collector/group.service';
import {Group} from '../../model/group.model';
import {User} from '../../model/user.model';
import {NzResizeEvent} from 'ng-zorro-antd/resizable';
import {StorageService} from '../../service/common/storage.service';
import {UserService} from '../../service/collector/user.service';
import {Role} from '../../const/role.const';
import {GroupType} from '../../const/group-type.const';
import {GenerateColorService} from '../../service/common/generate-color.service';
import {MessengerSidebarLeftComponent} from './messenger-sidebar-left/messenger-sidebar-left.component';
import {MessengerSidebarRightComponent} from './messenger-sidebar-right/messenger-sidebar-right.component';
import {ActivatedRoute, Router} from '@angular/router';
import * as _ from 'lodash';

@Component({
  selector: 'app-messenger',
  templateUrl: './messenger.component.html',
  styleUrls: ['./messenger.component.sass']
})
export class MessengerComponent implements OnInit, AfterViewInit {

  @ViewChild('messengerSidebarLeftComponent') messengerSidebarLeftComponent: MessengerSidebarLeftComponent;
  @ViewChild('messengerSidebarRightComponent') messengerSidebarRightComponent: MessengerSidebarRightComponent;

  public groups: Array<Group>;
  public groupSelected: Group;
  public memberOfGroup: Array<User>;
  public currentUser: User;
  public currentUserIsDoctor: boolean;
  public isMember: boolean;

  public idSidebarLeftResize = -1;
  public idSidebarRightResize = -1;
  public numColSidebarLeft = 5;
  public numColSidebarRight = 0;
  public numColContent = 14;
  public collapseSidebarRight: boolean;

  private DEFAULT_COL_SIDEBAR_RIGHT = 6;
  private currentGroupIdFromParam: number;

  constructor(private groupService: GroupService,
              private userService: UserService,
              private storageService: StorageService,
              private generateColorService: GenerateColorService,
              private route: ActivatedRoute,
              private router: Router) {
    this.groups = new Array<Group>();
    this.groupSelected = null;
  }

  ngOnInit(): void {
    this.fetchingCurrentUserInfo();
  }

  ngAfterViewInit() {
    this.route.params
      .subscribe(params => {
        this.currentGroupIdFromParam = _.get(params, 'groupId', null);

        if (!!this.currentGroupIdFromParam) {
          this.memberOfGroup = [];
          this.numColSidebarRight = 0;

          this.fetchingListGroup();
        }
      });

    this.fetchingListGroup();
  }

  // region Event
  public onResize({col}: NzResizeEvent, isSidebarLeft: boolean): void {
    if (isSidebarLeft) {
      cancelAnimationFrame(this.idSidebarLeftResize);
      this.idSidebarLeftResize = requestAnimationFrame(() => {
        this.numColSidebarLeft = col;
      });
    } else {
      cancelAnimationFrame(this.idSidebarRightResize);
      this.idSidebarRightResize = requestAnimationFrame(() => {
        this.numColSidebarRight = col;
      });
    }
  }

  public onChangeGroupSelected(group: Group): void {
    if (group && group.id) {
      this.fetchingMemberOfGroup(group.id);
      this.numColSidebarRight = this.DEFAULT_COL_SIDEBAR_RIGHT;
      this.router.navigate(['messages', this.groupSelected.id]);
    } else if (!group) {
      this.fetchingListGroup();
    }
  }

  public onChangeCollapseSidebar(collapsed: boolean) {
    this.numColSidebarRight = collapsed ? 0 : this.DEFAULT_COL_SIDEBAR_RIGHT;
  }

  // endregion

  private fetchingListGroup(): void {
    this.messengerSidebarLeftComponent.toggleLoading(true);
    this.groupService.getAllGroup()
      .subscribe(groups => {
          const groupMap = groups.map(group => {
            group.isOwner = group.owner === this.currentUser.userId;
            return group;
          });

          groupMap.forEach(group => {
            if (group.type === GroupType.ONE) {
              this.groupService.getAllMemberOfGroup(group.id)
                .subscribe(members => {
                    group.members = members;

                    if (members.length >= 2) {
                      const targetUser = members.filter(member => member.userId !== this.currentUser.userId)[0];
                      group.nameGroup = targetUser.fullName;
                    }
                  },
                  error => group.members = []);
            }
          });

          this.groups = groupMap;

          if (!!this.currentGroupIdFromParam) {
            // tslint:disable-next-line:triple-equals
            this.groupSelected = this.groups.find(group => group.id == this.currentGroupIdFromParam);
          } else {
            this.groupSelected = this.groups[0];
            this.router.navigate(['messages', this.groupSelected.id]);
          }

          if (!!this.groupSelected) {
            this.numColSidebarRight = this.DEFAULT_COL_SIDEBAR_RIGHT;
            this.fetchingMemberOfGroup(this.groupSelected.id);
          }
        },
        error => this.groups = [],
        () => this.messengerSidebarLeftComponent.toggleLoading(false));
  }

  private fetchingMemberOfGroup(groupId: number): void {
    this.messengerSidebarRightComponent.toggleLoading(true);
    this.groupService.getAllMemberOfGroup(groupId)
      .subscribe(members => {
          this.memberOfGroup = members.map(member => {
            member.color = this.generateColorService.generate();
            return member;
          });

          this.isMember = !!this.memberOfGroup.find(member => member.userId === this.currentUser.userId);
        },
        error => this.memberOfGroup = [],
        () => this.messengerSidebarRightComponent.toggleLoading(false));
  }

  private fetchingCurrentUserInfo(): void {
    // load current user info in caching
    this.currentUser = this.storageService.userInfo;
    this.currentUserIsDoctor = this.currentUser.role === Role.DOCTOR;
  }
}
