import {Component, Inject, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import * as _ from 'lodash';
import {KeycloakService} from '../../../service/auth/keycloak.service';
import {UserService} from '../../../service/collector/user.service';
import {DOCUMENT} from '@angular/common';

@Component({
  selector: 'app-integrated',
  templateUrl: './integrated.component.html',
  styleUrls: ['./integrated.component.sass']
})
export class IntegratedComponent implements OnInit {

  public authenticated = false;
  public loading = true;

  constructor(@Inject(DOCUMENT) private document: Document,
              private route: ActivatedRoute,
              private userService: UserService,
              private keycloakService: KeycloakService) {
    this.route.queryParams
      .subscribe(params => {
        const idToken = _.get(params, 'idToken', '');
        const accessToken = _.get(params, 'accessToken', '');
        const refreshToken = _.get(params, 'refreshToken', '');

        if (!!idToken && !!accessToken && !!refreshToken) {
          this.keycloakService.idToken = idToken;
          this.keycloakService.accessToken = accessToken;
          this.keycloakService.refreshToken = refreshToken;

          this.userService.getUserInfo()
            .subscribe(userInfo => {
              this.loading = false;
              this.authenticated = !!userInfo;

              setTimeout(() => {
                document.location.href = '/';
              }, 3000);
            });
        }
      });
  }

  ngOnInit(): void {
  }

}
