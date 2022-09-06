import styles from './Header.module.scss';

export const Header = () => {
  return (
    <header className={styles.header}>
      <div className={styles.left}>
        <div className={styles.circle}>
          {/* possibly ;logo */}
        </div>
        TBD
      </div>
      <div className={styles.right}>
        <span>TBD</span>
        <button className={styles.logOut}>
          LOG OUT
        </button>
      </div>
    </header>
  )
};
