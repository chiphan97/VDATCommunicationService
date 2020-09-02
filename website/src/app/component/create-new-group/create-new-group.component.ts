import { Component, OnInit } from '@angular/core';
import {MentionOnSearchTypes} from 'ng-zorro-antd';
import {UserService} from '../../service/user.service';
import {FormControl, FormGroup} from '@angular/forms';

@Component({
  selector: 'app-create-new-group',
  templateUrl: './create-new-group.component.html',
  styleUrls: ['./create-new-group.component.sass']
})
export class CreateNewGroupComponent implements OnInit {

  keyword = '';
  isGroupPrivate: boolean;
  loading: boolean;
  suggestions = [
    {
      fullName: 'Hồ Quốc Vửng',
      avatar: 'https://minio.nguyenchicuong.dev/public/57187617_2128229763962598_2406036693489549312_o.jpg'
    },
    {
      fullName: 'Nguyễn Thu Thảo',
      avatar: 'https://minio.nguyenchicuong.dev/public/57187617_2128229763962598_2406036693489549312_o.jpg'
    },
    {
      fullName: 'Nguyễn Thế Sơn',
      avatar: 'https://minio.nguyenchicuong.dev/public/57187617_2128229763962598_2406036693489549312_o.jpg'
    },
    {
      fullName: 'Nguyễn Chí Cường',
      avatar: 'https://minio.nguyenchicuong.dev/public/57187617_2128229763962598_2406036693489549312_o.jpg'
    },
    {
      fullName: 'Lê Hồng Nhu Em',
      avatar: 'https://minio.nguyenchicuong.dev/public/57187617_2128229763962598_2406036693489549312_o.jpg'
    }
  ];

  public formGroup: FormGroup;

  constructor(private userService: UserService) {
    this.formGroup = this.createFormGroup();
  }

  valueWith = (data: { fullName: string; avatar: string; }) => data.fullName;

  ngOnInit(): void {
  }

  onSearchChange({ value }: MentionOnSearchTypes): void {
    if (value) {
      this.loading = true;
      this.userService.findUserByKeyword(value)
        .subscribe(users => {
          console.log(users);
        });
    }
  }

  private createFormGroup() {
    return new FormGroup({
      groupName: new FormControl(),
      keyword: new FormControl()
    });
  }
}
