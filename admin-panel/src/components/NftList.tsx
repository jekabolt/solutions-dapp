import { useContext } from 'react';
import { useNavigate } from '@tanstack/react-location';

import type { NFTMintRequestWithStatus } from 'api';
import { Context } from 'context';
import { ROUTES } from 'constants/routes';
import { STATUS_COLORS, Status } from 'constants/values';
import styles from 'styles/NftList.module.scss';

const NtfListItem = ({
  nftRequestWithStatus: { nftMintRequest, sampleImages, status },
  handleClick
}: {
    nftRequestWithStatus: NFTMintRequestWithStatus,
    handleClick: () => void,
}) => (
  <div className={styles.nftListItem} onClick={handleClick}>
    <div className={styles.imgContainer}>
      {sampleImages && sampleImages.length > 0 &&
        <img src={sampleImages[0].compressed} alt="nft preview" />
      }
    </div>
    <p className={styles.status}>
      <span className={styles.statusIndicator} style={{
        // @ts-ignore // status should be as text not as number
        backgroundColor: STATUS_COLORS[Object.values(Status)[status]],
      }} />
      {/* @ts-ignore */}
      {Object.keys(Status)[status]}
    </p>
    <span className={styles.itemId}>{nftMintRequest?.id}</span>
  </div>
);

export const NftList = ({ nftMintRequests }: { nftMintRequests: NFTMintRequestWithStatus[] }) => {
  const navigate = useNavigate();
  const { state, dispatch } = useContext(Context);

  const handleNftClick = (value: NFTMintRequestWithStatus) => {
    dispatch({ type: 'setActiveNftMintRequest', payload: value });
    navigate({ to: ROUTES.nft, search: { id: value.nftMintRequest?.id } });
  };

  // todo: fix error
  const handleLoadMore = () => {
    dispatch({ type: 'setPage', payload: state.page + 1 });
  };

  const arr = Array(15).fill(nftMintRequests[0])

  return (
    <div className={styles.nftListContainer}>
      <div className={styles.nftList}>
        {arr.map((nftMintRequest) => (
          <NtfListItem
            nftRequestWithStatus={nftMintRequest}
            handleClick={() => handleNftClick(nftMintRequest)}
          />
        ))}
      </div>
      <button
        className={styles.loadMore}
        onClick={handleLoadMore}
      >
        LOAD MORE
      </button>
    </div>
  );
}
