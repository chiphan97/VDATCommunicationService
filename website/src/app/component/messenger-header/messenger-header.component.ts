import { Component, OnInit } from '@angular/core';
import {MessengerDrawerComponent} from '../messenger-drawer/messenger-drawer.component';
import {NzDrawerService} from 'ng-zorro-antd';

@Component({
  selector: 'app-messenger-header',
  templateUrl: './messenger-header.component.html',
  styleUrls: ['./messenger-header.component.sass']
})
export class MessengerHeaderComponent implements OnInit {

  constructor(private drawerService: NzDrawerService) { }

  ngOnInit(): void {
  }

  onOpenMessengerDrawer() {
    const drawerRef = this.drawerService.create<MessengerDrawerComponent>({
      nzTitle: 'Tùy chọn',
      nzContent: MessengerDrawerComponent,
      nzWidth: '25vw'
    });

    drawerRef.afterOpen.subscribe(() => {
      console.log('Drawer(Component) open');
    });

    drawerRef.afterClose.subscribe(data => {
      console.log(data);
    });
  }
}
