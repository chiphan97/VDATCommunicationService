const SERVER_URL = 'localhost:5000';
const IOH_SERVER_URL = 'vdat-mcsvc-collector-staging.vdatlab.com';

export const environment = {
  production: false,
  service: {
    apiUrl: `http://${SERVER_URL}`,
    wsUrl: `ws://${SERVER_URL}`,
    endpoint: {
      groups: '/api/v1/groups',
      user: '/api/v1/user',
      chat: '/chat',
      message: '/message'
    }
  },
  ioh: {
    apiUrl: `https://${IOH_SERVER_URL}/dcs/v1`,
    endpoint: {
      user: 'users'
    }
  },
  keycloak: {
    url: 'https://accounts.vdatlab.com/auth',
    realm: 'vdatlab.com',
    clientId: 'ioh.apps.vdatlab.com',
    redirectUrl: 'http://localhost:4200/auth'
  }
};
