import {createApp, onMounted} from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
// import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import V3waterfall from 'v3-waterfall'
// import "./style.css
import 'v3-waterfall/style.css'
import {createPinia} from "pinia";
import 'vue-waterfall-plugin-next/style.css'
import axios from "./request/index"
import hljs from 'highlight.js';
// import 'highlight.js/11.6.0/styles/default.min.css';


const pinia = createPinia()
const app = createApp(App)

app.use(router)
    .use(V3waterfall)
    .use(pinia);

app.directive('highlight', el => {
    let blocks = el.querySelectorAll('pre code');
    Array.prototype.forEach.call(blocks, hljs.highlightBlock);
});
app.mount("#app");
app.config.globalProperties.$axios = axios;