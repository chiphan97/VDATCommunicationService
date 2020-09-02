import {Component, OnInit} from '@angular/core';
import {UserOnlineService} from '../../service/user-online.service';
import {UserOnline} from '../../model/user-online.model';
import {formatDistanceToNow} from 'date-fns';
import {vi} from 'date-fns/locale';
import {GroupPayload} from '../../model/payload/group.payload';
import {GroupType} from '../../const/group-type.const';
import {GroupService} from '../../service/group.service';

@Component({
  selector: 'app-messenger-drawer',
  templateUrl: './messenger-drawer.component.html',
  styleUrls: ['./messenger-drawer.component.sass']
})
export class MessengerDrawerComponent implements OnInit {
  loading = false;

  public keyword: string;
  public users: Array<UserOnline> = new Array<UserOnline>();

  constructor(private userOnlineService: UserOnlineService,
              private groupService: GroupService) {
    this.users = this.userOnlineService.getUsersOnline();
  }

  ngOnInit(): void {
  }

  distanceDate(date: Date): string {
    return formatDistanceToNow(date, {locale: vi});
  }

  public onSearchChange() {
    if (this.keyword) {
      const listUser = this.userOnlineService.getUsersOnline();
      const listUserFilter = listUser.filter(user => {
        return user.fullName.search(this.keyword);
      });
      this.users = listUserFilter;
    } else {
      this.users = this.userOnlineService.getUsersOnline();
    }
  }

  public onCreateMessage(userId: string) {
    const groupPayload: GroupPayload = {
      nameGroup: '',
      private: true,
      type: GroupType.ONE,
      users: [userId]
    };

    this.groupService.createGroup(groupPayload)
      .subscribe(res => {});
  }
}
