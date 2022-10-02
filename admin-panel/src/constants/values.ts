export const AUTH_LOCAL_STORAGE_KEY = 'solutions-dapp-auth-token';

export enum Status {
  Any = 'Any',
  Unknown = 'Unknown',
  Pending = 'Pending',
  Failed = 'Failed',
  UploadedOffchain = 'UploadedOffchain',
  Uploaded = 'Uploaded',
  Burned = 'Burned',
  Shipped = 'Shipped',
}

export const STATUS_COLORS = {
  [Status.Any]: '#BFBFBF',
  [Status.Unknown]: '#FFFA7A',
  [Status.Pending]: '#B5FF57',
  [Status.Failed]: '#FF7557',
  [Status.UploadedOffchain]: '#00FF19',
  [Status.Uploaded]: '#0038FF',
  [Status.Burned]: '#0038FF',
  [Status.Shipped]: '#015211',
};
