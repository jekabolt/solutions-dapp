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

import { ROUTES } from 'constants/routes';
import { NftListPage } from 'pages/NftListPage';
import Test2 from 'pages/Test2';
// import { Sidebar } from 'components/Sidebar';
import 'styles/global.scss';
import styles from 'styles/index.module.scss';

const container = document.getElementById('root');
const root = createRoot(container!);

const location = new ReactLocation();
const routes: Route<DefaultGenerics>[] = [
  { path: ROUTES.home, element: <NftListPage /> },
  { path: ROUTES.nftMintRequest, element: <Test2 /> },
];

root.render(
  <StrictMode>
    <Router location={location} routes={routes}>
      <header className={styles.header}>
        <h1>User data if loggined, some cool content place</h1>
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
  </StrictMode>
);
