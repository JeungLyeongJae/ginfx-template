<script lang="ts">
import {defineComponent, onMounted, ref} from 'vue';


export default defineComponent({
  name: 'UserAddComponent',
  props: {
    open: Boolean
  },

  setup() {
    const isOpened = ref<Boolean>(false);

    const handleOk = (e: MouseEvent) => {
      console.log(e);
    };

    return {
      isOpened,
      handleOk
    }
  }
})
</script>

<template>
  <div>
    <!--    <a-button type="primary" @click="showModal">Open Modal</a-button>-->
    <a-modal v-model:open="isOpened" title="新增用户" @ok="handleOk">
      <p>Some contents...</p>
      <p>Some contents...</p>
      <p>Some contents...</p>
    </a-modal>
  </div>
</template>