const SERVER_URL = '2.tcp.ngrok.io:19429';

export const environment = {
  production: false,
  service: {
    apiUrl: `http://${SERVER_URL}`,
    wsUrl: `ws://${SERVER_URL}`,
    endpoint: {
      groups: '/groups',
      user: '/api/v1/user',
      chat: '/chat'
    }
  },
  keycloak: {
    url: 'https://accounts.vdatlab.com/auth',
    realm: 'vdatlab.com',
    clientId: 'chat.services.vdatlab.com',
    redirectUrl: 'http://localhost:4200/auth'
  }
};
