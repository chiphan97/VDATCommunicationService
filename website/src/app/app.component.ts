import { Component } from '@angular/core';
import {KeycloakService} from './service/auth/keycloak.service';
import {LanguageService} from './service/common/language.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent {

  constructor(private keycloakService: KeycloakService,
              private languageService: LanguageService) {
    this.keycloakService.initKeycloak();
    this.languageService.setDefaultLanguage();
  }

}
