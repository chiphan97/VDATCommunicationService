const SERVER_URL = 'vdat-mcsvc-chat.vdatlab.com';

export const environment = {
  production: true,
  apiUrl: `http://${SERVER_URL}`,
  wsUrl: `wss://${SERVER_URL}`,
  keycloak: {
    url: 'https://accounts.vdatlab.com/auth',
    realm: 'vdatlab.com',
    clientId: 'chat.services.vdatlab.com'
  },
};
