<script lang="ts">
import {defineComponent, onMounted, reactive, ref, UnwrapRef} from "vue";
import {PlusOutlined} from '@ant-design/icons-vue';
import UserAddComponent from './user-add.vue';
import {getUsers} from "./user.service.ts";
import {User} from "./user.ts";

const UserMapping: Map<string, string> = new Map();

UserMapping.set('name', '姓名');
UserMapping.set('phone', '手机号');
UserMapping.set('username', '用户名');
UserMapping.set('enable', '是否启用');
UserMapping.set('created_at', '创建日期');
UserMapping.set('updated_at', '更新日期');
UserMapping.set('last_login', '上传登录日期');

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
]

export default defineComponent({
  name: 'UserComponent',
  computed: {
    UserMapping() {
      return UserMapping
    }
  },
  components: {
    PlusOutlined,
    UserAddComponent
  },

  setup() {
    const search = ref<string>('');
    const isOpen = ref<Boolean>(false);
    const isLoading = ref<Boolean>(false);

    // user list
    const users = ref<User[]>([]);
    const page = ref(1);
    const limit = ref(10);
    const total = ref<number>();

    const fetchUsers = async () => {
      isLoading.value = true;
      try {
        const response = await getUsers({page_number: page.value, page_size: limit.value, condition: search.value});
        users.value = response.users!;
        total.value = response.total_count!;
        isLoading.value = false;
      } catch (err) {
        console.error('Failed to fetch users:', err);
      }
    };

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
      users.value = [];
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

    const editableData: UnwrapRef<Record<string, User>> = reactive({});

    const edit = (key: string) => {
      // editableData[key] = cloneDeep(users.value.filter(item => key === item.key)[0]);
    };
    const save = (key: string) => {
      // Object.assign(dataSource.value.filter(item => key === item.key)[0], editableData[key]);
      // delete editableData[key];
    };
    const cancel = (key: string) => {
      delete editableData[key];
    };

    function handleResizeColumn(w, col) {
      col.width = w;
    }

    return {
      search,
      isOpen,
      isLoading,
      onSearch,
      onReload,
      showModal,

      pagination,
      columns,
      handleTableChange,
      users,
      editingKey: '',
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
  <a-card style="width: 100%; height: 12%; margin: 10px" :hoverable="true" :bordered="false">
    <div style="align-items: center; display: flex; width: 100%;">
      <div style="width: 80%;">
        <span style="font-family: 'Microsoft YaHei', sans-serif; margin-right: 10px; color: #666666;">用户名：</span>
        <a-input-search
            v-model:value="search"
            placeholder="请输入用户名"
            enter-button
            @search="onSearch"
            style="width: 50%; margin-right: 10px"
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
  <UserAddComponent :open="isOpen.valueOf()"></UserAddComponent>


  <a-card style="width: 100%; height: 88%; margin: 10px" :hoverable="true" :bordered="false">
    <a-table
        :columns="columns.values()"
        :dataSource="users"
        :pagination="pagination"
        @change="handleTableChange"
        rowKey="id"
        bordered @resizeColumn="handleResizeColumn" :loading="isLoading.valueOf()">
<!--      <template #headerCell="{ column }">-->
<!--      </template>-->

      <!--      <template #bodyCell="{ column, text, record }">-->
      <!--        <template v-if="['name', 'age', 'address'].includes(column.dataIndex)">-->
      <!--          <div>-->
      <!--            <a-input-->
      <!--                v-if="editableData[record.key]"-->
      <!--                v-model:value="editableData[record.key][column.dataIndex]"-->
      <!--                style="margin: -5px 0"-->
      <!--            />-->
      <!--            <template v-else>-->
      <!--              {{ text }}-->
      <!--            </template>-->
      <!--          </div>-->
      <!--        </template>-->
      <!--        <template v-else-if="column.dataIndex === 'operation'">-->
      <!--          <div class="editable-row-operations">-->
      <!--              <span v-if="editableData[record.key]">-->
      <!--                <a-typography-link @click="save(record.key)">Save</a-typography-link>-->
      <!--                <a-popconfirm title="Sure to cancel?" @confirm="cancel(record.key)">-->
      <!--                  <a>Cancel</a>-->
      <!--                </a-popconfirm>-->
      <!--              </span>-->
      <!--            <span v-else>-->
      <!--                <a @click="edit(record.key)">Edit</a>-->
      <!--              </span>-->
      <!--          </div>-->
      <!--        </template>-->
      <!--      </template>-->
    </a-table>
  </a-card>
</template>

<style scoped>
.editable-row-operations a {
  margin-right: 8px;
}
</style>