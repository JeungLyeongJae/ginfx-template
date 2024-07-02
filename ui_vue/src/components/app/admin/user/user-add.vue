<script lang="ts" setup>
import {defineEmits, defineProps, reactive, ref, UnwrapRef, watch} from 'vue';
import {LockOutlined, MailOutlined, PhoneOutlined, TeamOutlined, UserOutlined} from '@ant-design/icons-vue';
import {
  validateEmail,
  validateName,
  validatePassword,
  validatePhone,
  validateUserName
} from "../../../../../public/model/validator.ts";
import {User} from "./user.ts";
import {Rule} from "ant-design-vue/es/form";
import {addUser} from "./user.service.ts";
import {message} from "ant-design-vue";
import {keys} from "lodash-es";

// model 传值 func
const isOpened = defineProps({
  isOpened: {
    type: Boolean,
    required: true
  }
})
const emit = defineEmits(['handleIsOpenedChange'])

// model 事件 func
const confirmLoading = ref<boolean>(false);
const okButtonProps = ref<boolean>(true);
const handleOk = () => {
  formRef.value.validateFields()
      .then(() => {
        confirmLoading.value = true;
        okButtonProps.value = true;
        setTimeout(async () => {
          try {
            await addUser(formState);
            confirmLoading.value = false;
            formRef.value.resetFields();
            formRef.value.clearValidate();
            message.success('添加成功！');
            emit('handleIsOpenedChange', {isOpened: false, submitStatus: true});
          } catch (err) {
            confirmLoading.value = false;
            message.error('添加失败！请重试');
          }
        }, 1000)
      })
      .catch(() => {
        okButtonProps.value = true;
      });
};
const handleCancel = () => {
  formRef.value.clearValidate();
  formRef.value.resetFields();
  okButtonProps.value = true;
  emit('handleIsOpenedChange', {isOpened: false, submitStatus: false});
};

// form表单绑定数据
const formRef = ref();
const formState: UnwrapRef<User> = reactive({
  name: '',
  username: '',
  phone: '',
  password: '',
  email: '',
  enable: true,
});
const formRules: Record<string, Rule[]> = {
  username: [{required: true, validator: validateUserName}],
  password: [{required: true, validator: validatePassword}],
  name: [{required: true, validator: validateName}],
  phone: [{required: true, validator: validatePhone}],
  email: [{validator: validateEmail}]
}
const layout = {
  labelCol: {span: 8},
  wrapperCol: {span: 16},
};

const okButtonStatus = () => {
  formRef.value.validateFields()
      .then(() => {
        okButtonProps.value = false;
      })
      .catch(() => {
        okButtonProps.value = true;
      });
}
watch((formState), () => {
  if (Object.values(formState).slice(0, 4).every(value => value !== '')) {
    okButtonStatus();
  }
}, { deep: true });
</script>

<template>
  <div>
    <a-modal :open="isOpened.isOpened.valueOf()"
             title="新增用户"
             :confirm-loading="confirmLoading"
             :ok-button-props="{ disabled: okButtonProps.valueOf() }"
             @ok="handleOk"
             @cancel="handleCancel">
      <a-form
          ref="formRef"
          :model="formState"
          :rules="formRules"
          v-bind="layout"
          style="width: 80%; margin: 10px">
        <a-form-item :name="['username']" label="用户名" has-feedback>
          <a-input v-model:value="formState.username">
            <template #prefix>
              <TeamOutlined class="site-form-item-icon"/>
            </template>
          </a-input>
        </a-form-item>
        <a-form-item :name="['password']" label="密码" has-feedback>
          <a-input-password v-model:value="formState.password">
            <template #prefix>
              <LockOutlined class="site-form-item-icon"/>
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item :name="['name']" label="姓名" has-feedback>
          <a-input v-model:value="formState.name">
            <template #prefix>
              <UserOutlined class="site-form-item-icon"/>
            </template>
          </a-input>
        </a-form-item>
        <a-form-item :name="['phone']" label="手机号" has-feedback>
          <a-input v-model:value="formState.phone">
            <template #prefix>
              <PhoneOutlined class="site-form-item-icon"/>
            </template>
          </a-input>
        </a-form-item>
        <a-form-item :name="['email']" label="邮箱">
          <a-input v-model:value="formState.email">
            <template #prefix>
              <MailOutlined class="site-form-item-icon"/>
            </template>
          </a-input>
        </a-form-item>
        <a-form-item :name="['enable']" label="是否启用">
          <a-switch v-model:checked="formState.enable"/>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>