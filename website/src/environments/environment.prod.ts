const SERVER_URL = 'vdat-mcsvc-chat.vdatlab.com';

export const environment = {
  production: true,
  service: {
    apiUrl: `http://${SERVER_URL}`,
    wsUrl: `wss://${SERVER_URL}`,
    endpoint: {
      groups: '/api/v1/groups',
      user: '/api/v1/users'
    }
  },
  keycloak: {
    url: 'https://accounts.vdatlab.com/auth',
    realm: 'vdatlab.com',
    clientId: 'chat.app.vdatlab.com'
  },
};
