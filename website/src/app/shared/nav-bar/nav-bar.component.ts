import {AfterViewChecked, Component, OnInit} from '@angular/core';
import {MessengerDrawerComponent} from '../../component/messenger-drawer/messenger-drawer.component';
import {NzDrawerService} from 'ng-zorro-antd';
import {KeycloakService} from '../../service/keycloak.service';
import {environment} from '../../../environments/environment';
import {StorageService} from '../../service/storage.service';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.sass']
})
export class NavBarComponent implements OnInit, AfterViewChecked {

  public userInfo: any;

  constructor(private drawerService: NzDrawerService,
              private keycloakService: KeycloakService,
              private storageService: StorageService) { }

  ngOnInit(): void {
  }

  ngAfterViewChecked() {
    setTimeout(() => {
      this.userInfo = this.storageService.getUserInfo();
    }, 1000);
  }

  onOpenMessengerDrawer() {
    const drawerRef = this.drawerService.create<MessengerDrawerComponent>({
      nzTitle: 'Danh sách người dùng online',
      nzContent: MessengerDrawerComponent,
      nzWidth: '300px'
    });

    drawerRef.afterOpen.subscribe(() => {
      console.log('Drawer(Component) open');
    });

    drawerRef.afterClose.subscribe(data => {
      console.log(data);
    });
  }

  public onLogout(): void {
    this.keycloakService.logout({
      redirectUri: environment.keycloak.redirectUrl
    });
  }
}
