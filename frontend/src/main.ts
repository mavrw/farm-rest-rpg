import './assets/styles/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import { createPinia } from 'pinia'
import router from './router'
import { useAuthStore } from './stores/authStore'

const app = createApp(App)

const pinia = createPinia()
app.use(pinia)
app.use(router)

const auth = useAuthStore();
// TODO: implement a listener for window.load before mounting the app
auth.fetchCurrentUser().finally(() => {
    app.mount('#app')
});