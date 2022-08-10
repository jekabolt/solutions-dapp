import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { ReactLocationDevtools } from '@tanstack/react-location-devtools';
import {
  Outlet,
  ReactLocation,
  Router,
  Route,
  DefaultGenerics,
} from "@tanstack/react-location";
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ContextProvider } from 'context';

import { ROUTES } from 'constants/routes';
import { NftListPage } from 'pages/NftListPage';
import Test2 from 'pages/Test2';
// import { Sidebar } from 'components/Sidebar';
import { LoginBlock } from 'components/LoginBlock';
import 'styles/global.scss';
import styles from 'styles/index.module.scss';

const container = document.getElementById('root');
const root = createRoot(container!);

const queryClient = new QueryClient();
const location = new ReactLocation();
const routes: Route<DefaultGenerics>[] = [
  { path: ROUTES.home, element: <NftListPage /> },
  { path: ROUTES.nftRequests, element: <Test2 /> },
];

root.render(
  <StrictMode>
    <ContextProvider>
      <QueryClientProvider client={queryClient}>
        <Router location={location} routes={routes}>
          <header className={styles.header}>
            <span>User data if loggined, some cool content place</span>
            <LoginBlock />
          </header>
          <section className={styles.mainView}>
            {/* think of page layout */}
            {/* <div className="sidebar"> */}
            {/* <Sidebar /> */}
            {/* </div> */}
            {/* find better className */}
            {/* <div className="route-page"> */}
            <Outlet />
            {/* </div> */}
          </section>
          <footer className={styles.footer}>
            Copyright © 2022 - {new Date().getFullYear()}. All Rights Reserved
          </footer>
          {process.env.NODE_ENV === 'development'
            ? <div style={{ padding: 0 }}><ReactLocationDevtools /></div>
            : null}
        </Router>
      </QueryClientProvider>
    </ContextProvider>
  </StrictMode>
);