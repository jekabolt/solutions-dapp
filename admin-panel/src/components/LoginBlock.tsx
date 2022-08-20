import { FC, useState, ChangeEvent, FormEvent, useEffect } from 'react';
import { useMutation } from '@tanstack/react-query';

import { login } from 'api';
import styles from 'styles/login-block.module.scss';

export const LoginBlock: FC = () => {
  const [isModalOpen, setModalVisability] = useState(false);
  const [password, setPassword] = useState('');
  const { mutate, data } = useMutation(login);

  useEffect(() => {
    if (data?.authToken) {
      localStorage.setItem('authToken', data.authToken);
    }
  }, [data?.authToken]);

  const toggleModal = () => {
    setModalVisability((v) => !v);
  };

  const handlePasswordSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    mutate(password);
    toggleModal();
  };

  const handlePasswordChange = (e: ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  return (
    <div className={styles.loginBlock}>
      <button onClick={toggleModal}>Login</button>
      {isModalOpen && (
        <>
          <div className={styles.loginModal}>
            <div className={styles.close} onClick={toggleModal}>
              close modal
            </div>
            <form onSubmit={handlePasswordSubmit}>
              <input onChange={handlePasswordChange} type='password' placeholder='password' />
              <button type='submit'>login</button>
            </form>
          </div>
          <div className={styles.overlay}></div>
        </>
      )}
    </div>
  );
};
