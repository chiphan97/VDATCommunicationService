const SERVER_URL = 'c010f9cf6ceb.ngrok.io';

export const environment = {
  production: false,
  apiUrl: `http://${SERVER_URL}`,
  wsUrl: `ws://${SERVER_URL}`,
  keycloak: {
    url: 'https://accounts.vdatlab.com/auth',
    realm: 'vdatlab.com',
    clientId: 'chat.services.vdatlab.com',
    redirectUrl: 'http://localhost:4200/auth'
  }
};
