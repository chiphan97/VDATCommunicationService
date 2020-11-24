import { Injectable } from '@angular/core';
import {ApiService} from '../common/api.service';
import {environment} from '../../../environments/environment';
import {Observable} from 'rxjs';
import {IohUser} from '../../model/ioh/ioh-user.model';

@Injectable({
  providedIn: 'root'
})
export class IohUserService {
  private readonly API_URL_USER = `${environment.ioh.apiUrl}/${environment.ioh.endpoint.user}`;

  constructor(private apiService: ApiService) {
  }

  getUserByUUID(uuid: string): Observable<IohUser> {
    const url = `${this.API_URL_USER}?uuid=${uuid}`;

    return new Observable<IohUser>(observer => {
      this.apiService.get(url)
        .then(value => {
          if (!!value && !!value.data) {
            console.log(value.data);
            const user = IohUser.fromJson(value.data);
            observer.next(user);
          } else {
            observer.next(null);
          }
          observer.complete();
        })
        .catch(err => {
          observer.error(err);
          observer.complete();
        });
    });
  }
}
