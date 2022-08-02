import { FC, Fragment } from 'react';
import { useQuery } from '@tanstack/react-query';

import { NftPreview } from 'components/NftPreview';
import styles from 'styles/nft-list-page.module.scss';

export const NftListPage: FC = () => {
  // const {} = useQuery(QUERIES.getNftRequest, );
  // get nft from backend 
  const nft = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

  

  return (
    <div className={styles.container}>
      <div className={styles.pageHeader}>
        <h3>some sort of data</h3>
      </div>
      <div className={styles.nftList}>
        {nft.map(({ }, index) => (
          // change key
          <Fragment key={index}>
            <NftPreview />
          </Fragment>
        ))}
      </div>
    </div>
  );
};
