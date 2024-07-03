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
import {authLoginApi, login} from "./auth.service.ts";

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
        await authLoginApi({username: form.username, password: form.password});
        message.success('Login successful！');
      } catch (err) {
        message.error('添加失败！请重试');
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
