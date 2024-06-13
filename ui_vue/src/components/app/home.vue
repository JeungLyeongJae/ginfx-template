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
      usingTheme.value.length == 1 ? usingTheme.value.push(theme.compactAlgorithm) : usingTheme.value.pop()
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

    <a-layout>
<!--      :style="{ position: 'fixed', zIndex: 1, width: '100%' }"  固定页头-->
      <a-layout-header>
        <div class="logo"/>
      </a-layout-header>
      <a-layout style="min-height: 100vh">
        <a-layout-sider v-model:collapsed="collapsed" width="200" collapsible>
          <a-menu v-model:selectedKeys="selectedKeys" :style="{height: '100%', borderRight: 0 }" mode="inline">
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
        <a-layout style="padding: 0 24px 24px">
          <a-breadcrumb style="margin: 16px 0">
            <a-breadcrumb-item>User</a-breadcrumb-item>
            <a-breadcrumb-item>Bill</a-breadcrumb-item>
          </a-breadcrumb>
          <a-layout-content style="margin: 0 16px">

            <div :style="{ padding: '24px',  minHeight: '360px' }">
              <a-button type="primary" @click="toggleColorTheme()">浅/深色模式</a-button>
              <br/>
              <br/>
              大/小模式：
              <a-switch v-model:checked="checked" @click="toggleTheme()"/>
            </div>
          </a-layout-content>
          <a-layout-footer style="text-align: center">
            Design ©2024 Created by JeungNyeongJae
          </a-layout-footer>
        </a-layout>
      </a-layout>
    </a-layout>
  </a-config-provider>
</template>

<style scoped>
.logo {
  float: left;
  width: 120px;
  height: 31px;
  margin: 16px 24px 16px 0;
  background: rgba(255, 255, 255, 0.3);
}
</style>
