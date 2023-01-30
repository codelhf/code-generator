<template>
  <table-layout title="模板字段">
    <!-- 左侧树 -->
    <template #tree>
      <tree-list
        ref="treeDom"
        :title="'分组列表'"
        :data="typeList"
        :disable-right-click="true"
        :show-check-box="false"
        :enable-click-check="true"
        :enable-parent-check="false"
        @node-click="handleNodeClick"
      />
    </template>
    <!-- 搜索表单 -->
    <template #form>
      <el-form :model="listQuery" :inline="true" label-width="120px" label-suffix=":">
        <el-form-item :label="$t('templateField.listQuery.name')">
          <el-input v-model="listQuery.name" :placeholder="$t('templateField.listQuery.placeholderName')" />
        </el-form-item>
        <el-form-item :label="$t('templateField.listQuery.desc')">
          <el-input v-model="listQuery.desc" :placeholder="$t('templateField.listQuery.placeholderDesc')" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :icon="useRenderIcon('search')" @click="handleSearch">{{ $t('templateField.listButton.search') }}</el-button>
          <el-button type="primary" :icon="useRenderIcon('refresh-left')" @click="handleReset">{{ $t('templateField.listButton.reset') }}</el-button>
        </el-form-item>
      </el-form>
    </template>
    <!--按钮组-->
    <template #button>
      <el-button type="primary" :icon="useRenderIcon('plus')" @click="handleDetail">{{ $t('templateField.listButton.add') }}</el-button>
      <el-button type="danger" :icon="useRenderIcon('delete')" @click="handleDelete">{{ $t('templateField.listButton.delete') }}</el-button>
    </template>
    <!-- 列表 -->
    <template #table>
      <table-list
        ref="tableListDom"
        :data="list"
        :loading="listLoading"
        :show-selection="true"
        :show-page="total>0"
        :page-total="total"
        v-model:page-num="listQuery.pageNum"
        v-model:page-size="listQuery.pageSize"
        @pagination="getList"
      >
        <el-table-column :label="$t('templateField.table.name')" align="center">
          <template #default="scope">
            <span>{{ scope.row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('templateField.table.desc')" align="center">
          <template #default="scope">
            <span>{{ scope.row.desc }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('templateField.table.value')" align="center">
          <template #default="scope">
            <span>{{ scope.row.value }}</span>
          </template>
        </el-table-column>
        <el-table-column v-if="listQuery.type === 2" :label="$t('templateField.table.operation')" align="center">
          <template #default="scope">
            <el-button circle size="small" :icon="useRenderIcon('edit')" @click="handleDetail(scope.row)" />
            <el-button circle size="small" :icon="useRenderIcon('delete')" @click="handleDelete(scope.row)" />
          </template>
        </el-table-column>
      </table-list>
      <!--dialog-->
      <el-dialog
        :title="templateField.id ? $t('templateField.item.editTitle') : $t('templateField.item.addTitle')"
        v-model="dialogFormVisible"
        @close="handleFormClose"
      >
        <el-form ref="templateFieldForm" :model="templateField" :rules="templateFieldRules" label-width="120px" label-suffix=":">
          <el-form-item :label="$t('templateField.item.name')" prop="name">
            <el-input v-model="templateField.name" :placeholder="$t('templateField.item.placeholderName')" @blur="checkField" />
          </el-form-item>
          <el-form-item :label="$t('templateField.item.desc')">
            <el-input v-model="templateField.desc" :placeholder="$t('templateField.item.placeholderDesc')" />
          </el-form-item>
          <el-form-item :label="$t('templateField.item.value')" prop="value">
            <el-input v-model="templateField.value" :placeholder="$t('templateField.item.placeholderValue')" />
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="handleFormClose">{{ $t('common.form.cancel') }}</el-button>
            <el-button type="primary" @click="handleFormSubmit">{{ $t('common.form.confirm') }}</el-button>
          </span>
        </template>
      </el-dialog>
    </template>
  </table-layout>
</template>

<script>
import TreeList from '@/components/TreeList/index.vue'
import TableList from '@/components/TableList/index.vue'
import TableLayout from "@/components/TableList/layout.vue"
export default {
  name: 'TemplateField',
  components: {
    TreeList,
    TableList,
    TableLayout
  }
}
</script>

<script setup>
import { onMounted, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { $t } from '@/i18n'
import {
  templateFieldList, templateFieldSelect, templateFieldInsert,
  templateFieldUpdate, templateFieldDelete, templateFieldCheck
} from '@/pages/code-generator/api/template-field'

// 左侧树
function handleNodeClick(data) {
  listQuery.value.type = data.id
  getList()
}

// 生命周期
onMounted(() => {
  getTypeList()
})

const typeList = ref([])
function getTypeList() {
  setTimeout(() => {
    typeList.value = [
      { id: 1, name: $t('templateField.table.type1') },
      { id: 2, name: $t('templateField.table.type2') }
    ]
    getList()
  }, 100)
}

// 搜索部分
const listQuery = ref({
  pageNum: 1,
  pageSize: 10,
  sort: '',
  name: '',
  desc: '',
  type: null
})

function handleSearch() {
  listQuery.value.pageNum = 1
  getList()
}

function handleReset() {
  listQuery.value = {
    pageNum: 1,
    pageSize: 10,
    sort: '',
    name: '',
    desc: '',
    type: null
  }
}

// 列表部分
const tableListDom = ref(null)
const listLoading = ref(false)
const list = ref([])
const total = ref(0)

function getList() {
  listLoading.value = true
  templateFieldList(listQuery.value).then(res => {
    list.value = res.data.list
    total.value = res.data.total
    // Just to simulate the time of the request
    setTimeout(() => {
      listLoading.value = false
    }, 1.5 * 1000)
  })
}

function handleDelete(row) {
  ElMessageBox.confirm($t('common.confirm.deleteOne'), $t('common.confirm.title'), {
    cancelButtonText: $t('common.confirm.cancel'),
    confirmButtonText: $t('common.confirm.confirm'),
    type: 'warning'
  }).then(() => {
    if (row && row.id) {
      templateFieldDelete(row.id).then(res => {
        getList()
      })
    } else {
      const keys = tableListDom.value.getSelectedRowKeys()
      if (keys.length > 0) {
        templateFieldDelete(keys.join(',')).then(res => {
          getList()
        })
      } else {
        ElMessage.error($t('common.confirm.chooseOne'))
      }
    }
  })
}
const dialogFormVisible = ref(false)
const templateField = ref({
  id: '',
  name: '',
  desc: '',
  type: null,
  value: '',
  cteTime: '',
  uteTime: ''
})
const templateFieldRules = ref({
  name: [{ required: true, message: $t('templateField.itemRules.name'), trigger: 'blur' }],
  desc: [{ required: true, message: $t('templateField.itemRules.desc'), trigger: 'blur' }],
  value: [{ required: true, message: $t('templateField.itemRules.value'), trigger: 'blur' }]
})
function handleDetail(row) {
  dialogFormVisible.value = true
  templateField.value = {}
  if (row && row.id) {
    templateFieldSelect(row.id).then(res => {
      templateField.value = res.data
    })
  }
}

function checkField() {
  templateFieldCheck(templateField.value).then(res => {
    // do nothing
  }, () => {
    ElMessage.error(templateField.value.name + ' ' + $t('common.message.exists'))
  })
}

const templateFieldForm = ref(null)
function handleFormClose() {
  dialogFormVisible.value = false
  templateFieldForm.value.resetFields()
}

function handleFormSubmit() {
  templateFieldForm.value.validate(validate => {
    if (validate) {
      if (templateField.value.id) {
        templateFieldUpdate(templateField.value).then(res => {
          dialogFormVisible.value = false
          getList()
        })
      } else {
        templateFieldInsert(this.templateField).then(res => {
          dialogFormVisible.value = false
          getList()
        })
      }
    }
  })
}
</script>

<style scoped>

</style>
