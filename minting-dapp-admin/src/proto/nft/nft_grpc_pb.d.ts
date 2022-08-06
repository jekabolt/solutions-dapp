// package: nft
// file: nft/nft.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as nft_nft_pb from "../nft/nft_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

interface INftService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    upsertNFTMintRequest: INftService_IUpsertNFTMintRequest;
    listNFTMintRequests: INftService_IListNFTMintRequests;
    deleteNFTMintRequestById: INftService_IDeleteNFTMintRequestById;
    updateNFTOffchainUrl: INftService_IUpdateNFTOffchainUrl;
    deleteNFTOffchainUrl: INftService_IDeleteNFTOffchainUrl;
    uploadOffchainMetadata: INftService_IUploadOffchainMetadata;
    burn: INftService_IBurn;
    getAllBurned: INftService_IGetAllBurned;
    getAllBurnedPending: INftService_IGetAllBurnedPending;
    getAllBurnedError: INftService_IGetAllBurnedError;
    updateBurnShippingStatus: INftService_IUpdateBurnShippingStatus;
    uploadIPFSMetadata: INftService_IUploadIPFSMetadata;
}

interface INftService_IUpsertNFTMintRequest extends grpc.MethodDefinition<nft_nft_pb.NFTMintRequestToUpload, nft_nft_pb.NFTMintRequestWithStatus> {
    path: "/nft.Nft/UpsertNFTMintRequest";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<nft_nft_pb.NFTMintRequestToUpload>;
    requestDeserialize: grpc.deserialize<nft_nft_pb.NFTMintRequestToUpload>;
    responseSerialize: grpc.serialize<nft_nft_pb.NFTMintRequestWithStatus>;
    responseDeserialize: grpc.deserialize<nft_nft_pb.NFTMintRequestWithStatus>;
}
interface INftService_IListNFTMintRequests extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, nft_nft_pb.NFTMintRequestListArray> {
    path: "/nft.Nft/ListNFTMintRequests";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<nft_nft_pb.NFTMintRequestListArray>;
    responseDeserialize: grpc.deserialize<nft_nft_pb.NFTMintRequestListArray>;
}
interface INftService_IDeleteNFTMintRequestById extends grpc.MethodDefinition<nft_nft_pb.DeleteId, nft_nft_pb.DeleteStatus> {
    path: "/nft.Nft/DeleteNFTMintRequestById";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<nft_nft_pb.DeleteId>;
    requestDeserialize: grpc.deserialize<nft_nft_pb.DeleteId>;
    responseSerialize: grpc.serialize<nft_nft_pb.DeleteStatus>;
    responseDeserialize: grpc.deserialize<nft_nft_pb.DeleteStatus>;
}
interface INftService_IUpdateNFTOffchainUrl extends grpc.MethodDefinition<nft_nft_pb.UpdateNFTOffchainUrlRequest, nft_nft_pb.NFTMintRequestWithStatus> {
    path: "/nft.Nft/UpdateNFTOffchainUrl";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<nft_nft_pb.UpdateNFTOffchainUrlRequest>;
    requestDeserialize: grpc.deserialize<nft_nft_pb.UpdateNFTOffchainUrlRequest>;
    responseSerialize: grpc.serialize<nft_nft_pb.NFTMintRequestWithStatus>;
    responseDeserialize: grpc.deserialize<nft_nft_pb.NFTMintRequestWithStatus>;
}
interface INftService_IDeleteNFTOffchainUrl extends grpc.MethodDefinition<nft_nft_pb.DeleteId, nft_nft_pb.NFTMintRequestWithStatus> {
    path: "/nft.Nft/DeleteNFTOffchainUrl";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<nft_nft_pb.DeleteId>;
    requestDeserialize: grpc.deserialize<nft_nft_pb.DeleteId>;
    responseSerialize: grpc.serialize<nft_nft_pb.NFTMintRequestWithStatus>;
    responseDeserialize: grpc.deserialize<nft_nft_pb.NFTMintRequestWithStatus>;
}
interface INftService_IUploadOffchainMetadata extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, nft_nft_pb.MetadataOffchainUrl> {
    path: "/nft.Nft/UploadOffchainMetadata";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<nft_nft_pb.MetadataOffchainUrl>;
    responseDeserialize: grpc.deserialize<nft_nft_pb.MetadataOffchainUrl>;
}
interface INftService_IBurn extends grpc.MethodDefinition<nft_nft_pb.BurnRequest, google_protobuf_empty_pb.Empty> {
    path: "/nft.Nft/Burn";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<nft_nft_pb.BurnRequest>;
    requestDeserialize: grpc.deserialize<nft_nft_pb.BurnRequest>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}
