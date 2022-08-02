import axios from 'axios';

export enum QUERIES {
  getNftRequest = 'getNftRequest',
}

export enum MUTATIONS {
  login = 'login',
}

export function login(password: string) {
  return axios.post('/api/auth/login', { password });
}

function getNftRequests(address: string) {
  return axios.get(`/api/nft/requests/${address}`);
}

