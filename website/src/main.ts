import { enableProdMode } from '@angular/core';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { AppModule } from './app/app.module';
import { environment } from './environments/environment';
import * as Keycloak from 'keycloak-js';
import {KeycloakConfig} from 'keycloak-js';

const initOptions: KeycloakConfig = {
  url: environment.keycloak.url,
  realm: environment.keycloak.realm,
  clientId: environment.keycloak.clientId
};

const keycloak = Keycloak(initOptions);

keycloak.init({onLoad: 'login-required'})
  .then((auth) => {
    if (!auth) {
      window.location.reload();
    } else {
      console.log('Authenticated');
    }

    localStorage.setItem('TOKEN', keycloak.token);
    localStorage.setItem('REFRESH_TOKEN', keycloak.refreshToken);

    setTimeout(() => {
      keycloak.updateToken(1)
        .then((refreshed) => {
          if (refreshed) {
            console.log('Token refreshed' + refreshed);
          } else {
            console.warn('Token not refreshed, valid for '
              + Math.round(keycloak.tokenParsed.exp + keycloak.timeSkew - new Date().getTime() / 1000) + ' seconds');
          }
        }).catch(() => {
        console.error('Failed to refresh token');
      });
    }, 60000);

  }).catch(() => {
  console.error('Authenticated Failed');
});

if (environment.production) {
  enableProdMode();
}

platformBrowserDynamic().bootstrapModule(AppModule)
  .catch(err => console.error(err));
