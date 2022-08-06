// package: auth
// file: auth/auth.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as auth_auth_pb from "../auth/auth_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

interface IAuthService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    login: IAuthService_ILogin;
}

interface IAuthService_ILogin extends grpc.MethodDefinition<auth_auth_pb.LoginRequest, auth_auth_pb.LoginResponse> {
    path: "/auth.Auth/Login";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<auth_auth_pb.LoginRequest>;
    requestDeserialize: grpc.deserialize<auth_auth_pb.LoginRequest>;
    responseSerialize: grpc.serialize<auth_auth_pb.LoginResponse>;
    responseDeserialize: grpc.deserialize<auth_auth_pb.LoginResponse>;
}

export const AuthService: IAuthService;

export interface IAuthServer {
    login: grpc.handleUnaryCall<auth_auth_pb.LoginRequest, auth_auth_pb.LoginResponse>;
}

export interface IAuthClient {
    login(request: auth_auth_pb.LoginRequest, callback: (error: grpc.ServiceError | null, response: auth_auth_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    login(request: auth_auth_pb.LoginRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_auth_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    login(request: auth_auth_pb.LoginRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_auth_pb.LoginResponse) => void): grpc.ClientUnaryCall;
}

export class AuthClient extends grpc.Client implements IAuthClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public login(request: auth_auth_pb.LoginRequest, callback: (error: grpc.ServiceError | null, response: auth_auth_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    public login(request: auth_auth_pb.LoginRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_auth_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    public login(request: auth_auth_pb.LoginRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_auth_pb.LoginResponse) => void): grpc.ClientUnaryCall;
}
