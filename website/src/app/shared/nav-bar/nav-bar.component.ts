import {AfterViewChecked, Component, OnInit} from '@angular/core';
import {KeycloakService} from '../../service/auth/keycloak.service';
import {environment} from '../../../environments/environment';
import {StorageService} from '../../service/common/storage.service';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.sass']
})
export class NavBarComponent implements OnInit, AfterViewChecked {

  public userInfo: any;

  constructor(private keycloakService: KeycloakService,
              private storageService: StorageService) { }

  ngOnInit(): void {
  }

  ngAfterViewChecked() {
    setTimeout(() => {
      this.userInfo = this.storageService.userInfo;
    }, 1000);
  }

  public onLogout(): void {
    this.keycloakService.logout({
      redirectUri: environment.keycloak.redirectUrl
    });
  }
}