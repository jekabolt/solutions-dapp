// package: nft
// file: nft/nft.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class ImageList extends jspb.Message { 
    getFullsize(): string;
    setFullsize(value: string): ImageList;
    getCompressed(): string;
    setCompressed(value: string): ImageList;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ImageList.AsObject;
    static toObject(includeInstance: boolean, msg: ImageList): ImageList.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ImageList, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ImageList;
    static deserializeBinaryFromReader(message: ImageList, reader: jspb.BinaryReader): ImageList;
}

export namespace ImageList {
    export type AsObject = {
        fullsize: string,
        compressed: string,
    }
}

export class ImageToUpload extends jspb.Message { 
    getRaw(): string;
    setRaw(value: string): ImageToUpload;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ImageToUpload.AsObject;
    static toObject(includeInstance: boolean, msg: ImageToUpload): ImageToUpload.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ImageToUpload, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ImageToUpload;
    static deserializeBinaryFromReader(message: ImageToUpload, reader: jspb.BinaryReader): ImageToUpload;
}

export namespace ImageToUpload {
    export type AsObject = {
        raw: string,
    }
}

export class NFTMintRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): NFTMintRequest;
    getEthaddress(): string;
    setEthaddress(value: string): NFTMintRequest;
    getTxhash(): string;
    setTxhash(value: string): NFTMintRequest;
    getMintsequencenumber(): number;
    setMintsequencenumber(value: number): NFTMintRequest;
    getDescription(): string;
    setDescription(value: string): NFTMintRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NFTMintRequest.AsObject;
    static toObject(includeInstance: boolean, msg: NFTMintRequest): NFTMintRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NFTMintRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NFTMintRequest;
    static deserializeBinaryFromReader(message: NFTMintRequest, reader: jspb.BinaryReader): NFTMintRequest;
}

export namespace NFTMintRequest {
    export type AsObject = {
        id: number,
        ethaddress: string,
        txhash: string,
        mintsequencenumber: number,
        description: string,
    }
}

export class NFTMintRequestToUpload extends jspb.Message { 
    clearSampleimagesList(): void;
    getSampleimagesList(): Array<ImageToUpload>;
    setSampleimagesList(value: Array<ImageToUpload>): NFTMintRequestToUpload;
    addSampleimages(value?: ImageToUpload, index?: number): ImageToUpload;

    hasNftmintrequest(): boolean;
    clearNftmintrequest(): void;
    getNftmintrequest(): NFTMintRequest | undefined;
    setNftmintrequest(value?: NFTMintRequest): NFTMintRequestToUpload;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NFTMintRequestToUpload.AsObject;
    static toObject(includeInstance: boolean, msg: NFTMintRequestToUpload): NFTMintRequestToUpload.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NFTMintRequestToUpload, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NFTMintRequestToUpload;
    static deserializeBinaryFromReader(message: NFTMintRequestToUpload, reader: jspb.BinaryReader): NFTMintRequestToUpload;
}

export namespace NFTMintRequestToUpload {
    export type AsObject = {
        sampleimagesList: Array<ImageToUpload.AsObject>,
        nftmintrequest?: NFTMintRequest.AsObject,
    }
}

export class NFTMintRequestWithStatus extends jspb.Message { 
    clearSampleimagesList(): void;
    getSampleimagesList(): Array<ImageList>;
    setSampleimagesList(value: Array<ImageList>): NFTMintRequestWithStatus;
    addSampleimages(value?: ImageList, index?: number): ImageList;

    hasNftmintrequest(): boolean;
    clearNftmintrequest(): void;
    getNftmintrequest(): NFTMintRequest | undefined;
    setNftmintrequest(value?: NFTMintRequest): NFTMintRequestWithStatus;
    getStatus(): string;
    setStatus(value: string): NFTMintRequestWithStatus;
    getNftoffchainurl(): string;
    setNftoffchainurl(value: string): NFTMintRequestWithStatus;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NFTMintRequestWithStatus.AsObject;
    static toObject(includeInstance: boolean, msg: NFTMintRequestWithStatus): NFTMintRequestWithStatus.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NFTMintRequestWithStatus, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NFTMintRequestWithStatus;
    static deserializeBinaryFromReader(message: NFTMintRequestWithStatus, reader: jspb.BinaryReader): NFTMintRequestWithStatus;
}

