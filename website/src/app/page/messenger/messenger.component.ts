import {Component, OnInit} from '@angular/core';
import {NzResizeEvent} from 'ng-zorro-antd/resizable';
import {UserOnlineService} from '../../service/user-online.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-messenger',
  templateUrl: './messenger.component.html',
  styleUrls: ['./messenger.component.sass']
})
export class MessengerComponent implements OnInit {

  public collapseSidebar = true;

  public col = 5;
  id = -1;

  constructor(private userOnlineService: UserOnlineService,
              private route: ActivatedRoute) {
    this.userOnlineService.initWebSocket();
  }

  ngOnInit(): void {
  }

  onResize({col}: NzResizeEvent): void {
    cancelAnimationFrame(this.id);
    this.id = requestAnimationFrame(() => {
      this.col = col;
    });
  }
}