interface INftService_IGetAllBurned extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, nft_nft_pb.BurnList> {
    path: "/nft.Nft/GetAllBurned";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<nft_nft_pb.BurnList>;
    responseDeserialize: grpc.deserialize<nft_nft_pb.BurnList>;
}
interface INftService_IGetAllBurnedPending extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, nft_nft_pb.BurnList> {
    path: "/nft.Nft/GetAllBurnedPending";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<nft_nft_pb.BurnList>;
    responseDeserialize: grpc.deserialize<nft_nft_pb.BurnList>;
}
interface INftService_IGetAllBurnedError extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, nft_nft_pb.BurnList> {
    path: "/nft.Nft/GetAllBurnedError";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<nft_nft_pb.BurnList>;
    responseDeserialize: grpc.deserialize<nft_nft_pb.BurnList>;
}
interface INftService_IUpdateBurnShippingStatus extends grpc.MethodDefinition<nft_nft_pb.ShippingStatusUpdateRequest, google_protobuf_empty_pb.Empty> {
    path: "/nft.Nft/UpdateBurnShippingStatus";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<nft_nft_pb.ShippingStatusUpdateRequest>;
    requestDeserialize: grpc.deserialize<nft_nft_pb.ShippingStatusUpdateRequest>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}
interface INftService_IUploadIPFSMetadata extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, google_protobuf_empty_pb.Empty> {
    path: "/nft.Nft/UploadIPFSMetadata";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}

export const NftService: INftService;

export interface INftServer {
    upsertNFTMintRequest: grpc.handleUnaryCall<nft_nft_pb.NFTMintRequestToUpload, nft_nft_pb.NFTMintRequestWithStatus>;
    listNFTMintRequests: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, nft_nft_pb.NFTMintRequestListArray>;
    deleteNFTMintRequestById: grpc.handleUnaryCall<nft_nft_pb.DeleteId, nft_nft_pb.DeleteStatus>;
    updateNFTOffchainUrl: grpc.handleUnaryCall<nft_nft_pb.UpdateNFTOffchainUrlRequest, nft_nft_pb.NFTMintRequestWithStatus>;
    deleteNFTOffchainUrl: grpc.handleUnaryCall<nft_nft_pb.DeleteId, nft_nft_pb.NFTMintRequestWithStatus>;
    uploadOffchainMetadata: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, nft_nft_pb.MetadataOffchainUrl>;
    burn: grpc.handleUnaryCall<nft_nft_pb.BurnRequest, google_protobuf_empty_pb.Empty>;
    getAllBurned: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, nft_nft_pb.BurnList>;
    getAllBurnedPending: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, nft_nft_pb.BurnList>;
    getAllBurnedError: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, nft_nft_pb.BurnList>;
    updateBurnShippingStatus: grpc.handleUnaryCall<nft_nft_pb.ShippingStatusUpdateRequest, google_protobuf_empty_pb.Empty>;
    uploadIPFSMetadata: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, google_protobuf_empty_pb.Empty>;
}

