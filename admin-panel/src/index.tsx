import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import {
  Outlet,
  ReactLocation,
  Router,
  Route,
  DefaultGenerics,
} from '@tanstack/react-location';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import { ContextProvider } from 'context';
import { ROUTES } from 'constants/routes';
import { Home } from 'pages/Home';
import { Auth } from 'pages/Auth';

import 'styles/global.scss';
import styles from 'styles/index.module.scss';

const container = document.getElementById('root') as HTMLElement;
const root = createRoot(container);

const queryClient = new QueryClient();
const location = new ReactLocation();
const routes: Route<DefaultGenerics>[] = [
  { path: ROUTES.home, element: <Home /> },
  { path: ROUTES.auth, element: <Auth /> },
];

root.render(
  <StrictMode>
    <ContextProvider>
      <QueryClientProvider client={queryClient}>
        <Router location={location} routes={routes}>
          <section className={styles.mainView}>
            <Outlet />
          </section>
          <footer className={styles.footer}>
            Copyright © 2022 - {new Date().getFullYear()}. All Rights Reserved
          </footer>
        </Router>
      </QueryClientProvider>
    </ContextProvider>
  </StrictMode>,
);
