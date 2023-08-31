import { createApp } from 'vue'
import { createPinia } from 'pinia';
import 'normalize.css'
import 'element-plus/theme-chalk/dark/css-vars.css';
import App from './App.vue'
import router from "./router"

createApp(App).use(router).use(createPinia()).mount('#app')