export namespace NFTMintRequestWithStatus {
    export type AsObject = {
        sampleimagesList: Array<ImageList.AsObject>,
        nftmintrequest?: NFTMintRequest.AsObject,
        status: string,
        nftoffchainurl: string,
    }
}

export class NFTMintRequestListArray extends jspb.Message { 
    clearNftmintrequestsList(): void;
    getNftmintrequestsList(): Array<NFTMintRequestWithStatus>;
    setNftmintrequestsList(value: Array<NFTMintRequestWithStatus>): NFTMintRequestListArray;
    addNftmintrequests(value?: NFTMintRequestWithStatus, index?: number): NFTMintRequestWithStatus;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NFTMintRequestListArray.AsObject;
    static toObject(includeInstance: boolean, msg: NFTMintRequestListArray): NFTMintRequestListArray.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NFTMintRequestListArray, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NFTMintRequestListArray;
    static deserializeBinaryFromReader(message: NFTMintRequestListArray, reader: jspb.BinaryReader): NFTMintRequestListArray;
}

export namespace NFTMintRequestListArray {
    export type AsObject = {
        nftmintrequestsList: Array<NFTMintRequestWithStatus.AsObject>,
    }
}

export class DeleteId extends jspb.Message { 
    getId(): string;
    setId(value: string): DeleteId;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteId.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteId): DeleteId.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteId, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteId;
    static deserializeBinaryFromReader(message: DeleteId, reader: jspb.BinaryReader): DeleteId;
}

export namespace DeleteId {
    export type AsObject = {
        id: string,
    }
}

export class DeleteStatus extends jspb.Message { 
    getMessage(): string;
    setMessage(value: string): DeleteStatus;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteStatus.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteStatus): DeleteStatus.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteStatus, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteStatus;
    static deserializeBinaryFromReader(message: DeleteStatus, reader: jspb.BinaryReader): DeleteStatus;
}

export namespace DeleteStatus {
    export type AsObject = {
        message: string,
    }
}

export class UpdateNFTOffchainUrlRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): UpdateNFTOffchainUrlRequest;

    hasNftoffchainurl(): boolean;
    clearNftoffchainurl(): void;
    getNftoffchainurl(): ImageToUpload | undefined;
    setNftoffchainurl(value?: ImageToUpload): UpdateNFTOffchainUrlRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateNFTOffchainUrlRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateNFTOffchainUrlRequest): UpdateNFTOffchainUrlRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateNFTOffchainUrlRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateNFTOffchainUrlRequest;
    static deserializeBinaryFromReader(message: UpdateNFTOffchainUrlRequest, reader: jspb.BinaryReader): UpdateNFTOffchainUrlRequest;
}

export namespace UpdateNFTOffchainUrlRequest {
    export type AsObject = {
        id: string,
        nftoffchainurl?: ImageToUpload.AsObject,
    }
}

export class MetadataOffchainUrl extends jspb.Message { 
    getUrl(): string;
    setUrl(value: string): MetadataOffchainUrl;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MetadataOffchainUrl.AsObject;
    static toObject(includeInstance: boolean, msg: MetadataOffchainUrl): MetadataOffchainUrl.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MetadataOffchainUrl, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MetadataOffchainUrl;
    static deserializeBinaryFromReader(message: MetadataOffchainUrl, reader: jspb.BinaryReader): MetadataOffchainUrl;
}

export namespace MetadataOffchainUrl {
    export type AsObject = {
        url: string,
    }
}

export class BurnRequest extends jspb.Message { 
    getTxid(): string;
    setTxid(value: string): BurnRequest;
    getAddress(): string;
    setAddress(value: string): BurnRequest;
    getMintsequencenumber(): number;
    setMintsequencenumber(value: number): BurnRequest;

    hasShipping(): boolean;
    clearShipping(): void;
    getShipping(): ShippingTo | undefined;
    setShipping(value?: ShippingTo): BurnRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BurnRequest.AsObject;
    static toObject(includeInstance: boolean, msg: BurnRequest): BurnRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BurnRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BurnRequest;
    static deserializeBinaryFromReader(message: BurnRequest, reader: jspb.BinaryReader): BurnRequest;
}

