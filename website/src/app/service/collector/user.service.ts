import { Injectable } from '@angular/core';
import {ApiService} from '../common/api.service';
import {Observable} from 'rxjs';
import {environment} from '../../../environments/environment';
import * as _ from 'lodash';
import {User} from '../../model/user.model';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private readonly API_ENDPOINT = `${environment.service.apiUrl}${environment.service.endpoint.user}`;

  constructor(private apiService: ApiService) { }

  public getUserInfo(): Observable<User> {
    return new Observable<User>(observer => {
      this.apiService.get(`${this.API_ENDPOINT}/info`)
        .then(userInfo => {
          console.log(userInfo);
          observer.next(null);
        })
        .catch(err => observer.error(err))
        .finally(() => observer.complete());
    });
  }

  public findUserByKeyword(keyword: string): Observable<any> {
    return new Observable<any>(observer => {
      observer.next(null);
      observer.complete();
    });
  }
}
