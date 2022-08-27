import { useState, ChangeEvent, useEffect } from 'react';
import { useMutation } from '@tanstack/react-query';
import { useNavigate } from '@tanstack/react-location';

import { login } from 'api';
import { ROUTES } from 'constants/routes';
import styles from 'styles/AuthPage.module.scss';

export const Auth = () => {
  const [password, setPassword] = useState('');
  const { mutate, data, error } = useMutation(login);
  const navigate = useNavigate();

  const handlePasswordSubmit = (e: any) => {
    e.preventDefault();
    mutate(password);
  };

  const handlePasswordChange = (e: ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  useEffect(() => {
    if (data?.authToken) {
      localStorage.setItem('authToken', data.authToken);
      // should we replace history or just push to it
      navigate({ to: ROUTES.home, replace: true });
    }
  }, [data?.authToken]);

  return (
    <div className={styles.authPage}>
      <h1>auth page</h1>
      <input
        onChange={handlePasswordChange}
        type="password"
        className={styles.passwordInput}
        placeholder="enter pass"
      />
      {!!error && <p className={styles.errorMessage}>enter correct password</p>}
      <br />
      <button className={styles.passwordSubmit} type="submit" onClick={handlePasswordSubmit}>Sign in</button>
    </div>
  );
};
