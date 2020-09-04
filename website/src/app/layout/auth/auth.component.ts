import {AfterViewChecked, Component, OnInit} from '@angular/core';
import {KeycloakService} from '../../service/keycloak.service';
import {environment} from '../../../environments/environment';
import {StorageService} from '../../service/storage.service';
import {Router} from '@angular/router';
import * as _ from 'lodash';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.sass']
})
export class AuthComponent implements OnInit, AfterViewChecked {


  constructor(private keycloakService: KeycloakService,
              private router: Router) {
  }

  ngOnInit(): void {
  }

  async ngAfterViewChecked() {
    if (this.keycloakService.authenticated) {
      await this.router.navigateByUrl('/');
    }
  }

  public onLogin(): void {
    const redirectUri = environment.keycloak.redirectUrl;
    this.keycloakService.login({
      redirectUri
    });
  }
}
