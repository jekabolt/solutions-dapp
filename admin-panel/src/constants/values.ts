export const AUTH_LOCAL_STORAGE_KEY = 'solutions-dapp-auth-token';

export enum Status {
  Unknown = 'Unknown',
  Pending = 'Pending',
  Done = 'Done',
  Failed = 'Failed',
  Finalised = 'Finalised',
  Burned = 'Burned',
  Shiped = 'Shiped',
}

export const STATUS_COLORS = {
  [Status.Unknown]: '#BFBFBF',
  [Status.Pending]: '#FFFA7A',
  [Status.Done]: '#B5FF57',
  [Status.Failed]: '#FF7557',
  [Status.Finalised]: '#00FF19',
  [Status.Burned]: '#0038FF',
};
