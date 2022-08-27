import { FC, useEffect } from 'react';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from '@tanstack/react-location';

import { getNftRequests, QUERIES } from 'api';
import { ROUTES } from 'constants/routes';
import { AUTH_LOCAL_STORAGE_KEY } from 'constants/values';
import { NftPreview } from 'components/NftPreview';

import styles from 'styles/nft-list-page.module.scss';

export const Home: FC = () => {
  const token = localStorage.getItem(AUTH_LOCAL_STORAGE_KEY);
  const navigate = useNavigate();
  // fix request on every render...
  // add types..import from proto and use generics
  const { data, error, refetch } = useQuery([QUERIES.getNftRequests], () =>
    getNftRequests(token || ''),
    { enabled: false, retry: false },
  );

  console.log(data);

  useEffect(() => {
    // redirect to auth if there is no token in LS or token is invalid
    //@ts-ignore
    if (!token || error?.response?.status === 401) {
      navigate({ to: ROUTES.auth });
    } else {
      refetch();
    }
    // @ts-ignore
  }, [token, error?.response?.status]);

  return (
    <div className={styles.container}>
      home page

      <button onClick={() => localStorage.setItem(AUTH_LOCAL_STORAGE_KEY, '')}>clear localst</button>
    </div>
  );
};
