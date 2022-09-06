import { AUTH_LOCAL_STORAGE_KEY } from 'constants/values';
import styles from 'styles/Header.module.scss';

export const Header = () => {
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
            window.location.reload();
          }}
        >
          LOG OUT
        </button>
      </div>
    </header>
  )
};
