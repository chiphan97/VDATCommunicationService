import {AfterContentInit, AfterViewChecked, AfterViewInit, Component, Inject, OnInit} from '@angular/core';
import {KeycloakService} from '../../service/auth/keycloak.service';
import {environment} from '../../../environments/environment';
import {Router} from '@angular/router';
import {DOCUMENT} from '@angular/common';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.sass']
})
export class AuthComponent implements OnInit {

  constructor(@Inject(DOCUMENT) private document: Document,
              private keycloakService: KeycloakService) {
    this.keycloakService.keycloak.onReady = authenticated => {
      if (authenticated) {
        document.location.href = '/';
      }
    };
  }

  ngOnInit() {
  }

  public onLogin(): void {
    this.keycloakService.login();
  }
}
