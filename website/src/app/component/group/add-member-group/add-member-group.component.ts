import { Component, OnInit } from '@angular/core';
import {NzModalRef} from 'ng-zorro-antd';
import {User} from '../../../model/user.model';
import {UserService} from '../../../service/collector/user.service';

@Component({
  selector: 'app-add-member-group',
  templateUrl: './add-member-group.component.html',
  styleUrls: ['./add-member-group.component.sass']
})
export class AddMemberGroupComponent implements OnInit {

  public loading: boolean;
  public listUser: Array<User>;

  constructor(private modalService: NzModalRef,
              private userService: UserService) { }

  ngOnInit(): void {
  }

  onSubmit(): void {
    console.log('submit');
  }

  onClose() {
    this.modalService.destroy('destroy');
  }

}
