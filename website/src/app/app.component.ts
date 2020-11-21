import {Component, HostListener} from '@angular/core';
import {KeycloakService} from './service/auth/keycloak.service';
import {LanguageService} from './service/common/language.service';
import {UserService} from './service/collector/user.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent {

  constructor(private keycloakService: KeycloakService,
              private languageService: LanguageService,
              private userService: UserService) {
    this.keycloakService.getKeycloakInstance();
    this.languageService.setDefaultLanguage();
  }

  @HostListener('window:beforeunload', [ '$event' ])
  beforeUnloadHandler(event) {
    this.userService.logout().subscribe(() => {
      event.returnValue = true;
    });
  }
}
