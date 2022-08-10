import { FC, createContext, useReducer, ReactNode, Dispatch } from 'react';

const initialContexValue = {
  authToken: '',
};

interface IState {
  authToken: string;
}
type ActionsType = 'setAuthToken' | 'resetAuthToken';

const reducer = (state: IState, action: { type: ActionsType; payload: string }) => {
  switch (action.type) {
    case 'setAuthToken':
      return { ...state, authToken: action.payload };
    case 'resetAuthToken':
      return { ...state, authToken: '' };
    default:
      return state;
  }
};

interface IContextValue {
  state: IState;
  dispatch: Dispatch<{ type: ActionsType; payload: string }>;
}
export const Context = createContext<IContextValue>({
  state: initialContexValue,
  dispatch: () => null,
});

export const ContextProvider: FC<{ children: ReactNode }> = ({ children }) => {
  const [state, dispatch] = useReducer(reducer, initialContexValue);

  return <Context.Provider value={{ state, dispatch }}>{children}</Context.Provider>;
};
