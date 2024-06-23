<script lang="ts">
import {defineComponent, onMounted, PropType, toRefs} from 'vue';


export default defineComponent({
  name: 'UserAddComponent',
  props: {
    isOpened: {
      type: Boolean as PropType<boolean>,
      required: true
    }
  },

  setup(props) {

    let isOpen  = toRefs(props);

    onMounted(() => {
      console.log(isOpen)
    })

    const handleOk = (e: MouseEvent) => {
      console.log(e);
    };

    const handleCancel = (e: MouseEvent) => {
      console.log(e);
    };

    return {
      isOpen,
      handleOk,
      handleCancel
    }
  }
})
</script>

<template>
  <div>
    <!--    <a-button type="primary" @click="showModal">Open Modal</a-button>-->
    <a-modal v-model:open="isOpen" title="新增用户" @ok="handleOk" @cancel="handleCancel">
      <p>Some contents...</p>
      <p>Some contents...</p>
      <p>Some contents...</p>
    </a-modal>
  </div>
</template>