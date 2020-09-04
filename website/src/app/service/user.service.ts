import { Injectable } from '@angular/core';
import {ApiService} from './api.service';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment';
import * as _ from 'lodash';
import {User} from '../model/user.model';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private apiService: ApiService) { }

  public findUserByKeyword(keyword: string): Observable<any> {
    return new Observable<any>(observer => {
      const url = `${environment.apiUrl}/api/v1/users`;

      this.apiService.get(url, {keyword})
        .then(res => {
          const data = res.data;
          if (_.isArray(data)) {
            const users = data.map(item => User.fromJson(item));
            observer.next(users);
          }
        })
        .catch(err => observer.error(err))
        .finally(() => observer.complete());
    });
  }
}
