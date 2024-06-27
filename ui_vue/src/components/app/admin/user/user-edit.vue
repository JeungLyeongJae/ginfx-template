<script lang="ts" setup>
import {defineEmits, defineProps, PropType, reactive, ref, UnwrapRef, watch} from 'vue';
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
import {updateUser} from "./user.service.ts";
import {message} from "ant-design-vue";

// model 传值 func
const editModel = defineProps({
  isOpened: {
    type: Boolean,
    required: true
  },
  user: {
    type: Object as PropType<User>,
    required: true
  },
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
            await updateUser(formState);
            confirmLoading.value = false;
            formRef.value.resetFields();
            formRef.value.clearValidate();
            message.info('修改成功！');
            emit('handleIsOpenedChange', {isOpened: false, submitStatus: true});
          } catch (err) {
            confirmLoading.value = false;
            message.error('修改失败！请重试');
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
  id: -1,
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

watch(
    () => editModel.user,
    (newUser) => {
      if (newUser) {
        formState.id = newUser.id;
        formState.name = newUser.name;
        formState.username = newUser.username;
        formState.phone = newUser.phone;
        formState.password = newUser.password;
        formState.email = newUser.email;
        formState.enable = newUser.enable;
      }
    },
    {immediate: true}
);

const okButtonStatus = () => {
  if (formRef.value) {
    formRef.value.validateFields()
        .then(() => {
          okButtonProps.value = false;
        })
        .catch(() => {
          okButtonProps.value = true;
        });
  }
}
watch((formState), () => {
  if (Object.values(formState).every(value => value !== '')) {
    okButtonStatus();
  }
}, {deep: true});
</script>

<template>
  <div>
    <a-modal :open="editModel.isOpened.valueOf()"
             title="编辑用户"
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