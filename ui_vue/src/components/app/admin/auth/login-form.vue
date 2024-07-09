<template>
  <a-form
      :form="form"
      @submit="handleSubmit"
      class="login-form"
  >
    <a-form-item>
      <a-input
          v-model:value="form.username"
          placeholder="Username"
          prefix="user"
      />
    </a-form-item>
    <a-form-item>
      <a-input-password
          v-model:value="form.password"
          placeholder="Password"
      />
    </a-form-item>
    <a-form-item>
      <a-button type="primary" html-type="submit" class="login-form-button">
        Log in
      </a-button>
    </a-form-item>
  </a-form>
</template>

<script lang="ts" setup>
import {reactive} from 'vue';
import {message} from 'ant-design-vue';
import {login} from "./auth.service.ts";
import {useRouter} from "vue-router";

const router = useRouter();

const form = reactive({
  username: '',
  password: '',
})

const handleSubmit = (e: any) => {
  e.preventDefault();
  if (!form.username || !form.password) {
    message.error('Please fill in both fields');
    return;
  } else {
    setTimeout(async () => {
      try {
        const data = await login({username: form.username, password: form.password});
        message.success('Login successful！');

        // 将 token 存储在 localStorage 中，用于后续的身份验证
        sessionStorage.setItem('token', data.token);

        // 使用 Vue Router 的 router.push() 方法进行页面跳转
        await router.push('home'); // 跳转到名为 home 的路由
      } catch (err) {
        message.error('Login defeated');
      }
    }, 1000)
  }
};

</script>

<style scoped>
.login-form {
  max-width: 300px;
  margin: auto;
  padding: 20px;
  border: 1px solid #d9d9d9;
  border-radius: 4px;
  background-color: #fff;
}

.login-form-button {
  width: 100%;
}
</style>
