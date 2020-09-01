import { Component } from '@angular/core';
import {ApiService} from './service/api.service';
import {UserOnlineService} from './service/user-online.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent {
  title = 'website';

  constructor(private userOnlineService: UserOnlineService) {
    this.userOnlineService.initWebSocket();
  }

}
