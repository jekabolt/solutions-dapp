import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { ReactLocationDevtools } from '@tanstack/react-location-devtools';
<<<<<<< HEAD

=======
>>>>>>> 910ce0f57929e011abf7ac9af44dd93f7c3ef56c
import {
  Outlet,
  ReactLocation,
  Router,
  Route,
  DefaultGenerics,
} from "@tanstack/react-location";
<<<<<<< HEAD

import { ROUTES } from 'constants/routes';
import Test from 'pages/Test';
import Test2 from 'pages/Test2';
import { Sidebar } from 'components/Sidebar';
import 'styles/global.scss';
=======
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ContextProvider } from 'context';

;import { ROUTES } from 'constants/routes';
import { NftListPage } from 'pages/NftListPage';
import Test2 from 'pages/Test2';
// import { Sidebar } from 'components/Sidebar';
import { LoginBlock } from 'components/LoginBlock';
import 'styles/global.scss';
import styles from 'styles/index.module.scss';
>>>>>>> 910ce0f57929e011abf7ac9af44dd93f7c3ef56c

const container = document.getElementById('root');
const root = createRoot(container!);

<<<<<<< HEAD
const location = new ReactLocation();
const routes: Route<DefaultGenerics>[] = [
  { path: ROUTES.home, element: <Test /> },
  { path: ROUTES.nftMintRequest, element: <Test2 /> },
=======
const queryClient = new QueryClient();
const location = new ReactLocation();
const routes: Route<DefaultGenerics>[] = [
  { path: ROUTES.home, element: <NftListPage /> },
  { path: ROUTES.nftRequests, element: <Test2 /> },
>>>>>>> 910ce0f57929e011abf7ac9af44dd93f7c3ef56c
];

root.render(
  <StrictMode>
<<<<<<< HEAD
    <Router location={location} routes={routes}>
      <header>
        <h1>User data if loggined, some cool content place, slider with nfts we sell etc</h1>
      </header>
      <section className="main-view">
        <div className="sidebar">
          <Sidebar />
        </div>
        {/* find better className */}
        <div className="route-page">
          <Outlet />
        </div>
      </section>
      <footer>
        Copyright © 2022 - {new Date().getFullYear()}. All Rights Reserved
      </footer>
      {process.env.NODE_ENV === 'development' ? <ReactLocationDevtools /> : null}
    </Router>
=======
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
          {false && process.env.NODE_ENV === 'development'
            ? <div style={{ padding: 0 }}><ReactLocationDevtools /></div>
            : null}
        </Router>
      </QueryClientProvider>
    </ContextProvider>
>>>>>>> 910ce0f57929e011abf7ac9af44dd93f7c3ef56c
  </StrictMode>
);
