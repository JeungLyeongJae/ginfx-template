import { createApp } from 'vue';
import Antd from 'ant-design-vue';
import App from './App.vue';
import 'ant-design-vue/dist/reset.css';
import BpmnVueTemp from "./components/app/bpmn";
import './assets/themes/scrollbar.css'
import router from "./router.ts"; // 引入全局样式

const app = createApp(App);

app.use(Antd)
    .use(BpmnVueTemp)
    .use(router)
    .mount('#app');