import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { ReactLocationDevtools } from '@tanstack/react-location-devtools';

import {
  Link,
  Outlet,
  ReactLocation,
  Router,
  Route,
  DefaultGenerics,
} from "@tanstack/react-location";

import { ROUTES } from './constants';
import Test from './pages/Test';
import Test2 from './pages/Test2';

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
      <Link to={ROUTES.home} activeOptions={{ exact: true }}>/home</Link>
      <br />
      <Link to={ROUTES.nftMintRequest} activeOptions={{ exact: true }}>/nft-mint-request</Link>
      <Outlet />
      {process.env.NODE_ENV === 'development' ? <ReactLocationDevtools /> : null}
    </Router>
  </StrictMode>
);
