import {Inject, Injectable, PLATFORM_ID} from '@angular/core';
import {isPlatformBrowser} from '@angular/common';
import {User} from '../../model/user.model';

@Injectable({
  providedIn: 'root'
})
export class StorageService {

  private readonly USER_INFO = 'VDAT_USER_INFO';
  private readonly isBrowser: boolean;

  constructor(@Inject(PLATFORM_ID) platformId: any) {
    this.isBrowser = isPlatformBrowser(platformId);
  }

  // region User info
  public set userInfo(userInfo: User) {
    if (this.isBrowser) {
      localStorage.setItem(this.USER_INFO, JSON.stringify(userInfo));
    }
  }

  public get userInfo(): User {
    if (this.isBrowser) {
      const userInfoRaw = localStorage.getItem(this.USER_INFO);

      if (userInfoRaw) {
        return JSON.parse(userInfoRaw);
      }
    }

    return null;
  }
  // endregion
}
