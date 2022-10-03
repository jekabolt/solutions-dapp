import { useNavigate } from '@tanstack/react-location';

import { ROUTES } from 'constants/routes';
import { AUTH_LOCAL_STORAGE_KEY } from 'constants/values';
import styles from 'styles/Header.module.scss';

export const Header = () => {
  const navigate = useNavigate();

  return (
    <header className={styles.header}>
      <div className={styles.left}>
        <div className={styles.circle}>
          {/* possibly ;logo */}
        </div>
        {/* TBD */}
      </div>
      <div className={styles.right}>
        {/* <span>TBD</span> */}
        <button
          className={styles.logOut}
          onClick={() => {
            localStorage.setItem(AUTH_LOCAL_STORAGE_KEY, '');
            navigate({ to: ROUTES.home })
          }}
        >
          LOG OUT
        </button>
      </div>
    </header>
  )
};
