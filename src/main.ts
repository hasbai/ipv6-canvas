import { createApp } from 'vue'
import App from '@/App.vue'
import router from './plugins/router'
import { store, key } from './plugins/store'
import eventBus from 'vue3-eventbus'

import '@/css/index.css'

const app = createApp(App)

app.use(eventBus)
app.use(store, key)
app.use(router)
app.mount('#app')
