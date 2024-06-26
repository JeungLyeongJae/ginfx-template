<script lang="ts" setup>
import {onMounted, reactive, ref} from "vue";
import {PlusOutlined} from '@ant-design/icons-vue';
import {deleteUser, getUsers, updateUser} from "./user.service.ts";
import {User} from "./user.ts";
import moment from 'moment';
import {PaginationType} from "../../../../../public/model/pagination.ts";
import UserAddComponent from "./user-add.vue";
import {message} from "ant-design-vue";
import UserEdit from "./user-edit.vue";

// 搜索框
const search = ref<string>('');

// 新增用户model
const showAddModal = () => {
  isAddUserOpened.value = true
};
const isAddUserOpened = ref<Boolean>(false);
const handleIsAddUserOpenedChange = (data: { isOpened: boolean, submitStatus: boolean }) => {
  isAddUserOpened.value = data.isOpened
  if (data.submitStatus) {
    fetchUsers()
  }
};

// 编辑用户model
const showUpdateModal = (id: number) => {
  isUpdateUser.value = userTable.data.filter(item => id === item.id)[0]
  isUpdateUserOpened.value = true
};
const isUpdateUserOpened = ref<Boolean>(false);
const isUpdateUser = ref<User>(new User());
const handleIsUpdateUserOpenedChange = (data: { isOpened: boolean, submitStatus: boolean }) => {
  isUpdateUserOpened.value = data.isOpened
  if (data.submitStatus) {
    fetchUsers()
  }
};

// 编辑用户是否启用
const isEnabled = (id: number) => {
  const user = userTable.data.filter(item => id === item.id)[0]
  user.enable = !user.enable
  isLoading.value = true;
  setTimeout(async () => {
    try {
      await updateUser(user);
      isLoading.value = false;
      user.enable ? message.success('启用成功！') : message.success('禁用成功！');
    } catch (err) {
      isLoading.value = false;
      user.enable = !user.enable
      user.enable ? message.error('启用失败！') : message.error('禁用失败！');
    }
  }, 200)
}

// 删除用户
const deletedUser = (id: number) => {
  return new Promise<void>((resolve) => {
    setTimeout(async () => {
      try {
        await deleteUser({id: id})
        message.success('删除成功！');
      } catch (err) {
        message.error('删除失败！');
      }
    }, 1000);
    fetchUsers();
    resolve();
  });
}

// user table
const userTable = reactive({
  columns: [
    {
      title: '用户名',
      dataIndex: 'username',
      key: 'username'
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
      title: '邮箱',
      dataIndex: 'email',
      key: 'email',
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

// 分页操作
const handleTableChange = (p: PaginationType) => {
  userTable.pagination = p
  fetchUsers();
};
</script>

<template>
  <a-card style="width: 100%; height: 10%; margin: 10px; min-height: 90px; min-width: 500px" :hoverable="true"
          :bordered="false">
    <div style="align-items: center; display: flex; width: 100%; margin-top: 1%">
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
        <a-button type="primary" style="float: right" @click="showAddModal">
          <a-space style="margin-right: 5px">
            <PlusOutlined/>
          </a-space>
          <span style="font-family: 'Microsoft YaHei', sans-serif;">添加</span>
        </a-button>
      </div>
    </div>
  </a-card>
  <UserAddComponent :isOpened="isAddUserOpened.valueOf()"
                    @handleIsOpenedChange="handleIsAddUserOpenedChange">
  </UserAddComponent>

  <a-card style="width: 100%; height: auto; margin: 10px; min-width: 500px" :hoverable="true" :bordered="false">
    <a-table
        :columns="userTable.columns"
        :dataSource="userTable.data"
        :pagination="userTable.pagination"
        @change="handleTableChange"
        :loading="isLoading.valueOf()"
        :scroll="{ x: 1000 }">
      <template #bodyCell="{ column, text, record }">
        <template v-if="['enable'].includes(column.dataIndex)">
          <a-switch :checked="text" @change="isEnabled(record.id)"/>
        </template>
        <template v-else-if="['created_at','updated_at', 'last_login'].includes(column.dataIndex)">
          {{ text == '' || text == null ? '-' : moment(text).format('YYYY-MM-DD HH:mm:ss') }}
        </template>
        <template v-else-if="column.dataIndex === 'operation'">
          <span>
            <a @click="showUpdateModal(record.id)">修改</a>
          </span>
          <span style="margin-left: 3px; color: #666666">
            |
          </span>
          <span style="margin-left: 3px">
            <a-popconfirm title="确定要删除?" @confirm="deletedUser(record.id)" ok-text="确定" cancel-text="取消">
              <a>删除</a>
            </a-popconfirm>
          </span>
        </template>
      </template>
    </a-table>
  </a-card>
  <UserEdit :isOpened="isUpdateUserOpened.valueOf()" :user="isUpdateUser.valueOf()"
            @handleIsOpenedChange="handleIsUpdateUserOpenedChange"></UserEdit>
</template>

<style scoped>
</style>