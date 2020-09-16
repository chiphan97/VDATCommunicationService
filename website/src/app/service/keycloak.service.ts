import {Inject, Injectable, PLATFORM_ID} from '@angular/core';
import {isPlatformBrowser} from '@angular/common';
import {environment} from '../../environments/environment';
import {KeycloakInstance, KeycloakLoginOptions, KeycloakInitOptions, KeycloakLogoutOptions} from 'keycloak-js';
import {StorageService} from './storage.service';

@Injectable({
  providedIn: 'root'
})
export class KeycloakService {

  private keycloak: KeycloakInstance;
  private readonly isBrowser: boolean;
  public authenticated: boolean;

  constructor(@Inject(PLATFORM_ID) platformId: any,
              private storageService: StorageService) {
    this.isBrowser = isPlatformBrowser(platformId);
  }

  public initKeycloak(): void {
    if (this.isBrowser) {
      // @ts-ignore
      const Keycloak = require('keycloak-js');
      this.keycloak = Keycloak(this.getKeycloakConfig());

      const initOptions: KeycloakInitOptions = {
        onLoad: 'check-sso',
        checkLoginIframe: true,
        checkLoginIframeInterval: 240,
        idToken: this.storageService.idToken,
        token: this.storageService.accessToken,
        refreshToken: this.storageService.refreshToken
      };

      this.keycloak.init(initOptions)
        .then((auth) => {
          this.authenticated = auth;

          if (!auth) {
            return;
          }

          this.storageService.accessToken = this.keycloak.token;
          this.storageService.refreshToken = this.keycloak.refreshToken;
          this.storageService.idToken = this.keycloak.idToken;

          this.keycloak.loadUserInfo()
            .then(userInfo => {
              this.storageService.userInfo = userInfo;
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

        }).catch(() => {
        this.storageService.clearAuth();
        console.error('Authenticated Failed');
      });
    }
  }

  public login(options?: KeycloakLoginOptions): void {
    this.keycloak.login(options);
  }

  public logout(options?: KeycloakLogoutOptions) {
    this.keycloak.logout(options);
    this.storageService.clearAuth();
  }

  private getKeycloakConfig(): any {
    return {
      url: environment.keycloak.url,
      realm: environment.keycloak.realm,
      clientId: environment.keycloak.clientId
    };
  }
}
