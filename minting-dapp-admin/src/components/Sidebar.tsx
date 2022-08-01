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
          <Link to={ROUTES.nftRequests} activeOptions={{ exact: true }}>/nft-requests</Link>
        </li>
      </ul>
    </div>
  );
};
