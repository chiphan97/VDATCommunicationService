import {Component, OnInit} from '@angular/core';
import {UserService} from '../../service/collector/user.service';
import {StorageService} from '../../service/common/storage.service';

@Component({
  selector: 'app-master',
  templateUrl: './master.component.html',
  styleUrls: ['./master.component.sass']
})
export class MasterComponent implements OnInit {

  constructor(private userService: UserService,
              private storageService: StorageService) {
  }

  ngOnInit(): void {
    this.userService.getUserInfo()
      .subscribe(userInfo => {
        this.storageService.userInfo = userInfo;
        console.log(userInfo);
      });
  }
}
