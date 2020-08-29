import {Component, OnDestroy, OnInit} from '@angular/core';
import {NzResizeEvent} from 'ng-zorro-antd/resizable';
import {NzModalService} from 'ng-zorro-antd';
import {Group} from '../../model/group.model';
import {UserOnlineService} from '../../service/user-online.service';
import {UserOnline} from '../../model/user-online.model';
import {Subscription} from 'rxjs';
import * as _ from 'lodash';
import {formatDistanceToNow} from 'date-fns';
import {vi} from 'date-fns/locale';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-messenger-drawer',
  templateUrl: './messenger-drawer.component.html',
  styleUrls: ['./messenger-drawer.component.sass']
})
export class MessengerDrawerComponent implements OnInit {
  loading = false;

  public keyword: string;
  public users: Array<UserOnline> = new Array<UserOnline>();

  constructor(private userOnlineService: UserOnlineService) {
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
}
