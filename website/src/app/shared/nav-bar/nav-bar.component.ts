import { Component, OnInit } from '@angular/core';
import {MessengerDrawerComponent} from '../../component/messenger-drawer/messenger-drawer.component';
import {NzDrawerService} from 'ng-zorro-antd';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.sass']
})
export class NavBarComponent implements OnInit {

  constructor(private drawerService: NzDrawerService) { }

  ngOnInit(): void {
  }

  onOpenMessengerDrawer() {
    const drawerRef = this.drawerService.create<MessengerDrawerComponent>({
      nzTitle: 'Danh sách người dùng online',
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
