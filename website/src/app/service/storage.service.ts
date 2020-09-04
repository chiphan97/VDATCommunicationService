import {Inject, Injectable, PLATFORM_ID} from '@angular/core';
import {isPlatformBrowser} from '@angular/common';

@Injectable({
  providedIn: 'root'
})
export class StorageService {

  private static TOKEN_KEY = 'TOKEN';
  private static REFRESH_TOKEN_KEY = 'TOKEN';
  private static ID_TOKEN_KEY = 'TOKEN';
  private static USER_INFO_KEY = 'USER_INFO';

  private readonly isBrowser: boolean;

  constructor(@Inject(PLATFORM_ID) platformId: any) {
    this.isBrowser = isPlatformBrowser(platformId);
  }

  public setToken(accessToken: string, refreshToken?: string, idToken?: string): void {
    if (this.isBrowser) {
      localStorage.setItem(StorageService.TOKEN_KEY, accessToken);

      if (refreshToken) {
        localStorage.setItem(StorageService.REFRESH_TOKEN_KEY, refreshToken);
      }

      if (idToken) {
        localStorage.setItem(StorageService.ID_TOKEN_KEY, idToken);
      }
    }
  }

  public getToken(): string {
    if (this.isBrowser) {
      return localStorage.getItem(StorageService.TOKEN_KEY);
    }

    return '';
  }

  public getRefreshToken(): string {
    if (this.isBrowser) {
      return localStorage.getItem(StorageService.REFRESH_TOKEN_KEY);
    }

    return '';
  }

  public getIdToken(): string {
    if (this.isBrowser) {
      return localStorage.getItem(StorageService.REFRESH_TOKEN_KEY);
    }

    return '';
  }

  public setUserInfo(userInfo: any): void {
    if (this.isBrowser) {
      localStorage.setItem(StorageService.USER_INFO_KEY, JSON.stringify(userInfo));
    }
  }

  public getUserInfo(): any {
    if (this.isBrowser) {
      const userInfoRaw = localStorage.getItem(StorageService.USER_INFO_KEY);

      if (userInfoRaw) {
        return JSON.parse(userInfoRaw);
      }
    }

    return null;
  }

  public clearAuth(): void {
    if (this.isBrowser) {
      localStorage.removeItem(StorageService.TOKEN_KEY);
      localStorage.removeItem(StorageService.REFRESH_TOKEN_KEY);
      localStorage.removeItem(StorageService.USER_INFO_KEY);
    }
  }
}
