import { Injectable } from '@angular/core';
import {ApiService} from './api.service';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment';

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
          console.log(res);
        });
    });
  }
}
