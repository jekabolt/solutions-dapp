import axios from 'axios';

export enum QUERIES {
  getNftRequests = 'getNftRequests',
}

export enum MUTATIONS {
  login = 'login',
}

const getAuthHeaders = (authToken: string) => ({
  'Grpc-Metadata-Authorization': `Bearer ${authToken}`,
});

export function login(password: string) {
  return axios.post('/api/auth/login', { password });
};

export function getNftRequests(authToken: string) {
  return axios.get('/api/nft/requests', { headers: getAuthHeaders(authToken) });
};
