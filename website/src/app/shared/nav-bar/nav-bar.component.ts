import {AfterViewChecked, Component, OnInit} from '@angular/core';
import {KeycloakService} from '../../service/auth/keycloak.service';
import {environment} from '../../../environments/environment';
import {StorageService} from '../../service/common/storage.service';
import {UserService} from '../../service/collector/user.service';
import {User} from '../../model/user.model';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.sass']
})
export class NavBarComponent implements OnInit {

  public userInfo: User;

  constructor(private keycloakService: KeycloakService,
              private storageService: StorageService,
              private userService: UserService) { }

  ngOnInit(): void {
    this.userInfo = this.storageService.userInfo;

    this.userService.getUserInfo()
      .subscribe(user => {
        this.userInfo = user;
        this.storageService.userInfo = user;
      }, err => {
        this.userInfo = null;
        this.storageService.userInfo = null;
      });
  }

  public onLogout(): void {
    this.userService.logout()
      .subscribe(() => {
        this.keycloakService.logout({
          redirectUri: environment.keycloak.redirectUrl
        });
      });
  }
}
