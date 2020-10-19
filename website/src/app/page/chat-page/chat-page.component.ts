import {Component, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import {Group} from '../../model/group.model';
import {NzResizeEvent} from 'ng-zorro-antd/resizable';
import {ActivatedRoute, Router} from '@angular/router';
import {UserService} from '../../service/collector/user.service';
import {User} from '../../model/user.model';
import { CommonModule } from '@angular/common';
import {KeycloakService} from '../../service/auth/keycloak.service';

@Component({
  selector: 'app-chat-page',
  templateUrl: './chat-page.component.html',
  styleUrls: ['./chat-page.component.sass']
})
export class ChatPageComponent implements OnInit {

  public collapseSidebar = false;
  public groupSelected: Group;
  public changed: boolean;
  public colResize = 5;
  public currentUser: User;
  public isMember: boolean;
  private idResize = -1;

  constructor(private route: ActivatedRoute,
              private router: Router,
              private userService: UserService) {
    this.route.queryParams
      .subscribe(params => {
      });

    this.userService.getUserInfo()
      .subscribe(userInfo => this.currentUser = userInfo);
  }

  ngOnInit(): void {
  }

  onEventChange(isChange: boolean) {
    this.changed = isChange;
    this.groupSelected = null;
  }

  onGroupChange(group: Group) {
    this.router.navigate(['messages', group.id]);
  }

  onResize({col}: NzResizeEvent): void {
    cancelAnimationFrame(this.idResize);
    this.idResize = requestAnimationFrame(() => {
      this.colResize = col;
    });
  }
}
