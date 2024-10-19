import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { App } from './App';
import { MarketProvider } from './components/MarketContext/MarketProvider';
import './main.scss';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <MarketProvider>
      <App />
    </MarketProvider>
  </StrictMode>
);
