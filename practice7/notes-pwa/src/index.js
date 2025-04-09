// index.js
import { createRoot } from 'react-dom/client';
import App from './App';
import * as serviceWorkerRegistration from './serviceWorkerRegistration';
import './index.css';

const container = document.getElementById('root');
const root = createRoot(container);
root.render(<App />);

// Изменяем unregister() на register()
serviceWorkerRegistration.register();