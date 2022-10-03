import { useState, ChangeEvent, useEffect, MouseEvent } from 'react';
import { useMutation } from '@tanstack/react-query';
import { useNavigate } from '@tanstack/react-location';

import { login } from 'api';
import { ROUTES } from 'constants/routes';
import { AUTH_LOCAL_STORAGE_KEY } from 'constants/values';

import styles from 'styles/AuthPage.module.scss';

export const Auth = () => {
  const [password, setPassword] = useState('');
  const { mutate, data, error } = useMutation(login);
  const navigate = useNavigate();

  const handlePasswordSubmit = (e: MouseEvent) => {
    e.preventDefault();
    mutate(password);
  };

  const handlePasswordChange = (e: ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  useEffect(() => {
    if (data?.authToken) {
      localStorage.setItem(AUTH_LOCAL_STORAGE_KEY, data.authToken);
      navigate({ to: ROUTES.home });
    }
  }, [data?.authToken]);

  return (
    <div className={styles.authPage}>
      <input
        onChange={handlePasswordChange}
        type="password"
        className={styles.passwordInput}
      />
      {!!error && <p className={styles.errorMessage}>enter correct password</p>}
      <button className={styles.passwordSubmit} type="submit" onClick={handlePasswordSubmit}>LOG IN</button>
    </div>
  );
};
