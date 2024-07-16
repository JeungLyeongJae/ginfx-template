<template>
  <div class="body">
    <a-card class="login-card" :bordered="false">
      <a-form
          :form="form"
          @submit="handleSubmit"
          class="login-form"
      >
        <a-form-item>
          <a-input
              v-model:value="form.username"
              placeholder="用户名"
              style="border-radius: 16px;"
          />
        </a-form-item>
        <a-form-item>
          <a-input-password
              v-model:value="form.password"
              placeholder="密码"
              style="border-radius: 16px;"
          />
        </a-form-item>
        <a-form-item style="text-align: center">
          <a-button :loading="isLoading.valueOf()" type="primary" html-type="submit" class="login-form-button">
            登录
          </a-button>
        </a-form-item>
        <Puzzle @success="onSuccess"
               :show="isShow"
               :canvas-width="310"
               :canvas-height="200"
               :puzzle-scale="0.7"
               :slider-size="25"
               :range="5"
        />
      </a-form>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from 'vue';
import {message} from 'ant-design-vue';
import {login} from "./auth.service.ts";
import {useRouter} from "vue-router";
import Puzzle from 'vue3-puzzle-vcode';

const router = useRouter();

const form = reactive({
  username: '',
  password: '',
})

const isShow = ref(false);
const onShow = () => {
  isShow.value = true;
};
const onClose = () => {
  isShow.value = false;
};
const onSuccess = () => {
  onClose();
  if (isLoading.value) {
    loginSubmit();
  }
};

const isLoading = ref(false);
const handleSubmit = (e: any) => {
  e.preventDefault();
  if (!form.username || !form.password) {
    message.error('请输入用户名与密码！');
    return;
  } else {
    isLoading.value = true;
    setTimeout(() => {
      onShow();
    }, 1000)
  }
};

const loginSubmit = () => {
  setTimeout(async () => {
    try {
      const data = await login({username: form.username, password: form.password});
      message.success('登录成功！');

      // 将 token 存储在 localStorage 中，用于后续的身份验证
      sessionStorage.setItem('token', data.token);

      // 使用 Vue Router 的 router.push() 方法进行页面跳转
      await router.push('home'); // 跳转到名为 home 的路由
    } catch (err) {
      message.error('登录失败');
    } finally {
      isLoading.value = false;
    }
  }, 1000)
}

</script>

<style scoped>
.body {
  background-image: url('../../../../assets/img/background.png');
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  background-color: #d3e3fd;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100vw;
  height: 100vh;
  margin: 0;
  padding: 0;
}

.login-card {
  width: 500px;
  height: 300px;
  margin: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.5);
  border-radius: 16px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #d3e3fd;
  opacity: 0.75;
}

.login-form {
  width: 350px;
  margin: auto;
  padding: 20px 20px 10px 20px;
  border-radius: 16px;
}

.login-form-button {
  width: 80%;
  border-radius: 16px;
}
</style>
