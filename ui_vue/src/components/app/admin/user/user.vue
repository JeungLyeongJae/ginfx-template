<script lang="ts" setup>
import {onMounted, reactive, ref, UnwrapRef} from "vue";
import {PlusOutlined} from '@ant-design/icons-vue';
import {getUsers, updateUser} from "./user.service.ts";
import {User} from "./user.ts";
import {cloneDeep} from "lodash-es";
import moment from 'moment';
import {PaginationType} from "../../../../../public/model/pagination.ts";
import {Rule} from "ant-design-vue/es/form";
import UserAddComponent from "./user-add.vue";

// 搜索框
const search = ref<string>('');

// 新增用户model
const isOpened = ref<Boolean>(false);
const handleIsOpenedChange = (value: boolean) => {
  isOpened.value = value
};

const validateUserName = async (_rule: Rule, value: string) => {
  console.log(value)
  if (!value) {
    return Promise.reject('请输入用户名!');
  }
  if (value.length < 3) {
    return Promise.reject('用户名不能小于3位英文');
  }
  return Promise.resolve();
};

// user table
const userTable = reactive({
  columns: [
    {
      title: '用户名',
      dataIndex: 'username',
      key: 'username',
      rules: [{required: true, validator: validateUserName}]
    },
    {
      title: '姓名',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: '手机号',
      dataIndex: 'phone',
      key: 'phone',
    },
    {
      title: '是否启用',
      dataIndex: 'enable',
      key: 'enable',
    },
    {
      title: '创建日期',
      dataIndex: 'created_at',
      key: 'created_at',
    },
    {
      title: '更新日期',
      dataIndex: 'updated_at',
      key: 'updated_at',
    },
    {
      title: '上次登录',
      dataIndex: 'last_login',
      key: 'last_login',
    },
    {
      title: '操作',
      dataIndex: 'operation',
    }
  ],
  data: [] as User[],
  pagination: {
    current: 1,
    pageSize: 10,
    total: 0,
    showSizeChanger: true,
    showQuickJumper: true,
  } as PaginationType,
  isLoading: false,
})
const isLoading = ref<Boolean>(false);
const alert = reactive({
  visible: false,
  message: '',
  type: 'info' as 'success' | 'info' | 'warning' | 'error'
});

// 初始化执行
onMounted(() => {
  fetchUsers();
});

// 用户名搜索
const onSearch = () => {
  fetchUsers();
};

// 刷新
const onReload = () => {
  search.value = '';
  fetchUsers();
};

// 新增用户
const showModal = () => {
  isOpened.value = true
};

// getUserList
const fetchUsers = async () => {
  isLoading.value = true;
  setTimeout(async () => {
    try {
      const response = await getUsers({
        page_number: userTable.pagination.current,
        page_size: userTable.pagination.pageSize,
        condition: search.value
      });
      userTable.data = [];
      userTable.data = response.users!;
      userTable.pagination.total = response.total_count!;

      isLoading.value = false;
    } catch (err) {
      isLoading.value = false;
      console.error('Failed to fetch users:', err);
    }
  }, 1000)
};

// updateUser
const saveUser = async (user: User) => {
  try {
    await updateUser(user);
    alert.message = 'Data saved successfully!';
    alert.type = 'success';
  } catch (err) {
    console.error('Failed to update user:', err);
    alert.message = 'Failed to save data.';
    alert.type = 'error';
  } finally {
    const u = userTable.data.filter(item => user.id === item.id)[0]
    Object.assign(u, editableData[user.id!]);
    delete editableData[user.id!];
    alert.visible = true;
    await fetchUsers()
  }
}

// 分页操作
const handleTableChange = (p: PaginationType) => {
  userTable.pagination = p
  fetchUsers();
};

const editableData: UnwrapRef<Record<number, User>> = reactive({});

const edit = (id: number) => {
  editableData[id] = cloneDeep(userTable.data.filter(item => id === item.id)[0]);
};

const save = (id: number) => {
  return new Promise((resolve) => {
    setTimeout(() => resolve(saveUser(editableData[id])), 1000);
  });

};

const cancel = (id: number) => {
  delete editableData[id];
};

const form = ref<User>();
</script>

<template>
  <a-card style="width: 100%; height: 12%; margin: 10px; min-height: 90px; min-width: 500px" :hoverable="true"
          :bordered="false">
    <div style="align-items: center; display: flex; width: 100%; margin-top: 1.5%">
      <div style="width: 80%;">
        <span style="font-family: 'Microsoft YaHei', sans-serif; margin-right: 10px; color: #666666;">用户名：</span>
        <a-input-search
            v-model:value="search"
            placeholder="请输入用户名"
            enter-button
            @search="onSearch"
            style="width: 200px; margin-right: 10px"
        />
        <a-button type="primary" @click="onReload">
          <span style="font-family: 'Microsoft YaHei', sans-serif;">刷新</span>
        </a-button>
      </div>
      <div style="width: 20%; justify-content: flex-end;">
        <a-button type="primary" style="float: right" @click="showModal">
          <a-space style="margin-right: 5px">
            <PlusOutlined/>
          </a-space>
          <span style="font-family: 'Microsoft YaHei', sans-serif;">添加</span>
        </a-button>
      </div>
    </div>
  </a-card>
  <UserAddComponent :isOpened="isOpened.valueOf()" @handleIsOpenedChange="handleIsOpenedChange"></UserAddComponent>

  <a-card style="width: 100%; height: 88%; margin: 10px; min-width: 500px" :hoverable="true" :bordered="false">
    <a-form :form="form" @submit.prevent="save">
      <a-table
          :columns="userTable.columns"
          :dataSource="userTable.data"
          :pagination="userTable.pagination"
          @change="handleTableChange"
          rowKey="id"
          :loading="isLoading.valueOf()"
          :scroll="{ x: 1000 }">
        <template #bodyCell="{ column, text, record }">
          <template v-if="['phone','username', 'age'].includes(column.dataIndex)">
            <div>
              <a-form-item v-if="editableData[record.id]"
                           :name="column.dataIndex"
                           :rules="column.rules">
                <a-input v-model:value="record[column.dataIndex]"/>
                <!--                <a-input-->
                <!--                    v-model:value="editableData[record.id][column.dataIndex]"-->
                <!--                    style="margin: -5px 0"-->
                <!--                />-->
              </a-form-item>
              <template v-else>
                {{ text }}
              </template>
            </div>
          </template>
          <template v-else-if="['enable'].includes(column.dataIndex)">
            <a-switch :checked="text == 1"/>
          </template>
          <template v-else-if="['created_at','updated_at', 'last_login'].includes(column.dataIndex)">
            {{ text == '' || text == null ? '-' : moment(text).format('YYYY-MM-DD HH:mm:ss') }}
          </template>
          <template v-else-if="column.dataIndex === 'operation'">
            <div class="editable-row-operations">
          <span v-if="editableData[record.id]">
            <a-popconfirm title="确定要修改?" @confirm="save(record.id)" ok-text="确定" cancel-text="取消">
              <a>保存</a>
            </a-popconfirm>
            <a-typography-link @click="cancel(record.id)">取消</a-typography-link>
          </span>
              <span v-else>
            <a @click="edit(record.id)">修改</a>
          </span>
            </div>
          </template>
        </template>
      </a-table>
    </a-form>
  </a-card>
</template>

<style scoped>
.editable-row-operations a {
  margin-right: 8px;
}
</style>