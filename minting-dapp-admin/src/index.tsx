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
import Test from 'pages/Test';
import Test2 from 'pages/Test2';
import { Sidebar } from 'components/Sidebar';
import 'styles/global.scss';

const container = document.getElementById('root');
const root = createRoot(container!);

const location = new ReactLocation();
const routes: Route<DefaultGenerics>[] = [
  { path: ROUTES.home, element: <Test /> },
  { path: ROUTES.nftMintRequest, element: <Test2 /> },
];

root.render(
  <StrictMode>
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
  </StrictMode>
);
