import '@carbon/charts-react/styles.css';
import { createRoot } from 'react-dom/client';
import App from './App';
import './App.css';
import AppDataProvider from './components/AppDataProvider/AppDataProvider';
import './main.scss';
import '/node_modules/react-grid-layout/css/styles.css';
import '/node_modules/react-resizable/css/styles.css';

const container = document.getElementById('root');

const root = createRoot(container!);

root.render(
  <AppDataProvider>
    <App />
  </AppDataProvider>
);
