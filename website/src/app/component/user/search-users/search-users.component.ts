import { Component, OnInit } from '@angular/core';
import {User} from '../../../model/user.model';
import {UserService} from '../../../service/collector/user.service';

@Component({
  selector: 'app-search-users',
  templateUrl: './search-users.component.html',
  styleUrls: ['./search-users.component.sass']
})
export class SearchUsersComponent implements OnInit {

  public users: Array<User>;
  public usersSelected: Array<User>;
  public loading: boolean;
  public keyword: string;

  constructor(private userService: UserService) {
    this.users = new Array<User>();
    this.usersSelected = new Array<User>();
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
}