export namespace BurnRequest {
    export type AsObject = {
        txid: string,
        address: string,
        mintsequencenumber: number,
        shipping?: ShippingTo.AsObject,
    }
}

export class ShippingTo extends jspb.Message { 
    getFullname(): string;
    setFullname(value: string): ShippingTo;
    getAddress(): string;
    setAddress(value: string): ShippingTo;
    getZipcode(): string;
    setZipcode(value: string): ShippingTo;
    getCity(): string;
    setCity(value: string): ShippingTo;
    getCountry(): string;
    setCountry(value: string): ShippingTo;
    getEmail(): string;
    setEmail(value: string): ShippingTo;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ShippingTo.AsObject;
    static toObject(includeInstance: boolean, msg: ShippingTo): ShippingTo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ShippingTo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ShippingTo;
    static deserializeBinaryFromReader(message: ShippingTo, reader: jspb.BinaryReader): ShippingTo;
}

export namespace ShippingTo {
    export type AsObject = {
        fullname: string,
        address: string,
        zipcode: string,
        city: string,
        country: string,
        email: string,
    }
}

export class ShippingStatus extends jspb.Message { 
    getTracknumber(): string;
    setTracknumber(value: string): ShippingStatus;
    getTimesent(): number;
    setTimesent(value: number): ShippingStatus;
    getError(): string;
    setError(value: string): ShippingStatus;
    getSuccess(): boolean;
    setSuccess(value: boolean): ShippingStatus;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ShippingStatus.AsObject;
    static toObject(includeInstance: boolean, msg: ShippingStatus): ShippingStatus.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ShippingStatus, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ShippingStatus;
    static deserializeBinaryFromReader(message: ShippingStatus, reader: jspb.BinaryReader): ShippingStatus;
}

export namespace ShippingStatus {
    export type AsObject = {
        tracknumber: string,
        timesent: number,
        error: string,
        success: boolean,
    }
}

export class ShippingStatusUpdateRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): ShippingStatusUpdateRequest;

    hasStatus(): boolean;
    clearStatus(): void;
    getStatus(): ShippingStatus | undefined;
    setStatus(value?: ShippingStatus): ShippingStatusUpdateRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ShippingStatusUpdateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ShippingStatusUpdateRequest): ShippingStatusUpdateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ShippingStatusUpdateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ShippingStatusUpdateRequest;
    static deserializeBinaryFromReader(message: ShippingStatusUpdateRequest, reader: jspb.BinaryReader): ShippingStatusUpdateRequest;
}

export namespace ShippingStatusUpdateRequest {
    export type AsObject = {
        id: string,
        status?: ShippingStatus.AsObject,
    }
}

export class BurnShippingInfo extends jspb.Message { 
    getId(): number;
    setId(value: number): BurnShippingInfo;

    hasBurn(): boolean;
    clearBurn(): void;
    getBurn(): BurnRequest | undefined;
    setBurn(value?: BurnRequest): BurnShippingInfo;

    hasStatus(): boolean;
    clearStatus(): void;
    getStatus(): ShippingStatus | undefined;
    setStatus(value?: ShippingStatus): BurnShippingInfo;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BurnShippingInfo.AsObject;
    static toObject(includeInstance: boolean, msg: BurnShippingInfo): BurnShippingInfo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BurnShippingInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BurnShippingInfo;
    static deserializeBinaryFromReader(message: BurnShippingInfo, reader: jspb.BinaryReader): BurnShippingInfo;
}

export namespace BurnShippingInfo {
    export type AsObject = {
        id: number,
        burn?: BurnRequest.AsObject,
        status?: ShippingStatus.AsObject,
    }
}

export class BurnList extends jspb.Message { 
    clearDataList(): void;
    getDataList(): Array<BurnShippingInfo>;
    setDataList(value: Array<BurnShippingInfo>): BurnList;
    addData(value?: BurnShippingInfo, index?: number): BurnShippingInfo;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BurnList.AsObject;
    static toObject(includeInstance: boolean, msg: BurnList): BurnList.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BurnList, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BurnList;
    static deserializeBinaryFromReader(message: BurnList, reader: jspb.BinaryReader): BurnList;
}

export namespace BurnList {
    export type AsObject = {
        dataList: Array<BurnShippingInfo.AsObject>,
    }
}
