import { FC, createContext, useReducer, ReactNode, Dispatch } from 'react';

import type { NFTMintRequestWithStatus } from 'api';
import { Status } from 'constants/values';

const initialContextValue = {
  status: Status.Any,
  page: 1,
  nftMintRequests: [],
  activeNftMintRequest: undefined,
  isLoading: false,
};

interface IState {
  page: number;
  status: Status;
  nftMintRequests: NFTMintRequestWithStatus[];
  activeNftMintRequest?: NFTMintRequestWithStatus;
  isLoading: boolean;
}

type ActionsTypes =
  { type: 'setPage', payload: number } |
  { type: 'setStatus', payload: Status } |
  { type: 'setNftMintRequests', payload: NFTMintRequestWithStatus[] } |
  { type: 'setActiveNftMintRequest', payload: NFTMintRequestWithStatus } |
  { type: 'setLoadingStatus', payload: boolean };

const reducer = (state: IState, action: ActionsTypes) => {
  switch (action.type) {
    case 'setPage':
      return {
        ...state,
        page: action.payload,
      };
    case 'setStatus':
      return {
        ...state,
        status: action.payload,
      };
    case 'setNftMintRequests':
      return {
        ...state,
        nftMintRequests: [...state.nftMintRequests, ...action.payload],
      };
    case 'setActiveNftMintRequest':
      return {
        ...state,
        activeNftMintRequest: action.payload,
      };
    case 'setLoadingStatus':
      return {
        ...state,
        isLoading: action.payload,
      };
    default:
      return state;
  }
};

interface IContextValue {
  state: IState;
  dispatch: Dispatch<ActionsTypes>;
}
export const Context = createContext<IContextValue>({
  state: initialContextValue,
  dispatch: () => null,
});

export const ContextProvider: FC<{ children: ReactNode }> = ({ children }) => {
  const [state, dispatch] = useReducer(reducer, initialContextValue);

  return <Context.Provider value={{ state, dispatch }}>{children}</Context.Provider>;
};
