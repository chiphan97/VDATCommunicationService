import {
  AfterContentInit,
  AfterViewInit,
  Component,
  OnInit,
  ViewChild,
} from '@angular/core';
import { GroupService } from '../../service/collector/group.service';
import { Group } from '../../model/group.model';
import { User } from '../../model/user.model';
import { NzResizeEvent } from 'ng-zorro-antd/resizable';
import { StorageService } from '../../service/common/storage.service';
import { UserService } from '../../service/collector/user.service';
import { Role } from '../../const/role.const';
import { GroupType } from '../../const/group-type.const';
import { GenerateColorService } from '../../service/common/generate-color.service';
import { MessengerSidebarLeftComponent } from './messenger-sidebar-left/messenger-sidebar-left.component';
import { MessengerSidebarRightComponent } from './messenger-sidebar-right/messenger-sidebar-right.component';
import { ActivatedRoute, Router } from '@angular/router';
import * as _ from 'lodash';
import {ChatService} from '../../service/ws/chat.service';
import {Message} from '../../model/message.model';

@Component({
  selector: 'app-messenger',
  templateUrl: './messenger.component.html',
  styleUrls: ['./messenger.component.sass'],
})
export class MessengerComponent
  implements OnInit, AfterContentInit, AfterViewInit {
  @ViewChild('messengerSidebarLeftComponent')
  messengerSidebarLeftComponent: MessengerSidebarLeftComponent;
  @ViewChild('messengerSidebarRightComponent')
  messengerSidebarRightComponent: MessengerSidebarRightComponent;

  public groups: Array<Group>;
  public groupSelected: Group;
  public memberOfGroup: Array<User>;
  public currentUser: User;
  public currentUserIsDoctor: boolean;
  public isMember: boolean;
  public messages: Array<Message>;

  public idSidebarLeftResize = -1;
  public idSidebarRightResize = -1;
  public numColSidebarLeft = 5;
  public numColSidebarRight = 0;
  public numColContent = 14;
  public collapseSidebarRight: boolean;

  private DEFAULT_COL_SIDEBAR_RIGHT = 6;
  private currentGroupIdFromParam: number;

  constructor(
    private groupService: GroupService,
    private userService: UserService,
    private chatService: ChatService,
    private storageService: StorageService,
    private generateColorService: GenerateColorService,
    private route: ActivatedRoute,
    private router: Router
  ) {
    this.groups = new Array<Group>();
    this.groupSelected = null;
    this.messages = new Array<Message>();
  }

  ngOnInit(): void {
    this.fetchingCurrentUserInfo();
  }

  ngAfterViewInit() {
    this.fetchingListGroup();
  }

  ngAfterContentInit() {
    this.route.params.subscribe((params) => {
      this.currentGroupIdFromParam = parseInt(
        _.get(params, 'groupId', null),
        0
      );

      if (!!this.currentGroupIdFromParam) {
        this.memberOfGroup = [];
        this.numColSidebarRight = 0;

        // kiểm tra group có được lưu trữ ở client chưa
        // nếu chưa tiến hành load dữ liệu lại
        const groupFind = this.groups.find(
          (group) => group.id === this.currentGroupIdFromParam
        );

        if (!!!groupFind) {
          setTimeout(
            () => {
              this.fetchingListGroup();
            },
            !!this.messengerSidebarLeftComponent ? 0 : 1000
          );
        }
      }
    });
  }

  // region Event
  public onResize({ col }: NzResizeEvent, isSidebarLeft: boolean): void {
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
      this.numColSidebarRight = this.DEFAULT_COL_SIDEBAR_RIGHT;
      this.router.navigate(['messages', this.groupSelected.id]);
      this.fetchingMemberOfGroup(group.id);
    } else if (!group) {
      this.fetchingListGroup();
    }
  }

  public onChangeCollapseSidebar(collapsed: boolean) {
    this.numColSidebarRight = collapsed ? 0 : this.DEFAULT_COL_SIDEBAR_RIGHT;
  }

  public onLoadMoreMessages() {
  }
  // endregion

  private fetchingListGroup(): void {
    this.messengerSidebarLeftComponent.toggleLoading(true);
    this.groupService.getAllGroup().subscribe(
      (groups) => {
        const groupMap = groups.map((group) => {
          group.isOwner = group.owner === this.currentUser.userId;
          return group;
        });

        groupMap.forEach((group) => {
          if (group.type === GroupType.ONE) {
            this.groupService.getAllMemberOfGroup(group.id).subscribe(
              (members) => {
                group.members = members;

                if (members.length >= 2) {
                  const targetUser = members.filter(
                    (member) => member.userId !== this.currentUser.userId
                  )[0];
                  group.nameGroup = targetUser.fullName;
                }
              },
              () => (group.members = [])
            );
          }
        });

        this.groups = groupMap;

        if (!!this.currentGroupIdFromParam) {
          this.groupSelected = this.groups.find(
            (group) => group.id === this.currentGroupIdFromParam
          );
        } else {
          this.groupSelected = this.groups[0];
          this.router.navigate(['messages', this.groupSelected.id]);
        }

        if (!!this.groupSelected) {
          this.numColSidebarRight = this.DEFAULT_COL_SIDEBAR_RIGHT;
          this.fetchingMemberOfGroup(this.groupSelected.id);
        }

        this.messengerSidebarLeftComponent.toggleLoading(false);
      },
      () => {
        this.groups = [];
        this.messengerSidebarLeftComponent.toggleLoading(false);
      }
    );
  }

  private fetchingMemberOfGroup(groupId: number): void {
    this.messengerSidebarRightComponent.toggleLoading(true);
    this.groupService.getAllMemberOfGroup(groupId).subscribe(
      (members) => {
        this.memberOfGroup = members.map((member) => {
          member.color = this.generateColorService.generate();
          return member;
        });

        this.isMember = !!this.memberOfGroup.find(
          (member) => member.userId === this.currentUser.userId
        );

        this.messengerSidebarRightComponent.toggleLoading(false);

        this.fetchingHistoryMessages();
      },
      () => {
        this.memberOfGroup = [];
        this.messengerSidebarRightComponent.toggleLoading(false);
      }
    );
  }

  private fetchingCurrentUserInfo(): void {
    // load current user info in caching
    this.currentUser = this.storageService.userInfo;
    this.currentUserIsDoctor = this.currentUser.role === Role.DOCTOR;
  }

  private fetchingHistoryMessages(): void {
    this.messages = new Array<Message>();

    this.chatService.getChatHistory(this.groupSelected.id)
      .subscribe(ready => {
        if (ready) {
          this.chatService.listener()
            .subscribe(messageDto => {
              const sender = this.memberOfGroup.find(member => member.userId === messageDto.senderId);

              const message = new Message(
                messageDto.id,
                this.groupSelected.id === messageDto.groupId ? this.groupSelected : null,
                sender,
                messageDto.content,
                messageDto.createdAt
              );

              this.messages.push(message);
              this.messages = [].concat(this.messages);
            });
        }
      });
  }
}
