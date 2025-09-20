import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { Dapp } from './Dapp.tsx';
import 'bootstrap/dist/css/bootstrap.css';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Dapp />
  </StrictMode>
);
