const SERVER_URL = 'f6b3406d55d4.ngrok.io';

export const environment = {
  production: false,
  service: {
    apiUrl: `http://${SERVER_URL}`,
    wsUrl: `ws://${SERVER_URL}`,
    endpoint: {
      groups: '/api/v1/groups',
      user: '/api/v1/user',
      chat: '/chat',
      message: 'message'
    }
  },
  keycloak: {
    url: 'https://accounts.vdatlab.com/auth',
    realm: 'vdatlab.com',
    clientId: 'chat.apps.vdatlab.com',
    redirectUrl: 'http://localhost:4200/auth'
  }
};
