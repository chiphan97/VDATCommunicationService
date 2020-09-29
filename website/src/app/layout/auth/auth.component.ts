import {AfterContentInit, AfterViewChecked, AfterViewInit, Component, OnInit} from '@angular/core';
import {KeycloakService} from '../../service/auth/keycloak.service';
import {environment} from '../../../environments/environment';
import {Router} from '@angular/router';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.sass']
})
export class AuthComponent implements OnInit {

  constructor(private keycloakService: KeycloakService,
              private router: Router) {
  }

  ngOnInit() {
    const checkLogin = setInterval(() => {
      console.log(this.keycloakService.authenticated);
      if (this.keycloakService.authenticated) {
        this.router.navigateByUrl('/').then();
        clearInterval(checkLogin);
      }
    }, 1000);
  }

  public onLogin(): void {
    this.keycloakService.login();
  }
}
