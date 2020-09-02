import { Injectable } from '@angular/core';
import {ApiService} from './api.service';
import {GroupPayload} from '../model/payload/group.payload';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment';
import axios from 'axios';
import {group} from '@angular/animations';

@Injectable({
  providedIn: 'root'
})
export class GroupService {

  constructor(private apiService: ApiService) { }

  public getAllGroup(): Observable<any> {
    return new Observable<any>(observer => {
      const url = `${environment.apiUrl}/api/v1/groups`;

      this.apiService.get(url)
        .then(res => {
          console.log(res);
          observer.next(res);
        });
    });
  }

  public createGroup(groupPayload: GroupPayload): Observable<any> {
    return new Observable<any>(observer => {
      const url = `${environment.apiUrl}/api/v1/groups`;

      this.apiService.post(url, groupPayload)
        .then(res => {
          console.log(res);
          observer.next(res);
        });
    });
  }
}
