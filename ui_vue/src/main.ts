import {createApp} from 'vue';
import Antd from 'ant-design-vue';
import App from './App.vue';
import 'ant-design-vue/dist/reset.css';
import './assets/themes/scrollbar.css';
import router from "./router.ts"; // 引入全局样式
import "../public/utils/update";

const app = createApp(App);

app.use(Antd)
    .use(router)
    .mount('#app');