import { Fragment } from 'react';
import { useNavigate } from '@tanstack/react-location';

import { ROUTES } from 'constants/routes';
import styles from 'styles/NftList.module.scss';

const PAGE_SIZE = 12
const arr = new Array(PAGE_SIZE).fill('null');

const NtfListItem = ({
  index,
  handleClick
}: {
  index: number,
  handleClick: (id: any) => void,
}) => (
  <div className={styles.nftListItem} onClick={() => handleClick(index)}>
    <span className={styles.statusIndicator} style={{
      backgroundColor: 'yellow',
    }} />
    <div className={styles.img} />
    {/* <img src="" alt="" /> */}
    <span className={styles.itemId}>№45675678</span>
  </div>
);

export const NftList = () => {
  const navigate = useNavigate();

  const handleNftClick = (nftId: any) => {
    navigate({ to: ROUTES.nft, search: { id: nftId } });
  };

  return (
    <div className={styles.nftListContainer}>
      <div className={styles.nftList}>
        {arr.map((_, index) => (
          <Fragment key={index}>
            <NtfListItem index={index} handleClick={handleNftClick} />
          </Fragment>
        ))}
      </div>
      <button className={styles.loadMore}>LOAD MORE</button>
    </div>
  );
}
