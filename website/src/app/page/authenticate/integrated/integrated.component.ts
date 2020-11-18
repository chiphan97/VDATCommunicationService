import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import * as _ from 'lodash';
import {KeycloakService} from '../../../service/auth/keycloak.service';
import {UserService} from '../../../service/collector/user.service';

@Component({
  selector: 'app-integrated',
  templateUrl: './integrated.component.html',
  styleUrls: ['./integrated.component.sass']
})
export class IntegratedComponent implements OnInit {

  public authenticated = false;
  public loading = true;

  constructor(private route: ActivatedRoute,
              private userService: UserService,
              private keycloakService: KeycloakService) {
    this.route.queryParams
      .subscribe(params => {
        const token = _.get(params, 'token', '');

        if (!!token) {
          this.keycloakService.accessToken = token;

          this.userService.getUserInfo()
            .subscribe(userInfo => {
              this.loading = false;
              this.authenticated = !!userInfo;
            });
        }
      });
  }

  ngOnInit(): void {
  }

}
