import {Component, OnInit} from '@angular/core';
import {UserService} from '../../service/collector/user.service';

@Component({
  selector: 'app-master',
  templateUrl: './master.component.html',
  styleUrls: ['./master.component.sass']
})
export class MasterComponent implements OnInit {

  constructor(private userService: UserService) {
  }

  ngOnInit(): void {
    this.userService.getUserInfo()
      .subscribe(userInfo => {
        console.log(userInfo);
      });
  }
}
