<script setup lang="ts">
import {RouterLink, RouterView} from 'vue-router'
</script>

<template>
  <div class="app">
    <RouterLink to="/login"></RouterLink>
    <RouterView/>
  </div>
</template>

<style scoped>
  div.app {
    margin: 1%;
    padding: 1%;
    border: 1px solid #282828;
    border-radius: 3%;
    width: 26%;
    height: 42%;
    position: fixed;
    right: 15%;
    top: 45%; /* 可选，可以根据需要设置垂直位置 */
    transform: translateY(-50%); /* 可选，用于垂直居中 */
    display: grid;
    place-items: center; /* 水平和垂直居中 */
  }
</style>
