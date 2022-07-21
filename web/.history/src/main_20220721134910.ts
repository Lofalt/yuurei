import { createApp } from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import V3waterfall from 'v3-waterfall'


const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.use(router).use(V3waterfall).mount("#app");