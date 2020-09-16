import {Inject, Injectable, PLATFORM_ID} from '@angular/core';
import {isPlatformBrowser} from '@angular/common';

@Injectable({
  providedIn: 'root'
})
export class StorageService {

  private static ACCESS_TOKEN = 'ACCESS_TOKEN';
  private static REFRESH_TOKEN = 'REFRESH_TOKEN';
  private static ID_TOKEN = 'ID_TOKEN';
  private static USER_INFO = 'USER_INFO';

  private readonly isBrowser: boolean;

  constructor(@Inject(PLATFORM_ID) platformId: any) {
    this.isBrowser = isPlatformBrowser(platformId);
  }

  // region Access Token
  public set accessToken(accessToken: string) {
    if (this.isBrowser) {
      localStorage.setItem(StorageService.ACCESS_TOKEN, accessToken);
    }
  }
  public get accessToken(): string {
    if (this.isBrowser) {
      return localStorage.getItem(StorageService.ACCESS_TOKEN);
    }

    return '';
  }
  // endregion

  // region Refresh Token
  public set refreshToken(refreshToken: string) {
    if (this.isBrowser) {
      localStorage.setItem(StorageService.REFRESH_TOKEN, refreshToken);
    }
  }
  public get refreshToken(): string {
    if (this.isBrowser) {
      return localStorage.getItem(StorageService.REFRESH_TOKEN);
    }

    return '';
  }
  // endregion

  // region ID Token
  public set idToken(idToken: string) {
    if (this.isBrowser) {
      localStorage.setItem(StorageService.ID_TOKEN, idToken);
    }
  }
  public get idToken(): string {
    if (this.isBrowser) {
      return localStorage.getItem(StorageService.ID_TOKEN);
    }

    return '';
  }
  // endregion

  // region User info
  public set userInfo(userInfo: any) {
    if (this.isBrowser) {
      localStorage.setItem(StorageService.USER_INFO, JSON.stringify(userInfo));
    }
  }
  public get userInfo(): any {
    if (this.isBrowser) {
      const userInfoRaw = localStorage.getItem(StorageService.USER_INFO);

      if (userInfoRaw) {
        return JSON.parse(userInfoRaw);
      }
    }

    return null;
  }
  // endregion

  public clearAuth(): void {
    if (this.isBrowser) {
      localStorage.removeItem(StorageService.ACCESS_TOKEN);
      localStorage.removeItem(StorageService.REFRESH_TOKEN);
      localStorage.removeItem(StorageService.USER_INFO);
    }
  }
}
