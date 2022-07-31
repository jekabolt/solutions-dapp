import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';

import Test from './pages/Test';
import './global.scss';

const container = document.getElementById('root');

const root = createRoot(container!);
root.render(
  <StrictMode>
    <div className="global">
      <Test />
    </div>
  </StrictMode>
);
