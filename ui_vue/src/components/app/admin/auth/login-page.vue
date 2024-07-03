<template>
  <div class="login-page">
    <login-form/>
  </div>
</template>

<script lang="ts" setup>
import LoginForm from "./login-form.vue";
</script>

<style scoped>
.login-page {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f0f2f5;
}
</style>
