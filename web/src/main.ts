import { createApp } from 'vue'
import App from '/@/App.vue'
import router from '/@/router';
import store from "/@/store";
import { i18n } from '/@/i18n';

import ElementPlus from "element-plus"
import "element-plus/dist/index.css"
import 'element-plus/dist/index.css';
import '/@/theme/index.scss';

import "/@/assets/css/setting.css"
import "/@/assets/css/global.css"

const app = createApp(App);
app.use(router);
app.use(store);
app.use(ElementPlus, { i18n: i18n.global.t, size: "mini" });
app.use(i18n);

app.mount('#app')

export default app;
