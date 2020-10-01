import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {User} from '../../../model/user.model';
import {UserService} from '../../../service/collector/user.service';
import {NzModalRef} from 'ng-zorro-antd';

@Component({
  selector: 'app-search-users',
  templateUrl: './search-users.component.html',
  styleUrls: ['./search-users.component.sass']
})
export class SearchUsersComponent implements OnInit {

  @Input() usersSelected: Array<User> = new Array<User>();
  @Output() usersSelectedChange = new EventEmitter<Array<User>>();

  public users: Array<User>;
  public loading: boolean;
  public keyword: string;

  constructor(private userService: UserService,
              private modalRef: NzModalRef) {
    this.users = new Array<User>();
  }

  ngOnInit(): void {
  }

  public onSearchUsers(): void {
    if (this.keyword && this.keyword.length > 0) {
      this.loading = true;
      this.userService.findUserByKeyword(this.keyword)
        .subscribe(users => {
          this.users = users;
        }, error => {
          this.loading = false;
        }, () => this.loading = false);
    }
  }

  public onCloseModal(): void {
    this.modalRef.close('Closed');
  }
}
