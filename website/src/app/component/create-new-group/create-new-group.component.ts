import { Component, OnInit } from '@angular/core';
import {MentionOnSearchTypes} from 'ng-zorro-antd';

@Component({
  selector: 'app-create-new-groups',
  templateUrl: './create-new-groups.component.html',
  styleUrls: ['./create-new-groups.component.sass']
})
export class CreateNewGroupComponent implements OnInit {

  inputValue?: string;
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

  valueWith = (data: { fullName: string; avatar: string; }) => data.fullName;

  constructor() { }

  ngOnInit(): void {
  }

  onSearchChange({ value }: MentionOnSearchTypes): void {
    if (value) {
      this.loading = true;
      console.log(this.loading);
      // this.fetchingData();
    }
  }

  fetchingData() {
    setTimeout(() => {
      this.loading = false;
    }, 10000);
  }
}
