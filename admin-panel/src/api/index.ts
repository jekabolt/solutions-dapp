import axios, { AxiosResponse } from 'axios';

import { createAuthClient, LoginRequest, LoginResponse } from './proto-http/auth';
import * as nftProto from './proto-http/nft';

export * from './proto-http/auth';
export * from './proto-http/nft';

export enum QUERIES {
  getNftRequests = 'getNftRequests',
  getNftAllBurned = 'getNftAllBurned',
  getNftAllBurnedError = 'getNftAllBurnedError',
  getNftAllBurnedPending = 'getNftAllBurnedPending',
}

type RequestType = {
  path: string;
  method: string;
  body: string | null;
};

const getAuthHeaders = (authToken: string) => ({
  'Grpc-Metadata-Authorization': `Bearer ${authToken}`,
});

// used in code
export function login(password: string): Promise<LoginResponse> {
  const authClient = createAuthClient(({ path, body }: RequestType): Promise<LoginResponse> => {
    return axios
      .post<LoginRequest, AxiosResponse<LoginResponse>>(path, body && JSON.parse(body))
      .then((response) => response.data);
  });

  return authClient.Login({ password });
}

const createAuthorizedNftClient = (authToken: string) => {
  return nftProto.createNftClient(
    ({ path, method, body }: RequestType): Promise<any> => {
      switch (method.toLowerCase()) {
        case 'post':
          return axios
            .post(path, body, { headers: getAuthHeaders(authToken) })
            .then((response) => response.data);
        case 'delete':
          return axios
            .delete(path, { headers: getAuthHeaders(authToken) })
            .then((response) => response.data);
        case 'get':
        default:
          return axios
            .get(path, { headers: getAuthHeaders(authToken) })
            .then((response) => response.data);
      }
    },
  );
};

export function getMintRequestsPagedList(
  authToken: string,
  requestBody: nftProto.ListPagedRequest
): Promise<nftProto.NFTMintRequestListArray> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.ListNFTMintRequestsPaged(requestBody);
}

// +
// upload images
export function submitNewNftMintRequest(
  authToken: string,
  requestBody: nftProto.NFTMintRequestToUpload,
): Promise<nftProto.NFTMintRequestWithStatus> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.NewNFTMintRequest(requestBody);
}

// todo: handle empty responses
export function uploadIpfsMetadata(authToken: string): Promise<{}> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.UploadIPFSMetadata({});
}

// +
export function burnNft(authToken: string, requestBody: nftProto.BurnRequest): Promise<{}> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.Burn(requestBody);
}

export function updateNftOffchainUrl(
  authToken: string,
  requestBody: nftProto.UpdateNFTOffchainUrlRequest,
): Promise<nftProto.NFTMintRequestWithStatus> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.UpdateNFTOffchainUrl(requestBody);
}

export function uploadOffchainMetadata(authToken: string): Promise<nftProto.MetadataOffchainUrl> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.UploadOffchainMetadata({});
}

export function deleteNftMintRequestById(
  authToken: string,
  id: nftProto.DeleteId,
): Promise<nftProto.DeleteStatus> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.DeleteNFTMintRequestById(id);
}

export function deleteNftOffchainUrl(
  authToken: string,
  id: nftProto.DeleteId,
): Promise<nftProto.NFTMintRequestWithStatus> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.DeleteNFTOffchainUrl(id);
}

// +
export function submitTrackingNumber(
  authToken: string,
  requestBody: nftProto.SetTrackingNumberRequest,
): Promise<{}> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.SetTrackingNumber(requestBody);
}
