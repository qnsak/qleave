import { createApp } from 'vue'
import './style.css'
import './index.css'
import App from './App.vue'
import { createPinia } from 'pinia'
import piniaPersist from 'pinia-plugin-persist'
import { router } from './router'

import { useRegisterSW } from 'virtual:pwa-register/vue';
import { registerLayouts } from './layouts/register';

useRegisterSW();

const app = createApp(App);

const pinia = createPinia();
pinia.use(piniaPersist);

app.use(pinia);
app.use(router);

app.mount('#app');
registerLayouts(app);