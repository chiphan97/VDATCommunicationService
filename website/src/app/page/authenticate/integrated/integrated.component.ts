import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import * as _ from 'lodash';
import {KeycloakService} from '../../../service/auth/keycloak.service';

@Component({
  selector: 'app-integrated',
  templateUrl: './integrated.component.html',
  styleUrls: ['./integrated.component.sass']
})
export class IntegratedComponent implements OnInit {

  constructor(private route: ActivatedRoute,
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
        }
      });
  }

  ngOnInit(): void {
  }

}
