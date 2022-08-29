import { FC, useEffect } from 'react';
import type { AxiosError } from 'axios';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from '@tanstack/react-location';

import { getNftRequests, QUERIES, NFTMintRequestListArray } from 'api';
import { ROUTES } from 'constants/routes';
import { AUTH_LOCAL_STORAGE_KEY } from 'constants/values';

import styles from 'styles/HomePage.module.scss';

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

export const Home: FC = () => {
  const token = localStorage.getItem(AUTH_LOCAL_STORAGE_KEY);
  const navigate = useNavigate();
  const { data, error, refetch } = useQuery<NFTMintRequestListArray, AxiosError>(
    [QUERIES.getNftRequests],
    () => getNftRequests(token || ''),
    { enabled: false, retry: false },
  );


  useEffect(() => {
    if (!token || error?.response?.status === 401) {
      navigate({ to: ROUTES.auth });
    } else {
      refetch();
    }
  }, [token, error?.response?.status]);

  return (
    <div className={styles.homePage}>
      home page
      <br />
      {data?.nftMintRequests?.map(() => (
        'single nft mint request'
      ))}
      <button onClick={() => localStorage.setItem(AUTH_LOCAL_STORAGE_KEY, '')}>clear localst</button>
    </div>
  );
};
