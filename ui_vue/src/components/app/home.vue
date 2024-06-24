<script lang="ts" setup>

import {theme} from "ant-design-vue";
import {onMounted, ref} from "vue";
import {
  BgColorsOutlined,
  CompressOutlined,
  DesktopOutlined,
  FileOutlined,
  PieChartOutlined,
  TeamOutlined,
  ToolOutlined,
  UserOutlined
} from '@ant-design/icons-vue';
import User from "./admin/user/user.vue";

const collapsed = ref<boolean>(false);
const selectedKeys = ref<string[]>(['1']);

let toggleColorThemeButtonType = ref<string>('default')
let toggleThemeButtonType = ref<string>('default')
let usingTheme = ref([theme.defaultAlgorithm])


onMounted(() => {
  const currentHour = new Date().getHours();
  if (sessionStorage.getItem('toggleColorThemeButtonType') == 'default' || (sessionStorage.getItem('toggleColorThemeButtonType') == null && currentHour >= 6 && currentHour < 18)) {
    usingTheme.value[0] = theme.defaultAlgorithm
    toggleColorThemeButtonType.value = 'default'
  } else {
    usingTheme.value[0] = theme.darkAlgorithm
    toggleColorThemeButtonType.value = 'primary'
  }

  if (sessionStorage.getItem('toggleThemeButtonType') == 'default' || sessionStorage.getItem('toggleThemeButtonType') == null) {
    toggleThemeButtonType.value = 'default'
  } else {
    usingTheme.value.push(theme.compactAlgorithm)
    toggleThemeButtonType.value = 'primary'
  }
})

const toggleColorTheme = () => {

  if (toggleColorThemeButtonType.value == 'default') {
    usingTheme.value[0] = theme.darkAlgorithm
    toggleColorThemeButtonType.value = 'primary'
  } else {
    usingTheme.value[0] = theme.defaultAlgorithm
    toggleColorThemeButtonType.value = 'default'
  }
  sessionStorage.setItem('toggleColorThemeButtonType', toggleColorThemeButtonType.value)
  // const isLight = usingTheme.value[0] == theme.defaultAlgorithm;
  // const _head = document.getElementsByTagName('head')[0] || null;
  // const _styles = Array.from(document.getElementsByTagName('link'));
  // console.log('_head:', _head)
  // console.log('_styles:', _styles)
  // _styles.forEach((style: any, index) => {
  //   if (index === 2 && _styles.length > 1) {
  //     _styles[index].parentElement!.removeChild(_styles[index]);
  //   }
  // });
  // if (_head) {
  //   const link = document.createElement('link');
  //   link.setAttribute('rel', 'stylesheet');
  //   link.setAttribute('name', 'theme');
  //   link.setAttribute('href', `/theme/bpmn/${isLight ? 'light' : 'dark'}.css`);
  //   _head.appendChild(link);
  // }
};
const toggleTheme = () => {
  if (toggleThemeButtonType.value == 'default') {
    usingTheme.value.push(theme.compactAlgorithm)
    toggleThemeButtonType.value = 'primary'
  } else {
    usingTheme.value.pop()
    toggleThemeButtonType.value = 'default'
  }
  sessionStorage.setItem('toggleThemeButtonType', toggleThemeButtonType.value)
};

function selectElement(info: any) {
  console.log(info);
}


</script>

<template>
  <a-config-provider :theme="{token: {borderRadius: 16}, algorithm: usingTheme}">
    <a-layout>
      <!--      :style="{ position: 'fixed', zIndex: 1, width: '100%' }"  固定页头-->
      <a-layout-header>
        <div class="logo" style="font-family: 'Microsoft YaHei', sans-serif;">JeungNyeongJae</div>
      </a-layout-header>

      <a-layout style="min-height: 100vh">
        <a-layout-sider v-model:collapsed="collapsed" width="200" collapsible breakpoint="lg">
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
            <a-sub-menu key="账号管理">
              <template #title>
            <span>
              <team-outlined/>
              <span>账号管理</span>
            </span>
              </template>
              <a-menu-item key="角色列表">角色列表</a-menu-item>
              <a-menu-item key="人员信息">人员信息</a-menu-item>
            </a-sub-menu>
            <a-menu-item key="bpmn">
              <file-outlined/>
              <span>工作流</span>
            </a-menu-item>
          </a-menu>
        </a-layout-sider>

        <a-layout style="padding: 0 24px 24px">
          <a-breadcrumb style="margin: 16px 0">
            <a-breadcrumb-item>{{ selectedKeys[0] }}</a-breadcrumb-item>
            <a-breadcrumb-item>{{ selectedKeys[0] }}</a-breadcrumb-item>
          </a-breadcrumb>

          <a-layout-content style="margin: 0 16px; ">
            <div class="box" id="bpmn" v-if="selectedKeys[0] === 'bpmn'">
              <Bpmn-Vue :bpmnID="'test'" ref="bpmnRef" @select:element="selectElement"/>
            </div>
            <div class="box" id="role_list" v-if="selectedKeys[0] === '人员信息'">
              <user/>
            </div>
          </a-layout-content>

          <a-layout-footer style="text-align: center">
            Design ©2024 Created by JeungNyeongJae
          </a-layout-footer>
        </a-layout>
      </a-layout>
    </a-layout>
  </a-config-provider>

  <a-float-button-group shape="circle" :style="{ right: '18px', position: 'fixed' }" trigger="hover">
    <template #icon>
      <ToolOutlined/>
    </template>
    <a-float-button :type="toggleColorThemeButtonType.valueOf()"
                    @click="toggleColorTheme()">
      <template #tooltip>
        <div>浅/深色模式</div>
      </template>
      <template #icon>
        <BgColorsOutlined/>
      </template>
    </a-float-button>
    <a-float-button :type="toggleThemeButtonType.valueOf()"
                    @click="toggleTheme()">
      <template #tooltip>
        <div>大/小模式</div>
      </template>
      <template #icon>
        <CompressOutlined/>
      </template>
    </a-float-button>
    <a-back-top tooltip="返回顶部" :visibility-height="0"/>
  </a-float-button-group>
</template>

<style scoped>
.logo {
  align-items: center;
  display: flex;
  width: 120px;
  height: 25px;
  margin: 15px 5px 5px -10px;
  border-radius: 10px;
  padding: 5px;
  font-size: 13px;
  background: rgba(255, 255, 255, 0.3);
}

.box {
  height: 95%;
  width: 98%;
}
</style>
