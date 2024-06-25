<script lang="ts" setup>
import {defineEmits, defineProps, reactive, ref, UnwrapRef} from 'vue';
import {LockOutlined, MailOutlined, PhoneOutlined, TeamOutlined, UserOutlined} from '@ant-design/icons-vue';
import {
  isExist,
  validateName,
  validatePassword,
  validatePhone,
  validateUserName
} from "../../../../../public/model/validator.ts";
import {User} from "./user.ts";
import {Rule} from "ant-design-vue/es/form";

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
  confirmLoading.value = true;
  setTimeout(() => {
    confirmLoading.value = false;
    emit('handleIsOpenedChange', false);
  }, 1000);

};
const handleCancel = () => {
  formRef.value.resetFields();
  emit('handleIsOpenedChange', false);
};

// form表单绑定数据
const formRef = ref();
const formState: UnwrapRef<User> = reactive({
  name: '',
  username: '',
  phone: '',
  password: '',
});
const formRules: Record<string, Rule[]> = {
  username: [{required: true, validator: validateUserName, trigger: 'blur'}],
  password: [{required: true, validator: validatePassword}],
  name: [{required: true, validator: validateName}],
  phone: [{required: true, validator: validatePhone}]
}
const layout = {
  labelCol: {span: 8},
  wrapperCol: {span: 16},
};
// const validateStatus = reactive({});
//
// const getValidateStatus = (key, dataIndex) => {
//   return validateStatus[dataIndex] || '';
// };


</script>

<template>
  <div>
    <a-modal :open="isOpened.isOpened.valueOf()"
             title="新增用户"
             :confirm-loading="confirmLoading"
             @ok="handleOk"
             :ok-button-props="{ disabled: okButtonProps }"
             @cancel="handleCancel">
      <a-form
          ref="formRef"
          :model="formState"
          :rules="formRules"
          v-bind="layout"
          style="width: 80%; margin: 10px"
          :submit="okButtonProps = false">
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
        <a-form-item :name="['enable']" label="邮箱">
          <a-input v-model:value="formState.enable">
            <template #prefix>
              <MailOutlined class="site-form-item-icon"/>
            </template>
          </a-input>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>