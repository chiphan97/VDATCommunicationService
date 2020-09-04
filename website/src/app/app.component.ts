import { Component } from '@angular/core';
import {KeycloakService} from './service/keycloak.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent {

  constructor(private keycloakService: KeycloakService) {
    this.keycloakService.initKeycloak();
  }

}
