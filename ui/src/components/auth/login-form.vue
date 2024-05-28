<script lang="ts">
import {defineComponent} from 'vue';
import {Button, Form, FormItem, Input} from 'ant-design-vue';

export default defineComponent({
  name: "login-form",

  constructor() {

  },

  data() {
    return {
      form: {},
      username: '',
      password: ''
    };
  },

  components: {
    'a-form': Form,
    'a-form-item': FormItem,
    'a-input': Input,
    'a-button': Button
  },

  setup() {
  },


  methods: {
    handleSubmit() {

      // this.form.validateFields().then(() => {
      //   axios.post('your_login_endpoint', {
      //     username: this.username,
      //     password: this.password
      //   })
      //       .then(response => {
      //         console.log(response.data);
      //         // 处理登录成功逻辑
      //       })
      //       .catch(error => {
      //         console.error(error);
      //         // 处理登录失败逻形
      //       });
      // }).catch(errorInfo => {
      //   console.log('Validation error:', errorInfo);
      // });
    }
  }
});
</script>

<template>
  <div class="login-form">
    <a-form :form="form" @submit="handleSubmit">
      <a-form-item label="用户名" :rules="[{ required: true, message: 'Please enter your username' }]">
        <a-input v-model="username"/>
      </a-form-item>
      <a-form-item label="密码" :rules="[{ required: true, message: 'Please enter your password' }]">
        <a-input type="password" v-model="password"/>
      </a-form-item>
      <a-form-item>
        <a-button type="primary" html-type="submit">登录</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<style scoped>
div.login-form {
  margin: 1%;
  padding: 1%;
  width: 90%;
  height: 90%;
}

:where(.css-dev-only-do-not-override-16pw25h) {
  .ant-input {
    width: 80%;
  }

  .ant-btn {
    width: 60%;
  }
}
</style>