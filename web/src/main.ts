import { createApp } from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
// import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import V3waterfall from 'v3-waterfall'
// import "./style.css
import 'v3-waterfall/style.css'
import { createPinia } from "pinia";
import 'vue-waterfall-plugin-next/style.css'

const pinia = createPinia()
const app = createApp(App)

app.use(router)
    .use(V3waterfall)
    .use(pinia);
app.mount("#app");