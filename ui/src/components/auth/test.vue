<template>
  <div class="container">
    <div class="loginUser-container">
      <div class="loginUser-title">管理平台</div>
      <a-form
          :model="state.formState"
          :label-col="state.layoutConfig.labelCol"
          :wrapper-col="state.layoutConfig.wrapperCol"
          :rules="state.formRule"
          ref="formRef"
          layout="vertical"
          autocomplete="off"
      >
        <a-form-item label="账号" name="username">
          <a-input
              v-model:value="state.formState.username"
              allowClear
              placeholder="请输入账号"
              :disabled="state.spinning"
          />
        </a-form-item>

        <a-form-item label="密码" name="password">
          <a-input-password
              v-model:value="state.formState.password"
              :disabled="state.spinning"
              allowClear
              placeholder="请输入密码"
          />
        </a-form-item>

        <a-form-item name="remember" :wrapper-col="state.wrapperCol">
          <a-checkbox
              v-model:checked="state.formState.remember"
              :disabled="state.spinning"
          >记住密码</a-checkbox
          >
        </a-form-item>

        <a-form-item :wrapper-col="state.submitWrapperCol" class="submit-box">
          <a-button
              type="primary"
              html-type="submit"
              @click="loginAction"
              :loading="state.spinning"
              style="width: 100%; font-size: 16px; font-weight: bolder"
          >登录</a-button
          >
        </a-form-item>
      </a-form>
      <div class="description">
        <span class="description-prefix">没账号？</span>
        <span
            @click="jumpRegister"
            class="description-after"
            :disabled="state.spinning"
        >去注册</span
        >
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { reactive, ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import { message } from "ant-design-vue";
import { loginUser } from "../../service/user/userApi";

import type { FormInstance } from "ant-design-vue";

interface FormStateType {
  username: string;
  password: string;
  remember: boolean;
}
interface FormRuleType {
  username: Object;
  password: Object;
}
interface stateType {
  formState: FormStateType;
  formRule: FormRuleType;
  layoutConfig: any;
  wrapperCol: any;
  submitWrapperCol: any;
  spinning: boolean;
  backgroundImgUrl: string;
}
// 路由
const router = useRouter();
//store
const store = useStore();
const formRef = ref<FormInstance>();
const state: stateType = reactive({
  formState: {
    username: "",
    password: "",
    remember: false,
  },
  spinning: false,
  formRule: {
    username: [{ required: true, message: "请输入账号！" }],
    password: [{ required: true, message: "请输入密码！" }],
  },
  layoutConfig: {
    labelCol: {
      span: 8,
    },
    wrapperCol: {
      span: 24,
    },
  },
  wrapperCol: { offset: 0, span: 24 },
  submitWrapperCol: { offset: 0, span: 24 },
  backgroundImgUrl:
      "http://www.yongma16.xyz/staticFile/common/img/background.png",
});
/**
 * 初始化表单内容
 */
const initForm = () => {
  const userInfoItem: any = window.localStorage.getItem("userInfo");
  interface userInfoType {
    username: string;
    password: string;
    remember: boolean;
  }
  const userInfo: userInfoType = userInfoItem
      ? JSON.parse(userInfoItem)
      : {
        username: "",
        password: "",
        remember: false,
      };
  if (userInfo.username && userInfo.password) {
    state.formState.username = userInfo.username;
    state.formState.password = userInfo.password;
    state.formState.remember = userInfo.remember;
  }
};
/**
 * 前往注册！
 */
const jumpRegister = () => {
  // 带 hash，结果是 /about#team
  router.push({ path: "/register" });
};

/**
 * 前往home！
 */
const jumpHome = () => {
  // 带 hash，结果是 /about#team
  router.push({ path: "/" });
};
/**
 * 记住密码
 * @param params
 */
const rememberAction = (params: Object) => {
  window.localStorage.setItem("userInfo", JSON.stringify(params));
};
/**
 * 登录
 */
const loginAction = () => {
  if (formRef.value) {
    formRef.value.validate().then(async (res: any) => {
      state.spinning = true;
      const params = {
        username: state.formState.username,
        password: state.formState.password,
      };
      if (state.formState.remember) {
        rememberAction({ ...params, remember: state.formState.remember });
      }
      try {
        console.log('登录',params)
        // @ts-ignore
        await store.dispatch(
            "user/loginUser",
            params
        );
        // 跳转
        setTimeout(() => {
          jumpHome();
        }, 500);
        state.spinning = false;
      } catch (r: any) {
        message.error(JSON.stringify(r));
        state.spinning = false;
        throw Error(r);
      }
    });
  }
};

onMounted(() => {
  //初始化
  initForm();
});
</script>

<style lang="less">
.background {
  background-image: url("http://www.yongma16.xyz/staticFile/common/img/background.png");
  background-repeat: no-repeat;
  background-size: 100%;
}
.container {
  /*background: #262626;*/
  background-clip: border-box;
  position: absolute;
  width: 100vw;
  height: 100vh;
  .background();
}
.loginUser-container {
  position: absolute;
  width: 400px;
  height: 350px;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
.submit-box {
  text-align: center;
  width: 100%;
  margin: 0 auto;
}
.loginUser-container {
  background-color: rgba(255, 255, 255, 0.8);
  border-radius: 10px;
  padding: 0 20px;
}
.loginUser-title {
  margin-top: 20px;
  width: 100%;
  text-align: center;
  font-weight: bolder;
  font-size: 16px;
}
.description {
  margin-top: 20px;
  width: 100%;
  text-align: center;
  .description-after {
    color: #1890ff;
    cursor: pointer;
  }
}
</style>
