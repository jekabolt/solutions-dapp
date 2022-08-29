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

export enum MUTATIONS {
  login = 'login',
}

// proto files doesnt export this type
type RequestType = {
  path: string;
  method: string;
  body: string | null;
};

const getAuthHeaders = (authToken: string) => ({
  'Grpc-Metadata-Authorization': `Bearer ${authToken}`,
});

export function login(password: string): Promise<LoginResponse> {
  const authClient = createAuthClient(({ path, body }: RequestType): Promise<LoginResponse> => {
    return axios
      .post<LoginRequest, AxiosResponse<LoginResponse>>(path, body && JSON.parse(body))
      .then((response) => response.data);
  });

  return authClient.Login({ password });
}

// +GET /api/nft/requests
// +GET /api/nft/burn
// +GET /api/nft/burn/error
// +GET /api/nft/burn/pending

// +POST /api/auth/login
// -POST /api/nft/ipfs
// +POST /api/nft/burn
// +POST /api/nft
// +POST /api/nft/requests
// +POST /api/nft/offchain
// +POST /api/nft/shipping/status

// +DELETE /api/nft/requests/{id}
// +DELETE /api/nft/{id}

const createAuthorizedNftClient = (authToken: string) => {
  return nftProto.createNftClient(
    // todo: fix types
    ({ path, method, body }: RequestType): Promise<any> => {
      switch (method.toLowerCase()) {
        // todo: test
        case 'post':
          return axios
            .post(path, body, { headers: getAuthHeaders(authToken) })
            .then((response) => response.data);
        // todo: test ?? where is the body pf request
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

export function getNftRequests(authToken: string): Promise<nftProto.NFTMintRequestListArray> {
  const nftClient = createAuthorizedNftClient(authToken);
  
  return nftClient.ListNFTMintRequests(authToken);
}

export function getNftAllBurned(authToken: string): Promise<nftProto.BurnList> {
  const nftClient = createAuthorizedNftClient(authToken);
  
  return nftClient.GetAllBurned(authToken);
}

export function getNftAllBurnedError(authToken: string): Promise<nftProto.BurnList> {
  const nftClient = createAuthorizedNftClient(authToken);
  
  return nftClient.GetAllBurnedError(authToken);
}
export function getNftAllBurnedPending(authToken: string): Promise<nftProto.BurnList> {
  const nftClient = createAuthorizedNftClient(authToken);
  
  return nftClient.GetAllBurnedPending(authToken);
}

// todo: handle empty responses

export function uploadIpfsMetadata(authToken: string): Promise<{}> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.UploadIPFSMetadata({});
}

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

export function upsertNftMintRequest(
  authToken: string,
  requestBody: nftProto.NFTMintRequestToUpload,
): Promise<nftProto.NFTMintRequestWithStatus> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.UpsertNFTMintRequest(requestBody);
}

export function uploadOffchainMetadata(authToken: string): Promise<nftProto.MetadataOffchainUrl> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.UploadOffchainMetadata({});
}

export function updateBurnShippingStatus(
  authToken: string,
  requestBody: nftProto.ShippingStatusUpdateRequest,
): Promise<{}> {
  const nftClient = createAuthorizedNftClient(authToken);

  return nftClient.UpdateBurnShippingStatus(requestBody);
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
