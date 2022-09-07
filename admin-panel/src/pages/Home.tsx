import { FC, useEffect } from 'react';
import type { AxiosError } from 'axios';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from '@tanstack/react-location';

import { getNftRequests, QUERIES, NFTMintRequestListArray } from 'api';
import { ROUTES } from 'constants/routes';
import { AUTH_LOCAL_STORAGE_KEY } from 'constants/values';
import { Header } from 'components/Header';
import { StatusOptions } from 'components/StatusOptions';
import { NftList } from 'components/NftList';

import styles from 'styles/HomePage.module.scss';

export const Home: FC = () => {
  const token = localStorage.getItem(AUTH_LOCAL_STORAGE_KEY) || '';
  const navigate = useNavigate();
  const { data, error, refetch } = useQuery<NFTMintRequestListArray, AxiosError>(
    [QUERIES.getNftRequests],
    () => getNftRequests(token),
    { enabled: false, retry: false },
  );

  useEffect(() => {
    if (!token || error?.response?.status === 401) {
      navigate({ to: ROUTES.auth, replace: true });
    } else {
      // if server stops infinity call occur
      // refetch();
    }
  }, [token, error?.response?.status]);

  return (
    <div className={styles.homePage}>
      <Header />
      <StatusOptions />
      <NftList />
    </div>
  );
};
