import type { FC } from 'react';
import { Link } from '@tanstack/react-location';

import { ROUTES } from 'constants/routes';
import styles from 'styles/sidebar.module.scss';

export const Sidebar: FC = () => {
  return (
    <div className={styles.sidebar}>
      <ul className="navigation">
        <li>
          <Link to={ROUTES.home} activeOptions={{ exact: true }}>/home</Link>
        </li>
        <li>
<<<<<<< HEAD
          <Link to={ROUTES.nftMintRequest} activeOptions={{ exact: true }}>/nft-mint-request</Link>
=======
          <Link to={ROUTES.nftRequests} activeOptions={{ exact: true }}>/nft-requests</Link>
>>>>>>> 910ce0f57929e011abf7ac9af44dd93f7c3ef56c
        </li>
      </ul>
    </div>
  );
};
