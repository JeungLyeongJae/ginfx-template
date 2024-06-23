<script lang="ts">
import {defineComponent, onMounted, reactive, ref, UnwrapRef} from "vue";
import {PlusOutlined} from '@ant-design/icons-vue';
import UserAddComponent from './user-add.vue';
import {getUsers, updateUser} from "./user.service.ts";
import {User} from "./user.ts";
import {cloneDeep} from "lodash-es";

const columns = [
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
    title: '用户名',
    dataIndex: 'username',
    key: 'username',
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
    title: '上传登录日期',
    dataIndex: 'last_login',
    key: 'last_login',
  },
  {
    title: '操作',
    dataIndex: 'operation',
  }
]

export default defineComponent({
  name: 'UserComponent',
  components: {
    PlusOutlined,
    UserAddComponent
  },

  setup() {
    const search = ref<string>('');
    const isOpen = ref<Boolean>(false);
    const isLoading = ref<Boolean>(false);
    const isSaving = ref<Boolean>(false);
    const alert = reactive({
      visible: false,
      message: '',
      type: 'info' as 'success' | 'info' | 'warning' | 'error'
    });

    // user list
    const users = ref<User[]>([]);
    const page = ref(1);
    const limit = ref(10);
    const total = ref<number>();

    const fetchUsers = async () => {
      isLoading.value = true;
      setTimeout(async () => {
        try {
          const response = await getUsers({page_number: page.value, page_size: limit.value, condition: search.value});
          users.value = [];
          users.value = response.users!;
          total.value = response.total_count!;
          isLoading.value = false;
        } catch (err) {
          isLoading.value = false;
          console.error('Failed to fetch users:', err);
        }
      }, 1000)
    };

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
        const u = users.value.filter(item => user.id === item.id)[0]
        Object.assign(u, editableData[user.id!]);
        delete editableData[user.id!];
        alert.visible = true;
        await fetchUsers()
      }
    }

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
      isOpen.value = true
    };

    const pagination = ref({
      current: page.value,
      pageSize: limit.value,
      total: total.value,
      showSizeChanger: true,
      showQuickJumper: true,
    });

    const handleTableChange = () => {
      fetchUsers();
    };

    const editableData: UnwrapRef<Record<number, User>> = reactive({});

    const edit = (id: number) => {
      editableData[id] = cloneDeep(users.value.filter(item => id === item.id)[0]);
    };
    const save = (id: number) => {
      return new Promise((resolve) => {
        setTimeout(() => resolve(saveUser(editableData[id])), 1000);
      });

    };
    const cancel = (id: number) => {
      delete editableData[id];
    };

    function handleResizeColumn(w, col) {
      col.width = w;
    }

    return {
      search,
      isOpen,
      isLoading,
      isSaving,
      alert,
      onSearch,
      onReload,
      showModal,

      pagination,
      columns,
      handleTableChange,
      users,
      editableData,
      edit,
      save,
      cancel,
      handleResizeColumn,
    }
  },
});
</script>

<template>
  <a-card style="width: 100%; height: 12%; margin: 10px; min-height: 90px; min-width: 500px" :hoverable="true"
          :bordered="false">
    <div style="align-items: center; display: flex; width: 100%;">
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
  <!--  <UserAddComponent :isOpened="isOpen.valueOf()"></UserAddComponent>-->

  <a-card style="width: 100%; height: 88%; margin: 10px; min-width: 500px" :hoverable="true" :bordered="false">
    <a-alert
        v-if="alert.visible"
        :message="alert.message"
        :type="alert.type"
        show-icon
        closable
        :banner="true"
        @close="alert.visible = false"
    />
    <a-table
        :columns="columns"
        :dataSource="users"
        :pagination="pagination"
        @change="handleTableChange"
        rowKey="id"
        bordered @resizeColumn="handleResizeColumn"
        :loading="isLoading.valueOf()"
        :scroll="{ x: 800 }">
      <template #bodyCell="{ column, text, record }">
        <template v-if="['phone','username', 'age', 'enable'].includes(column.dataIndex)">
          <div>
            <a-input
                v-if="editableData[record.id]"
                v-model:value="editableData[record.id][column.dataIndex]"
                style="margin: -5px 0"
            />
            <template v-else>
              {{ text }}
            </template>
          </div>
        </template>
        <!--        <template v-else-if="['phone','username', 'age', 'enable'].includes(column.dataIndex)">-->
        <!--          <div class="editable-row-operations">-->
        <!--          <span v-if="editableData[record.id]">-->
        <!--            <a-popconfirm title="确定要修改?" @confirm="save(record.id)" ok-text="确定" cancel-text="取消">-->
        <!--              <a>保存</a>-->
        <!--            </a-popconfirm>-->
        <!--            <a-typography-link @click="cancel(record.id)">取消</a-typography-link>-->
        <!--          </span>-->
        <!--            <span v-else>-->
        <!--            <a @click="edit(record.id)">修改</a>-->
        <!--          </span>-->
        <!--          </div>-->
        <!--        </template>-->
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
  </a-card>
</template>

<style scoped>
.editable-row-operations a {
  margin-right: 8px;
}
</style>