<script  lang="ts">

import {theme} from "ant-design-vue";
import {defineComponent, ref} from "vue";

const checked = ref<boolean>(false);
let usingTheme = ref(theme.defaultAlgorithm)


export default defineComponent({
  name: 'home',
  computed: {
    theme() {
      return theme
    },
  },

  setup() {
    const toggleTheme = () => {
      usingTheme.value == theme.defaultAlgorithm? usingTheme.value = theme.darkAlgorithm : usingTheme.value = theme.defaultAlgorithm
    };
    return {
      usingTheme,
      checked,
      toggleTheme
    };
  }
})

</script>

<template>
  <a-config-provider :theme="{algorithm: usingTheme}">
    <a-layout>
      <a-layout-header class="header">
        <div class="logo" />
        <a-menu
            theme="dark"
            mode="horizontal"
            :style="{ lineHeight: '64px' }"
        >
          <a-menu-item key="1">nav 1</a-menu-item>
          <a-menu-item key="2">nav 2</a-menu-item>
          <a-menu-item key="3">nav 3</a-menu-item>
        </a-menu>
      </a-layout-header>
      <a-layout-content style="padding: 0 50px">
        <a-breadcrumb style="margin: 16px 0">
          <a-breadcrumb-item>Home</a-breadcrumb-item>
          <a-breadcrumb-item>List</a-breadcrumb-item>
          <a-breadcrumb-item>App</a-breadcrumb-item>
        </a-breadcrumb>
        <a-layout style="padding: 24px 0; background: #fff">
          <a-layout-sider width="200" style="background: #fff">
            <a-menu
                mode="inline"
                style="height: 100%"
            >
              <a-sub-menu key="sub1">
                <template #title>
                <span>
                  <user-outlined />
                  subnav 1
                </span>
                </template>
                <a-menu-item key="1">option1</a-menu-item>
                <a-menu-item key="2">option2</a-menu-item>
                <a-menu-item key="3">option3</a-menu-item>
                <a-menu-item key="4">option4</a-menu-item>
              </a-sub-menu>
              <a-sub-menu key="sub2">
                <template #title>
                <span>
                  <laptop-outlined />
                  subnav 2
                </span>
                </template>
                <a-menu-item key="5">option5</a-menu-item>
                <a-menu-item key="6">option6</a-menu-item>
                <a-menu-item key="7">option7</a-menu-item>
                <a-menu-item key="8">option8</a-menu-item>
              </a-sub-menu>
              <a-sub-menu key="sub3">
                <template #title>
                <span>
                  <notification-outlined />
                  subnav 3
                </span>
                </template>
                <a-menu-item key="9">option9</a-menu-item>
                <a-menu-item key="10">option10</a-menu-item>
                <a-menu-item key="11">option11</a-menu-item>
                <a-menu-item key="12">option12</a-menu-item>
              </a-sub-menu>
            </a-menu>
          </a-layout-sider>
          <a-layout-content :style="{ padding: '0 24px', minHeight: '280px' }">
            Content
            <a-button type="primary" @click="toggleTheme()">Primary Button</a-button>
            <a-switch v-model:checked="checked" @click="toggleTheme()"/>
          </a-layout-content>
        </a-layout>
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        Ant Design ©2018 Created by Ant UED
      </a-layout-footer>
    </a-layout>
  </a-config-provider>
</template>

<style scoped>
#components-layout-demo-top-side .logo {
  float: left;
  width: 120px;
  height: 31px;
  margin: 16px 24px 16px 0;
  background: rgba(255, 255, 255, 0.3);
}

.ant-row-rtl #components-layout-demo-top-side .logo {
  float: right;
  margin: 16px 0 16px 24px;
}

.site-layout-background {
  background: #fff;
}
</style>
