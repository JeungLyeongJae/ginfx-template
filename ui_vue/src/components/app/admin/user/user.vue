<script lang="ts">
import {defineComponent, reactive, ref, UnwrapRef} from "vue";
import {cloneDeep} from 'lodash-es';
import {PlusOutlined} from '@ant-design/icons-vue';
import UserAddComponent from './user-add.vue';

const columns = [
  {
    title: 'name',
    dataIndex: 'name',
    width: '25%',
    slots: {customRender: 'name'},
  },
  {
    title: 'age',
    dataIndex: 'age',
    width: '15%',
    slots: {customRender: 'age'},
  },
  {
    title: 'address',
    dataIndex: 'address',
    width: '40%',
    slots: {customRender: 'address'},
  },
  {
    title: 'operation',
    dataIndex: 'operation',
    slots: {customRender: 'operation'},
  },
];

interface DataItem {
  key: string;
  name: string;
  age: number;
  address: string;
}

const data: DataItem[] = [];
for (let i = 0; i < 100; i++) {
  data.push({
    key: i.toString(),
    name: `Edrward ${i}`,
    age: 32,
    address: `London Park no. ${i}`,
  });
}

export default defineComponent({
  name: 'UserComponent',
  components: {
    PlusOutlined,
    UserAddComponent
  },

  setup() {
    const value = ref<string>('');

    const open = ref<Boolean>(false);

    const onSearch = (searchValue: string) => {
      console.log('use value', searchValue);
      console.log('or use this.value', value.value);
    };

    const showModal = () => {
      open.value = true
    };

    const dataSource = ref(data);
    const editableData: UnwrapRef<Record<string, DataItem>> = reactive({});

    const edit = (key: string) => {
      editableData[key] = cloneDeep(dataSource.value.filter(item => key === item.key)[0]);
    };
    const save = (key: string) => {
      Object.assign(dataSource.value.filter(item => key === item.key)[0], editableData[key]);
      delete editableData[key];
    };
    const cancel = (key: string) => {
      delete editableData[key];
    };

    function handleResizeColumn(w, col) {
      col.width = w;
    }

    return {
      value,
      open,
      onSearch,
      showModal,

      dataSource,
      columns,
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
            v-model:value="value"
            placeholder="请输入用户名"
            enter-button
            @search="onSearch"
            style="width: 50%; margin-right: 10px"
        />
        <a-button type="primary">
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
  <UserAddComponent :open="open.valueOf()"></UserAddComponent>


  <a-card style="width: 100%; height: 88%; margin: 10px" :hoverable="true" :bordered="false">
    <a-table :columns="columns" :data-source="dataSource" bordered @resizeColumn="handleResizeColumn">
      <template #bodyCell="{ column, text, record }">
        <template v-if="['name', 'age', 'address'].includes(column.dataIndex)">
          <div>
            <a-input
                v-if="editableData[record.key]"
                v-model:value="editableData[record.key][column.dataIndex]"
                style="margin: -5px 0"
            />
            <template v-else>
              {{ text }}
            </template>
          </div>
        </template>
        <template v-else-if="column.dataIndex === 'operation'">
          <div class="editable-row-operations">
              <span v-if="editableData[record.key]">
                <a-typography-link @click="save(record.key)">Save</a-typography-link>
                <a-popconfirm title="Sure to cancel?" @confirm="cancel(record.key)">
                  <a>Cancel</a>
                </a-popconfirm>
              </span>
            <span v-else>
                <a @click="edit(record.key)">Edit</a>
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