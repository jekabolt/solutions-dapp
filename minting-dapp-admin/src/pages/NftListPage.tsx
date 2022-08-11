import { FC, Fragment } from 'react';
import { useQuery } from '@tanstack/react-query';

import { getNftRequests, QUERIES } from 'api';
import { NftPreview } from 'components/NftPreview';

import styles from 'styles/nft-list-page.module.scss';

export const NftListPage: FC = () => {
  const { data } = useQuery([QUERIES.getNftRequests], () =>
    getNftRequests(localStorage.getItem('authToken') || '')
  );

  return (
    <div className={styles.container}>
      <div className={styles.pageHeader}>
        <h3>some sort of data</h3>
      </div>
      <div className={styles.nftList}>
        <div style={{ backgroundColor: localStorage.getItem('authToken') ? 'lightgreen' : 'red' }}>auth</div>
        <br />
        <br />
        <div style={{ backgroundColor: data?.nftMintRequests ? 'lightgreen' : 'red' }}>
          getNftRequests
        </div>
        {data?.nftMintRequests &&
          data.nftMintRequests.map((nftRequest: any, index: number) => (
            // change key
            <Fragment key={index}>
              <NftPreview />
            </Fragment>
          ))}
      </div>
    </div>
  );
};
