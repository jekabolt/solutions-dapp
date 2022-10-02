import { FC, useEffect, useState } from 'react';
import type { AxiosError } from 'axios';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from '@tanstack/react-location';

import { Status as StatusType } from 'api/proto-http/nft';
import { getMintRequestsPagedList, QUERIES, NFTMintRequestListArray } from 'api';
import { ROUTES } from 'constants/routes';
import { AUTH_LOCAL_STORAGE_KEY, Status } from 'constants/values';
import { Header } from 'components/Header';
import { StatusOptions } from 'components/StatusOptions';
import { NftList } from 'components/NftList';

import styles from 'styles/HomePage.module.scss';

export const Home: FC = () => {
  const token = localStorage.getItem(AUTH_LOCAL_STORAGE_KEY) || '';
  const [activeStatus, setActiveStatus] = useState(Status.Any as StatusType);
  const page = 2;
  const navigate = useNavigate();
  // const { data } = useQuery<NFTMintRequestListArray, AxiosError>(
  //   [QUERIES.getNftRequests],
  //   () => getMintRequestsPagedList(token, { status: activeStatus, page }),
  //   { enabled: false, retry: false },
  // );

  // console.log(activeStatus);

  useEffect(() => {
    if (!token) {
      navigate({ to: ROUTES.auth, replace: true });
    } else {
      console.log('should fetch data');
      // refetch();
    }
  }, [token]);

  return (
    <div className={styles.homePage}>
      <Header />
      <StatusOptions activeStatus={activeStatus} setActiveStatus={setActiveStatus} />
      <NftList />
    </div>
  );
};
