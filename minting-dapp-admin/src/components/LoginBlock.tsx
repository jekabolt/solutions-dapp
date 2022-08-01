import { FC, useState } from 'react';
import styles from 'styles/login-block.module.scss';

export const LoginBlock: FC = () => {
  const [isModalOpen, setModalVisability] = useState(false);
  const handleLoginClick = () => {
    // implement login, auth token, cookies, etc.
    setModalVisability((v) => !v);
  };

  const handleModalClose = () => {
    setModalVisability(false);
  };

  return (
    <div className={styles.loginBlock}>
      <button onClick={handleLoginClick}>Login</button>
      {isModalOpen && (
        <>
          <div className={styles.loginModal}>
            <div className={styles.close} onClick={handleModalClose}>close modal</div>
            <input type="login" placeholder="Login" />
            <input type="password" placeholder="password" />
          </div>
          <div className={styles.overlay}></div>
        </>
      )}
    </div>
  );
};
