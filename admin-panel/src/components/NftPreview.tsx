import type { FC } from 'react';

import styles from 'styles/nft-preview.module.scss';

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface INftPreviewProps {}

export const NftPreview: FC<INftPreviewProps> = () => {
  return (
    <div className={styles.nftPreview}>
      nft preview image
      <br />
      description
      <br />
      button
    </div>
  );
};
