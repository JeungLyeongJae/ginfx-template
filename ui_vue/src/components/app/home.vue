<script lang="ts">

import {theme} from "ant-design-vue";
import {defineComponent, ref} from "vue";
import {
  DesktopOutlined,
  FileOutlined,
  LaptopOutlined,
  NotificationOutlined,
  PieChartOutlined,
  TeamOutlined,
  UserOutlined
} from '@ant-design/icons-vue';

const collapsed = ref<boolean>(false);
const selectedKeys = ref<string[]>(['1']);
const checked = ref<boolean>(false);
let usingTheme = ref([theme.defaultAlgorithm])


export default defineComponent({
  name: 'home',
  components: {
    UserOutlined,
    LaptopOutlined,
    NotificationOutlined,
    PieChartOutlined,
    DesktopOutlined,
    TeamOutlined,
    FileOutlined,
  },
  computed: {
    theme() {
      return theme
    },
  },

  setup() {
    const toggleColorTheme = () => {
      usingTheme.value[0] == theme.defaultAlgorithm ? usingTheme.value[0] = theme.darkAlgorithm : usingTheme.value[0] = theme.defaultAlgorithm
    };
    const toggleTheme = () => {
      usingTheme.value.length == 1? usingTheme.value.push(theme.compactAlgorithm): usingTheme.value.pop()
    };
    return {
      collapsed,
      selectedKeys,
      usingTheme,
      checked,
      toggleColorTheme,
      toggleTheme
    };
  }
})

</script>


<template>
  <a-config-provider :theme="{token: {borderRadius: 16}, algorithm: usingTheme}">
    <a-layout style="min-height: 100vh">
      <a-layout-sider v-model:collapsed="collapsed" collapsible >
        <div class="logo"/>
        <a-menu v-model:selectedKeys="selectedKeys" :style="{ height: '100%', borderRight: 0 }" mode="inline">
          <a-menu-item key="1">
            <pie-chart-outlined/>
            <span>Option 1</span>
          </a-menu-item>
          <a-menu-item key="2">
            <desktop-outlined/>
            <span>Option 2</span>
          </a-menu-item>
          <a-sub-menu key="sub1">
            <template #title>
            <span>
              <user-outlined/>
              <span>User</span>
            </span>
            </template>
            <a-menu-item key="3">Tom</a-menu-item>
            <a-menu-item key="4">Bill</a-menu-item>
            <a-menu-item key="5">Alex</a-menu-item>
          </a-sub-menu>
          <a-sub-menu key="sub2">
            <template #title>
            <span>
              <team-outlined/>
              <span>Team</span>
            </span>
            </template>
            <a-menu-item key="6">Team 1</a-menu-item>
            <a-menu-item key="8">Team 2</a-menu-item>
          </a-sub-menu>
          <a-menu-item key="9">
            <file-outlined/>
            <span>File</span>
          </a-menu-item>
        </a-menu>
      </a-layout-sider>
      <a-layout>
        <a-layout-header :style="{ border: '20px' }"/>
        <a-layout-content style="margin: 0 16px">
          <a-breadcrumb style="margin: 16px 0">
            <a-breadcrumb-item>User</a-breadcrumb-item>
            <a-breadcrumb-item>Bill</a-breadcrumb-item>
          </a-breadcrumb>
          <div :style="{ padding: '24px',  minHeight: '360px' }">
            <a-button type="primary" @click="toggleColorTheme()">切换</a-button>
            <br/>
            <br/>
            <a-switch v-model:checked="checked" @click="toggleTheme()"/>
          </div>
        </a-layout-content>
        <a-layout-footer style="text-align: center">
          Design ©2024 Created by JeungNyeongJae
        </a-layout-footer>
      </a-layout>
    </a-layout>
  </a-config-provider>
</template>

<style scoped>
.logo {
  height: 32px;
  margin: 16px;
  background: rgba(255, 255, 255, 0.3);
}
.site-layout .site-layout-background {
  background: #fff;
}
[data-theme='dark'] .site-layout .site-layout-background {
  background: #141414;
}
</style>
