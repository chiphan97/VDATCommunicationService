import { enableProdMode } from '@angular/core';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { AppModule } from './app/app.module';
import { environment } from './environments/environment';
import * as Sentry from '@sentry/angular';
import { Integrations } from '@sentry/tracing';

if (environment.production) {
  enableProdMode();

  Sentry.init({
    dsn: 'https://e9fb7c43ec8a414e804e57aab3e28a04@o399174.ingest.sentry.io/5444972',
    integrations: [
      new Integrations.BrowserTracing({
        tracingOrigins: ['localhost'],
        routingInstrumentation: Sentry.routingInstrumentation,
      }),
    ],
    tracesSampleRate: 1.0,
  });
}

platformBrowserDynamic().bootstrapModule(AppModule)
  .catch(err => console.error(err));
