<template>
  <table-layout title="模板字段">
    <!-- 左侧树 -->
    <template #tree>
      <tree-list
        ref="treeDom"
        :title="'分组列表'"
        :data="languageList"
        :disable-right-click="false"
        :show-check-box="false"
        :enable-click-check="true"
        :enable-parent-check="false"
        @node-click="handleNodeClick"
        @append-node="handleGroupCreate"
        @update-node="handleGroupUpdate"
        @delete-node="handleGroupDelete"
      />
    </template>
    <!-- 搜索表单 -->
    <template #form>
      <el-form :model="listQuery" :inline="true" label-width="120px" label-suffix=":">
        <el-form-item :label="$t('typeColumn.listQuery.dbType')">
          <el-select v-model="listQuery.dbType" :placeholder="$t('typeColumn.listQuery.placeholderDbType')" style="width: 200px">
            <el-option v-for="(item) in dbTypeList" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('typeColumn.listQuery.columnType')">
          <el-input v-model="listQuery.columnType" :placeholder="$t('typeColumn.listQuery.placeholderColumnType')" />
        </el-form-item>
        <el-form-item :label="$t('typeColumn.listQuery.fieldType')">
          <el-input v-model="listQuery.fieldType" :placeholder="$t('typeColumn.listQuery.placeholderFieldType')" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :icon="useRenderIcon('search')" @click="handleSearch">{{ $t('typeColumn.listButton.search') }}</el-button>
          <el-button type="primary" :icon="useRenderIcon('refresh-left')" @click="handleReset">{{ $t('typeColumn.listButton.reset') }}</el-button>
        </el-form-item>
      </el-form>
    </template>
    <!--按钮组-->
    <template #button>
      <el-button type="primary" :icon="useRenderIcon('plus')" @click="handleDetail">{{ $t('typeColumn.listButton.add') }}</el-button>
      <el-button type="danger" :icon="useRenderIcon('delete')" @click="handleDelete">{{ $t('typeColumn.listButton.delete') }}</el-button>
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
        <el-table-column :label="$t('typeColumn.table.dbType')" align="center">
          <template #default="scope">
            <span>{{ scope.row.dbName }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('typeColumn.table.regexpType')" align="center">
          <template #default="scope">
            <span v-if="scope.row.regexpType === 1">{{ $t('typeColumn.item.regexpType1') }}</span>
            <span v-if="scope.row.regexpType === 2">{{ $t('typeColumn.item.regexpType2') }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('typeColumn.table.columnType')" align="center">
          <template #default="scope">
            <span>{{ scope.row.columnType }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('typeColumn.table.fieldType')" align="center">
          <template #default="scope">
            <span>{{ scope.row.fieldType }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('typeColumn.table.jdbcType')" align="center">
          <template #default="scope">
            <span>{{ scope.row.jdbcType }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('typeColumn.table.operation')" align="center" width="160">
          <template #default="scope">
            <el-button circle size="small" :icon="useRenderIcon('edit')" @click="handleDetail(scope.row)" />
            <el-button circle size="small" :icon="useRenderIcon('delete')" @click="handleDelete(scope.row)" />
          </template>
        </el-table-column>
      </table-list>
      <!--字段映射-->
      <el-dialog
        :title="typeColumn.id ? $t('typeColumn.item.editTitle') : $t('typeColumn.item.addTitle')"
        v-model="dialogFormVisible"
        @close="handleFormClose"
      >
        <el-form ref="typeColumnForm" :model="typeColumn" :rules="typeColumnRules" label-width="120px" label-suffix=":">
          <el-form-item :label="$t('typeColumn.item.languageType')" prop="languageId">
            <el-select v-model="typeColumn.languageId" :placeholder="$t('typeColumn.item.placeholderLanguageType')">
              <el-option v-for="(item) in languageList" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('typeColumn.item.dbType')" prop="dbType">
            <el-select v-model="typeColumn.dbType" :placeholder="$t('typeColumn.item.placeholderDbType')">
              <el-option v-for="(item) in dbTypeList" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('typeColumn.item.regexpType')" prop="regexpType">
            <el-select v-model="typeColumn.regexpType" :placeholder="$t('typeColumn.item.placeholderRegexpType')">
              <el-option :label="$t('typeColumn.item.regexpType1')" :value="1" />
              <el-option :label="$t('typeColumn.item.regexpType2')" :value="2" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('typeColumn.item.columnType')" prop="columnType">
            <el-input v-model="typeColumn.columnType" :placeholder="$t('typeColumn.item.placeholderColumnType')" @blur="checkTypeColumn" />
          </el-form-item>
          <el-form-item :label="$t('typeColumn.item.fieldType')" prop="fieldType">
            <el-input v-model="typeColumn.fieldType" :placeholder="$t('typeColumn.item.placeholderFieldType')" />
          </el-form-item>
          <el-form-item :label="$t('typeColumn.item.jdbcType')" prop="jdbcType">
            <el-input v-model="typeColumn.jdbcType" :placeholder="$t('typeColumn.item.placeholderJdbcType')" />
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="handleFormClose">{{ $t('common.form.cancel') }}</el-button>
            <el-button type="primary" @click="handleFormSubmit">{{ $t('common.form.confirm') }}</el-button>
          </span>
        </template>
      </el-dialog>
      <!--语言类型-->
      <el-dialog
        :title="language.id ? $t('typeColumn.language.editTitle') : $t('typeColumn.language.addTitle')"
        v-model="dialogFormVisible2"
        @close="handleForm2Close"
      >
        <el-form ref="languageForm" :model="language" :rules="languageRules" label-width="120px" label-suffix=":">
          <el-form-item :label="$t('typeColumn.language.name')" prop="name">
            <el-input v-model="language.name" :placeholder="$t('typeColumn.language.placeholderName')" @blur="checkLanguage" />
          </el-form-item>
          <el-form-item :label="$t('typeColumn.language.desc')" prop="desc">
            <el-input v-model="language.desc" :placeholder="$t('typeColumn.language.placeholderDesc')" />
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="handleForm2Close">{{ $t('common.form.cancel') }}</el-button>
            <el-button type="primary" @click="handleForm2Submit">{{ $t('common.form.confirm') }}</el-button>
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
  name: 'TypeColumn',
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
  typeColumnList, typeColumnSelect, typeColumnCheck,
  typeColumnInsert, typeColumnUpdate, typeColumnDelete
} from '@/pages/code-generator/api/type-column'
import {
  typeLanguageList, typeLanguageSelect, typeLanguageCheck,
  typeLanguageInsert, typeLanguageUpdate, typeLanguageDelete
} from '@/pages/code-generator/api/type-language'
import { allDbType } from '@/pages/code-generator/api/type-column'
// 左侧树
const languageList = ref([])
function getGroupList() {
  typeLanguageList({}).then(res => {
    languageList.value = res.data
    getList()
  })
}
function handleNodeClick(data) {
  listQuery.value.languageId = data.id
  getList()
}
const dialogFormVisible2 = ref(false)
const language = ref({
  id: '',
  name: '',
  desc: ''
})
const languageRules = ref({
  name: [{ required: true, message: $t('typeColumn.languageRules.name'), trigger: 'blur' }]
})
function handleGroupCreate() {
  dialogFormVisible2.value = true
  language.value = {}
}

function handleGroupUpdate(item) {
  dialogFormVisible2.value = true
  language.value = {}
  if (item && item.id) {
    typeLanguageSelect(item.id).then(res => {
      language.value = res.data
    })
  }
}

function handleGroupDelete(item) {
  ElMessageBox.confirm($t('common.confirm.deleteOne'), $t('common.confirm.title'), {
    cancelButtonText: $t('common.confirm.cancel'),
    confirmButtonText: $t('common.confirm.confirm'),
    type: 'warning'
  }).then(() => {
    typeLanguageDelete(item.id).then(() => {
      getGroupList()
      getList()
    })
  })
}

function checkGroup() {
  typeLanguageCheck(language.value).then(res => {
    // do nothing
  }, () => {
    ElMessage.error(language.value.name + ' ' + $t('common.message.exists'))
  })
}

const languageForm = ref(null)
function handleForm2Close() {
  dialogFormVisible2.value = false
  languageForm.value.resetFields()
}

function handleForm2Submit() {
  languageForm.value.validate(validate => {
    if (validate) {
      if (language.value.id) {
        typeLanguageUpdate(language.value).then(res => {
          dialogFormVisible2.value = false
          getGroupList()
          getList()
        })
      } else {
        typeLanguageInsert(language.value).then(res => {
          dialogFormVisible2.value = false
          getGroupList()
          getList()
        })
      }
    }
  })
}

const dbTypeList = ref([])
function getAllDbType() {
  allDbType().then(res => {
    dbTypeList.value = res.data
  })
}

// 生命周期
onMounted(() => {
  getGroupList()
  getAllDbType()
})

// 搜索部分
const listQuery = ref({
  pageNum: 1,
  pageSize: 10,
  sort: '',
  languageId: '',
  dbType: null,
  columnType: '',
  fieldType: ''
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
    languageId: '',
    dbType: null,
    columnType: '',
    fieldType: ''
  }
}

// 列表部分
const tableListDom = ref(null)
const listLoading = ref(false)
const list = ref([])
const total = ref(0)

function getList() {
  listLoading.value = true
  typeColumnList(listQuery.value).then(res => {
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
      typeColumnDelete(row.id).then(res => {
        getList()
      })
    } else {
      const keys = tableListDom.value.getSelectedRowKeys()
      if (keys.length > 0) {
        typeColumnDelete(keys.join(',')).then(res => {
          getList()
        })
      } else {
        ElMessage.error($t('common.confirm.chooseOne'))
      }
    }
  })
}

const dialogFormVisible = ref(false)
const typeColumn = ref({
  id: '',
  languageId: '',
  dbType: null,
  dbName: '',
  regexpType: '',
  columnType: '',
  fieldType: '',
  jdbcType: ''
})
const typeColumnRules = ref({
  languageId: [{ required: true, message: $t('typeColumn.itemRules.languageType'), trigger: 'blur' }],
  dbType: [{ required: true, message: $t('typeColumn.itemRules.dbType'), trigger: 'blur' }],
  regexpType: [{ required: true, message: $t('typeColumn.itemRules.regexpType'), trigger: 'blur' }],
  columnType: [{ required: true, message: $t('typeColumn.itemRules.columnType'), trigger: 'blur' }],
  fieldType: [{ required: true, message: $t('typeColumn.itemRules.fieldType'), trigger: 'blur' }],
  jdbcType: [{ required: true, message: $t('typeColumn.itemRules.jdbcType'), trigger: 'blur' }]
})
function handleDetail(row) {
  dialogFormVisible.value = true
  typeColumn.value = {}
  if (row && row.id) {
    typeColumnSelect(row.id).then(res => {
      typeColumn.value = res.data
      if (typeColumn.value.dbType === 0) {
        typeColumn.value.dbType = null
      }
      if (typeColumn.value.languageId === 0) {
        typeColumn.value.languageId = null
      }
    })
  }
}

function checkTypeColumn() {
  typeColumnCheck(typeColumn.value).then(res => {
    // do nothing
  }, () => {
    ElMessage.error(typeColumn.value.columnType + ('common.message.exists'))
  })
}

const typeColumnForm = ref(null)
function handleFormClose() {
  dialogFormVisible.value = false
  typeColumnForm.value.resetFields()
}

function handleFormSubmit() {
  typeColumnForm.value.validate(validate => {
    if (validate) {
      if (typeColumn.value.id) {
        typeColumnUpdate(typeColumn.value).then(res => {
          dialogFormVisible.value = false
          getList()
        })
      } else {
        typeColumnInsert(typeColumn.value).then(res => {
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
