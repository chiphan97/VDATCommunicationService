import { Injectable } from '@angular/core';
import {ApiService} from '../common/api.service';
import {Observable} from 'rxjs';
import {environment} from '../../../environments/environment';
import * as _ from 'lodash';
import {User} from '../../model/user.model';
import {KeycloakService} from '../auth/keycloak.service';
import {StorageService} from '../common/storage.service';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private readonly API_ENDPOINT = `${environment.service.apiUrl}${environment.service.endpoint.user}`;

  constructor(private apiService: ApiService,
              private storageService: StorageService,
              private keycloakService: KeycloakService) { }

  public getUserInfo(): Observable<User> {
    return new Observable<User>(observer => {
      this.apiService.get(`${this.API_ENDPOINT}/info`)
        .then(res => {
          const user = User.fromJson(res.data);
          const userFromKeycloak = this.keycloakService.userInfo;

          if (userFromKeycloak) {
            user.firstName = _.get(userFromKeycloak, 'given_name', '');
            user.lastName = _.get(userFromKeycloak, 'family_name', '');
            user.fullName = _.get(userFromKeycloak, 'name', '');
            user.username = _.get(userFromKeycloak, 'preferred_username', '');
          }

          observer.next(user);
        })
        .catch(err => observer.error(err))
        .finally(() => observer.complete());
    });
  }

  public findUserByKeyword(keyword: string, page?: number, pageSize?: number): Observable<Array<User>> {
    return new Observable<any>(observer => {
      this.apiService.get(`${this.API_ENDPOINT}`, {keyword, page, pageSize})
        .then(res => {
          const arr = res.data;
          const users = arr.map(item => User.fromJson(item));
          observer.next(users);
        })
        .catch(err => observer.error(err))
        .finally(() => observer.complete());
    });
  }

  public logout(): Observable<any> {
    return new Observable<any>(observer => {
      const user: User = this.storageService.userInfo;

      if (!!user) {
        this.apiService.delete(`${this.API_ENDPOINT}/online`, {socketId: user.socketId, hostName: user.hostName})
          .then(() => observer.next())
          .catch(err => observer.error(err))
          .catch(() => observer.complete());
      } else {
        observer.complete();
      }
    });
  }
}
