import {Component, Inject, OnInit} from '@angular/core';
import {DOCUMENT} from '@angular/common';
import {KeycloakService} from '../../../service/auth/keycloak.service';

@Component({
  selector: 'app-authenticate',
  templateUrl: './authenticate.component.html',
  styleUrls: ['./authenticate.component.sass']
})
export class AuthenticateComponent implements OnInit {

  constructor(@Inject(DOCUMENT) private document: Document,
              private keycloakService: KeycloakService) {
    this.keycloakService.getKeycloakInstance()
      .subscribe(keycloak => {
        if (keycloak.authenticated) {
          document.location.href = '/';
        }
      });
  }

  ngOnInit() {
  }

  public onLogin(): void {
    this.keycloakService.login();
  }
}
