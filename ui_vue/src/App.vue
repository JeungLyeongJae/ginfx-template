<script setup lang="ts">


// const updateAvailable = ref<boolean>(false)

// onMounted(() => {
//   if ('serviceWorker' in navigator) {
//     navigator.serviceWorker.addEventListener('controllerchange', (e) => {
//       console.log(e)
//       updateAvailable.value = true;
//       message.info('New version available.');
//       refreshPage();
//     });
//   }
// })

// function refreshPage() {
//   window.location.reload();
// }
</script>

<template>
  <router-view></router-view>
</template>

<style scoped>
</style>
