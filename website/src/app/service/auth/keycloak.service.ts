import {Injectable} from '@angular/core';
import {environment} from '../../../environments/environment';
import {KeycloakInstance, KeycloakLoginOptions, KeycloakInitOptions, KeycloakLogoutOptions} from 'keycloak-js';
import {StorageConst} from '../../const/storage.const';
import * as Keycloak from 'keycloak-js';

@Injectable({
  providedIn: 'root'
})
export class KeycloakService {
  private readonly ACCESS_TOKEN = StorageConst.KC_ACCESS_TOKEN;
  private readonly REFRESH_TOKEN = StorageConst.KC_REFRESH_TOKEN;
  private readonly ID_TOKEN = StorageConst.KC_ID_TOKEN;
  private readonly USER_INFO = StorageConst.KC_USER_INFO;
  private readonly TIME_REFRESH = 120;
  private readonly REDIRECT_URL = environment.keycloak.redirectUrl;

  private keycloak: KeycloakInstance;
  private readonly isBrowser: boolean;
  public authenticated: boolean;

  public initKeycloak(): void {
    this.keycloak = Keycloak(this.getKeycloakConfig());

    const initOptions: KeycloakInitOptions = {
      onLoad: 'check-sso',
      checkLoginIframe: true,
      checkLoginIframeInterval: this.TIME_REFRESH,
      idToken: this.idToken,
      token: this.accessToken,
      refreshToken: this.refreshToken,
      redirectUri: this.REDIRECT_URL
    };

    this.keycloak.init(initOptions)
      .then((auth) => {
        this.authenticated = auth;

        if (!auth) {
          return;
        }

        this.accessToken = this.keycloak.token;
        this.refreshToken = this.keycloak.refreshToken;
        this.idToken = this.keycloak.idToken;

        this.keycloak.loadUserInfo()
          .then(userInfo => {
            console.log(userInfo);
            this.userInfo = userInfo;
          });

        setTimeout(() => {
          this.keycloak.updateToken(60)
            .then((refreshed) => {
              if (refreshed) {
                console.log('Token refreshed' + refreshed);
              } else {
                console.warn('Token not refreshed, valid for '
                  + Math.round(this.keycloak.tokenParsed.exp + this.keycloak.timeSkew - new Date().getTime() / 1000) + ' seconds');
              }
            }).catch(() => {
            console.error('Failed to refresh token');
          });
        }, 60000);
      })
      .catch(() => {
        console.log('bug');
        this.clearAuth();
        console.error('Authenticated Failed');
      });
  }

  public login(options?: KeycloakLoginOptions): void {
    this.keycloak.login(options);
  }

  public logout(options?: KeycloakLogoutOptions) {
    this.keycloak.logout(options);
    this.clearAuth();
  }

  private getKeycloakConfig(): any {
    return {
      url: environment.keycloak.url,
      realm: environment.keycloak.realm,
      clientId: environment.keycloak.clientId
    };
  }

  // region Storage
  public set userInfo(userInfo: any) {
    localStorage.setItem(this.USER_INFO, JSON.stringify(userInfo));
  }

  public get userInfo(): any {
    const userInfoRaw = localStorage.getItem(this.USER_INFO);

    if (userInfoRaw) {
      return JSON.parse(userInfoRaw);
    }
    return null;
  }

  public set idToken(idToken: string) {
    localStorage.setItem(this.ID_TOKEN, idToken);
  }

  public get idToken(): string {
    return localStorage.getItem(this.ID_TOKEN);
  }

  public set refreshToken(refreshToken: string) {
    localStorage.setItem(this.REFRESH_TOKEN, refreshToken);
  }

  public get refreshToken(): string {
    return localStorage.getItem(this.REFRESH_TOKEN);
  }

  public set accessToken(accessToken: string) {
    localStorage.setItem(this.ACCESS_TOKEN, accessToken);
  }

  public get accessToken(): string {
    return localStorage.getItem(this.ACCESS_TOKEN);
  }

  public clearAuth(): void {
    localStorage.removeItem(this.ACCESS_TOKEN);
    localStorage.removeItem(this.REFRESH_TOKEN);
    localStorage.removeItem(this.USER_INFO);
  }

  // endregion
}
