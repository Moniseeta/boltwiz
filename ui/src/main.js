import { createApp } from 'vue'
import {createRouter, createWebHashHistory} from 'vue-router'
import {createPinia} from 'pinia'

import App from './App.vue'
import { Quasar } from 'quasar'
import quasarUserOptions from './quasar-user-options'
import routes from "@/routes/routes";


const p = createPinia()
const app = createApp(App)
app.use(Quasar, quasarUserOptions);
app.use(p)

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})
app.use(router)

app.mount('#app')
