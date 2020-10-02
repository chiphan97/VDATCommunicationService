import {Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges} from '@angular/core';
import {User} from '../../../model/user.model';
import {UserService} from '../../../service/collector/user.service';
import {NzModalRef} from 'ng-zorro-antd';
import * as _ from 'lodash';

@Component({
  selector: 'app-search-users',
  templateUrl: './search-users.component.html',
  styleUrls: ['./search-users.component.sass']
})
export class SearchUsersComponent implements OnInit, OnChanges {

  @Input() usersSelected: Array<User>;
  @Output() usersSelectedChange = new EventEmitter<Array<User>>();

  public users: Array<User>;
  public loading: boolean;
  public keyword: string;

  public isSelected = (userId: string): boolean => !!this.usersSelected.find(user => user.userId === userId);

  constructor(private userService: UserService,
              private modalRef: NzModalRef) {
    this.users = new Array<User>();
  }

  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges) {
    console.log(this.usersSelected);
  }

  public onSearchUsers(): void {
    if (this.keyword && this.keyword.length > 0) {
      const oldValue = _.cloneDeep(this.keyword);

      setTimeout(() => {
        if (this.keyword === oldValue) {
          this.fetchingUsers(oldValue);
        }
      }, 1500);
    }
  }

  public onSelectUser(user: User): void {
    if (!!this.usersSelected.find(iter => iter.userId === user.userId)) {
      _.remove(this.usersSelected, iter => iter.userId === user.userId);
    } else {
      this.usersSelected.push(user);
    }

    this.usersSelectedChange.emit(this.usersSelected);
  }

  private fetchingUsers(keyword: string, page: number = 1, pageSize: number = 10) {
    this.loading = true;
    this.userService.findUserByKeyword(keyword, page, pageSize)
      .subscribe(users => {
        // filter user existed
        this.users = users.filter(user =>
          this.usersSelected.findIndex(iter =>
            iter.userId === user.userId) === -1);
      }, error => {
        this.loading = false;
      }, () => this.loading = false);
  }

  public onCloseModal(): void {
    this.modalRef.close('Closed');
  }
}
