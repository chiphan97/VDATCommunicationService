import { Injectable } from '@angular/core';
import {ApiService} from './common/api.service';
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

      observer.next(null);
      observer.complete();
    });
  }
}
