const SERVER_URL = 'vdat-mcsvc-chat.vdatlab.com';

export const environment = {
  production: false,
  service: {
    apiUrl: `https://${SERVER_URL}`,
    wsUrl: `wss://${SERVER_URL}`,
    endpoint: {
      groups: '/api/v1/groups',
      user: '/api/v1/user',
      chat: '/chat'
    }
  },
  keycloak: {
    url: 'https://accounts.vdatlab.com/auth',
    realm: 'vdatlab.com',
    clientId: 'chat.apps.vdatlab.com',
    redirectUrl: 'https://vdat-mcsvc-chat.vdatlab.com/auth'
  }
};

