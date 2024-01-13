import {createApp} from 'vue'
import App from './App.vue'
import { router } from './router'
import PrimveVue from 'primevue/config';
import 'primevue/resources/themes/saga-blue/theme.css';
import './style.css';

const app = createApp(App)
app.use(router)
app.use(PrimveVue)
app.mount('#app')