export interface INftClient {
    upsertNFTMintRequest(request: nft_nft_pb.NFTMintRequestToUpload, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    upsertNFTMintRequest(request: nft_nft_pb.NFTMintRequestToUpload, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    upsertNFTMintRequest(request: nft_nft_pb.NFTMintRequestToUpload, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    listNFTMintRequests(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestListArray) => void): grpc.ClientUnaryCall;
    listNFTMintRequests(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestListArray) => void): grpc.ClientUnaryCall;
    listNFTMintRequests(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestListArray) => void): grpc.ClientUnaryCall;
    deleteNFTMintRequestById(request: nft_nft_pb.DeleteId, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.DeleteStatus) => void): grpc.ClientUnaryCall;
    deleteNFTMintRequestById(request: nft_nft_pb.DeleteId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.DeleteStatus) => void): grpc.ClientUnaryCall;
    deleteNFTMintRequestById(request: nft_nft_pb.DeleteId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.DeleteStatus) => void): grpc.ClientUnaryCall;
    updateNFTOffchainUrl(request: nft_nft_pb.UpdateNFTOffchainUrlRequest, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    updateNFTOffchainUrl(request: nft_nft_pb.UpdateNFTOffchainUrlRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    updateNFTOffchainUrl(request: nft_nft_pb.UpdateNFTOffchainUrlRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    deleteNFTOffchainUrl(request: nft_nft_pb.DeleteId, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    deleteNFTOffchainUrl(request: nft_nft_pb.DeleteId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    deleteNFTOffchainUrl(request: nft_nft_pb.DeleteId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    uploadOffchainMetadata(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.MetadataOffchainUrl) => void): grpc.ClientUnaryCall;
    uploadOffchainMetadata(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.MetadataOffchainUrl) => void): grpc.ClientUnaryCall;
    uploadOffchainMetadata(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.MetadataOffchainUrl) => void): grpc.ClientUnaryCall;
    burn(request: nft_nft_pb.BurnRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    burn(request: nft_nft_pb.BurnRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    burn(request: nft_nft_pb.BurnRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    getAllBurned(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    getAllBurned(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    getAllBurned(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    getAllBurnedPending(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    getAllBurnedPending(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    getAllBurnedPending(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    getAllBurnedError(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    getAllBurnedError(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    getAllBurnedError(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    updateBurnShippingStatus(request: nft_nft_pb.ShippingStatusUpdateRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    updateBurnShippingStatus(request: nft_nft_pb.ShippingStatusUpdateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    updateBurnShippingStatus(request: nft_nft_pb.ShippingStatusUpdateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    uploadIPFSMetadata(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    uploadIPFSMetadata(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    uploadIPFSMetadata(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
}

export class NftClient extends grpc.Client implements INftClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public upsertNFTMintRequest(request: nft_nft_pb.NFTMintRequestToUpload, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    public upsertNFTMintRequest(request: nft_nft_pb.NFTMintRequestToUpload, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    public upsertNFTMintRequest(request: nft_nft_pb.NFTMintRequestToUpload, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    public listNFTMintRequests(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestListArray) => void): grpc.ClientUnaryCall;
    public listNFTMintRequests(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestListArray) => void): grpc.ClientUnaryCall;
    public listNFTMintRequests(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestListArray) => void): grpc.ClientUnaryCall;
    public deleteNFTMintRequestById(request: nft_nft_pb.DeleteId, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.DeleteStatus) => void): grpc.ClientUnaryCall;
    public deleteNFTMintRequestById(request: nft_nft_pb.DeleteId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.DeleteStatus) => void): grpc.ClientUnaryCall;
    public deleteNFTMintRequestById(request: nft_nft_pb.DeleteId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.DeleteStatus) => void): grpc.ClientUnaryCall;
    public updateNFTOffchainUrl(request: nft_nft_pb.UpdateNFTOffchainUrlRequest, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    public updateNFTOffchainUrl(request: nft_nft_pb.UpdateNFTOffchainUrlRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    public updateNFTOffchainUrl(request: nft_nft_pb.UpdateNFTOffchainUrlRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    public deleteNFTOffchainUrl(request: nft_nft_pb.DeleteId, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    public deleteNFTOffchainUrl(request: nft_nft_pb.DeleteId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    public deleteNFTOffchainUrl(request: nft_nft_pb.DeleteId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.NFTMintRequestWithStatus) => void): grpc.ClientUnaryCall;
    public uploadOffchainMetadata(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.MetadataOffchainUrl) => void): grpc.ClientUnaryCall;
    public uploadOffchainMetadata(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.MetadataOffchainUrl) => void): grpc.ClientUnaryCall;
    public uploadOffchainMetadata(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.MetadataOffchainUrl) => void): grpc.ClientUnaryCall;
    public burn(request: nft_nft_pb.BurnRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public burn(request: nft_nft_pb.BurnRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public burn(request: nft_nft_pb.BurnRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public getAllBurned(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    public getAllBurned(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    public getAllBurned(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    public getAllBurnedPending(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    public getAllBurnedPending(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    public getAllBurnedPending(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    public getAllBurnedError(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    public getAllBurnedError(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    public getAllBurnedError(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: nft_nft_pb.BurnList) => void): grpc.ClientUnaryCall;
    public updateBurnShippingStatus(request: nft_nft_pb.ShippingStatusUpdateRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public updateBurnShippingStatus(request: nft_nft_pb.ShippingStatusUpdateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public updateBurnShippingStatus(request: nft_nft_pb.ShippingStatusUpdateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public uploadIPFSMetadata(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public uploadIPFSMetadata(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public uploadIPFSMetadata(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
}
