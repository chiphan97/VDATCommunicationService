import { Injectable } from '@angular/core';
import {ApiService} from '../common/api.service';
import {GroupPayload} from '../../model/payload/group.payload';
import {Observable} from 'rxjs';
import {environment} from '../../../environments/environment';
import {Group} from '../../model/group.model';
import {User} from '../../model/user.model';
import * as _ from 'lodash';

@Injectable({
  providedIn: 'root'
})
export class GroupService {

  private readonly API_ENDPOINT = `${environment.service.apiUrl}${environment.service.endpoint.groups}`;
  private currentSelectedGroupId: string;
  constructor(private apiService: ApiService) { }

  // region GROUP
  public getAllGroup(): Observable<Array<Group>> {
    return new Observable<Array<Group>>(observer => {
      const url = `${this.API_ENDPOINT}`;
      this.apiService.get(url)
        .then(res => {
          const data = res.data;
          const groups: Array<Group> = data.map(item => Group.fromJson(item));
          observer.next(_.uniqBy(groups, 'id'));
        })
        .catch(err => observer.error(err))
        .finally(() => observer.complete());
    });
  }

  public createGroup(groupPayload: GroupPayload): Observable<Group> {
    return new Observable<any>(observer => {
      const url = `${this.API_ENDPOINT}`;
      this.apiService.post(url, groupPayload)
        .then(res => {
          const data = res.data;
          const group = Group.fromJson(data[0]);
          observer.next(group);
        })
        .catch(err => observer.error(err))
        .finally(() => observer.complete());
    });
  }

  public updateNameGroup(groupId: number, nameGroup: string): Observable<Group> {
    return new Observable<Group>(observer => {
      const url = `${this.API_ENDPOINT}/${groupId}`;
      this.apiService.put(url, {nameGroup})
        .then(res => {
          const data = res.data;
          const group = Group.fromJson(data);
          observer.next(group);
        })
        .catch(err => observer.error(err))
        .finally(() => observer.complete());
    });
  }

  public deleteGroup(groupId: number): Observable<boolean> {
    return new Observable<boolean>(observer => {
      const url = `${this.API_ENDPOINT}/${groupId}`;
      this.apiService.delete(url)
        .then(res => {
          if (res.data) {
            observer.next(res.data === true);
          } else {
            observer.next(false);
          }

        })
        .catch(err => observer.error(err))
        .finally(() => observer.complete());
    });
  }
  // endregion

  // region MEMBER GROUP
  public addMemberOfGroup(groupId: number, users: Array<User>): Observable<any> {
    return new Observable<any>(observer => {
      const userIds = users.filter(user => !!user.userId)
        .map(user => user.userId);

      this.apiService.patch(`${this.API_ENDPOINT}/${groupId}/members`, {users: userIds})
        .then(res => {
          observer.next(_.get(res.data, 'result', false));
        })
        .catch(err => observer.error(err))
        .finally(() => observer.complete());
    });
  }

  public getAllMemberOfGroup(groupId: number): Observable<Array<User>> {
    return new Observable<Array<User>>(observer => {
      const url = `${this.API_ENDPOINT}/${groupId}/members`;
      this.apiService.get(url)
        .then(res => {
          const data = res.data;
          const users = data.map(item => User.fromJson(item));
          observer.next(users);
        })
        .catch(err => observer.error(err))
        .finally(() => observer.complete());
    });
  }

  public deleteMemberOfGroup(groupId: number, userId: string): Observable<boolean> {
    return new Observable<boolean>(observer => {
      const url = `${this.API_ENDPOINT}/${groupId}/members/${userId}`;
      this.apiService.delete(url)
        .then(res => {
          observer.next(_.get(res.data, 'result', false));
        })
        .catch(err => observer.error(err))
        .finally(() => observer.complete());
    });
  }

  public memberOutGroup(groupId: number): Observable<boolean> {
    return new Observable<boolean>(observer => {
      const url = `${this.API_ENDPOINT}/${groupId}/members`;
      this.apiService.delete(url)
        .then(res => {
          observer.next(_.get(res.data, 'result', false));
        })
        .catch(err => observer.error(err))
        .finally(() => observer.complete());
    });
  }

  public getCurrentSelectedGroup(): string{
    return this.currentSelectedGroupId;
  }
  // endregion
  public getUserById(group: Group, senderId: string): User {
    return group.members.find(user => user.userId === senderId);
  }
}
