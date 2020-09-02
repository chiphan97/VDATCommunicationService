import { Injectable } from '@angular/core';
import {ApiService} from './api.service';
import {GroupPayload} from '../model/payload/group.payload';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class GroupService {

  constructor(private apiService: ApiService) { }

  public createGroup(groupPayload: GroupPayload): Observable<any> {
    return new Observable<any>(observer => {
      const url = `${environment.apiUrl}/groups`;
      this.apiService.post(url, JSON.stringify(groupPayload))
        .then(res => {
          console.log(res);
          observer.next(res);
        });
    });
  }
}
