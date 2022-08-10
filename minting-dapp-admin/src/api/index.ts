import axios, { AxiosResponse } from 'axios';

import { createAuthClient, LoginRequest, LoginResponse } from './proto-http/auth';
import { } from './proto-http/nft';

export enum QUERIES {
  getNftRequests = 'getNftRequests',
}

export enum MUTATIONS {
  login = 'login',
}

// copy of type inside generated file (no export, need to define explicitly)
type RequestType = {
  path: string;
  method: string;
  body: string | null;
};

export function login(password: string): Promise<LoginResponse> {
  const handleRequest = ({ path, body }: RequestType): Promise<LoginResponse> => {
    return axios.post<LoginRequest, AxiosResponse<LoginResponse>>(path, body && JSON.parse(body)).then(response => response.data)
  };

  const authClient = createAuthClient(handleRequest);

  return authClient.Login({ password });
}
