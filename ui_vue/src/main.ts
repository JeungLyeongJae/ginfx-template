import { createApp } from 'vue';
import Antd from 'ant-design-vue';
import App from './App.vue';
import 'ant-design-vue/dist/reset.css';
import BpmnVueTemp from "./components/app/bpmn";

const app = createApp(App);

app.use(Antd).use(BpmnVueTemp).mount('#app');