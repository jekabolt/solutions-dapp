import { FC, useEffect, useState, useContext } from 'react';
import type { AxiosError } from 'axios';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from '@tanstack/react-location';

import { getMintRequestsPagedList, QUERIES, NFTMintRequestListArray } from 'api';
import { ROUTES } from 'constants/routes';
import { AUTH_LOCAL_STORAGE_KEY } from 'constants/values';
import { Header } from 'components/Header';
import { Context } from 'context';
import { StatusOptions } from 'components/StatusOptions';
import { NftList } from 'components/NftList';

import styles from 'styles/HomePage.module.scss';

export const Home: FC = () => {
  const { state, dispatch } = useContext(Context);
  const token = localStorage.getItem(AUTH_LOCAL_STORAGE_KEY) || '';
  const navigate = useNavigate();
  const { data, refetch } = useQuery<NFTMintRequestListArray, AxiosError>(
    [QUERIES.getNftRequests],
    () => getMintRequestsPagedList(token, { status: state.status, page: state.page }),
    {
      retry: false,
      enabled: false,
      onSuccess: (data) => {
        if (data.nftMintRequests) {
          dispatch({ type: 'setNftMintRequests', payload: data.nftMintRequests });
        }
      }
    },
  );

  useEffect(() => {
    if (token && state.status && Number.isInteger(state.page)) {
      refetch();
    }

    if (!token) {
      navigate({ to: ROUTES.auth, replace: true });
    }
  }, [state.page, state.status, token]);

  return (
    <div className={styles.homePage}>
      <Header />
      <StatusOptions />
      {data?.nftMintRequests ? <NftList nftMintRequests={data.nftMintRequests} /> : (
        <div className={styles.noData}>
          <h1>No data available</h1>
        </div>
      )}
    </div>
  );
};
