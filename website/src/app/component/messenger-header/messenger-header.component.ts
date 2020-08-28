import { Component, OnInit } from '@angular/core';
import {MessengerDrawerComponent} from '../messenger-drawer/messenger-drawer.component';
import {NzDrawerService} from 'ng-zorro-antd';

@Component({
  selector: 'app-messenger-header',
  templateUrl: './messenger-header.component.html',
  styleUrls: ['./messenger-header.component.sass']
})
export class MessengerHeaderComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }


}
