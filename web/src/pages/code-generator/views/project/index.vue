<template>
  <table-layout title="项目列表">
    <!-- 搜索表单 -->
    <template #form>
      <el-form :model="listQuery" :inline="true" label-width="120px" label-suffix=":">
        <el-form-item label="部门名称">
          <el-input v-model="listQuery.name" placeholder="请输入部门名称" />
        </el-form-item>
        <el-form-item label="部门简称">
          <el-input v-model="listQuery.desc" placeholder="请输入部门简称" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :icon="useRenderIcon('search')" @click="handleSearch">{{ $t('project.listButton.search') }}</el-button>
          <el-button type="primary" :icon="useRenderIcon('refresh-left')" @click="handleReset">{{ $t('project.listButton.reset') }}</el-button>
        </el-form-item>
      </el-form>
    </template>
    <!--按钮组-->
    <template #button>
      <el-button type="primary" :icon="useRenderIcon('plus')" @click="handleDetail">{{ $t('project.listButton.add') }}</el-button>
      <el-button type="danger" :icon="useRenderIcon('delete')" @click="handleDelete">批量删除</el-button>
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
        <el-table-column :label="$t('project.table.name')" align="center">
          <template #default="scope">
            <span>{{ scope.row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('project.table.desc')" align="center">
          <template #default="scope">
            <span>{{ scope.row.desc }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('project.table.cteTime')" align="center">
          <template #default="scope">
            <span>{{ scope.row.cteTime }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('project.table.uteTime')" align="center">
          <template #default="scope">
            <span>{{ scope.row.uteTime }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('project.table.operation')" align="center">
          <template #default="scope">
            <el-button circle size="small" :icon="useRenderIcon('position')" @click="handleRun(scope.row)" />
            <el-button circle size="small" :icon="useRenderIcon('edit')" @click="handleDetail(scope.row)" />
            <el-button circle size="small" :icon="useRenderIcon('coin')" @click="handleDatabase(scope.row)" />
            <el-button circle size="small" :icon="useRenderIcon('delete')" @click="handleDelete(scope.row)" />
          </template>
        </el-table-column>
      </table-list>
      <!--project-->
      <el-dialog
        :title="project.id ? $t('project.item.editTitle') : $t('project.item.addTitle')"
        v-model="dialogFormVisible"
        @close="handleFormClose"
      >
        <el-form ref="projectForm" :model="project" :inline="true" :rules="projectRules" label-width="120px" label-suffix=":">
          <el-form-item :label="$t('project.item.name')" prop="name">
            <el-input v-model="project.name" :placeholder="$t('project.item.placeholderName')" />
          </el-form-item>
          <el-form-item :label="$t('project.item.desc')">
            <el-input v-model="project.desc" :placeholder="$t('project.item.placeholderDesc')" />
          </el-form-item>
          <el-form-item :label="$t('project.item.author')" prop="author">
            <el-input v-model="project.author" :placeholder="$t('project.item.placeholderAuthor')" />
          </el-form-item>
          <el-form-item :label="$t('project.item.organization')" prop="organization">
            <el-input v-model="project.organization" :placeholder="$t('project.item.placeholderOrganization')" />
          </el-form-item>
          <el-form-item :label="$t('project.item.dateFormat')" prop="dateFormat">
            <el-input v-model="project.dateFormat" :placeholder="$t('project.item.placeholderDateFormat')" />
          </el-form-item>
          <el-form-item :label="$t('project.item.fileEncode')" prop="fileEncode">
            <el-select v-model="project.fileEncode" :placeholder="$t('project.item.placeholderFileEncode')">
              <el-option label="UTF-8" :value="'UTF-8'" />
              <el-option label="GBK" :value="'GBK'" />
              <el-option label="GB2312" :value="'GB2312'" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('project.item.httpPrefix')" prop="httpPrefix">
            <el-input v-model="project.httpPrefix" :placeholder="$t('project.item.placeholderHttpPrefix')" />
          </el-form-item>
          <el-form-item :label="$t('project.item.generateRemark')" prop="generateRemark">
            <el-select v-model="project.generateRemark" :placeholder="$t('project.item.placeholderGenerateRemark')">
              <el-option :label="$t('project.item.generateRemark1')" :value="1" />
              <el-option :label="$t('project.item.generateRemark2')" :value="2" />
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="handleFormClose">{{ $t('common.form.cancel') }}</el-button>
            <el-button type="primary" @click="handleFormSubmit">{{ $t('common.form.confirm') }}</el-button>
          </span>
        </template>
      </el-dialog>
      <!--project db-->
      <el-dialog
        :title="database.id ? $t('project.database.editTitle') : $t('project.database.addTitle')"
        v-model="dialogForm2Visible"
        @close="handleForm2Close"
      >
        <el-form ref="databaseForm" :model="database" :rules="databaseRules" label-width="120px" label-suffix=":">
          <el-row>
            <el-col :span="10">
              <el-form-item :label="$t('project.database.type')" prop="type">
                <el-select v-model="database.type" :placeholder="$t('project.database.placeholderType')">
                  <el-option v-for="item in dbTypeList" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="10">
              <el-form-item :label="$t('project.database.database')" prop="database">
                <el-input v-model="database.database" :placeholder="$t('project.database.placeholderDatabase')" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="10">
              <el-form-item :label="$t('project.database.host')" prop="host">
                <el-input v-model="database.host" :placeholder="$t('project.database.placeholderHost')" />
              </el-form-item>
            </el-col>
            <el-col :span="10">
              <el-form-item :label="$t('project.database.port')" prop="port">
                <el-input v-model="database.port" :placeholder="$t('project.database.placeholderPort')" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="10">
              <el-form-item :label="$t('project.database.username')" prop="username">
                <el-input v-model="database.username" :placeholder="$t('project.database.placeholderUsername')" />
              </el-form-item>
            </el-col>
            <el-col :span="10">
              <el-form-item :label="$t('project.database.password')" prop="password">
                <el-input v-model="database.password" type="password" :placeholder="$t('project.database.placeholderPassword')" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="10">
              <el-form-item :label="$t('project.database.delimitKeyword')" prop="">
                <el-input v-model="database.delimitKeyword" :placeholder="$t('project.database.placeholderDelimitKeyword')" />
              </el-form-item>
            </el-col>
            <el-col :span="10">
              <el-form-item :label="$t('project.database.languageType')" prop="languageId">
                <el-select v-model="database.languageId" :placeholder="$t('project.database.placeholderLanguageType')">
                  <el-option v-for="item in languageList" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
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
import TableList from '@/components/TableList/index.vue'
import TableLayout from '@/components/TableList/layout.vue'
export default {
  name: 'Project',
  components: {
    TableList,
    TableLayout
  }
}
</script>

<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import {ElMessage, ElMessageBox} from 'element-plus'
import { $t } from '@/i18n'
import { projectList, projectSelect, projectInsert, projectUpdate, projectDelete } from '@/pages/code-generator/api/project'
import { projectDbSelect, projectDbInsert, projectDbUpdate } from '@/pages/code-generator/api/project-db'
import { allDbType } from '@/pages/code-generator/api/type-column'
import { typeLanguageList } from '@/pages/code-generator/api/type-language'

// 生命周期
onMounted(() => {
  getList()
  getAllDbType()
  getLanguageList()
})

// 搜索部分
const listQuery = ref({
  pageNum: 1,
  pageSize: 10,
  sort: '',
  name: '',
  desc: ''
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
    desc: ''
  }
}

// 列表部分
const tableListDom = ref(null)
const listLoading = ref(false)
const list = ref([])
const total = ref(0)

function getList() {
  listLoading.value = true
  projectList(listQuery.value).then(res => {
    list.value = res.data.list
    total.value = res.data.total
    // Just to simulate the time of the request
    setTimeout(() => {
      listLoading.value = false
    }, 1.5 * 1000)
  })
}

const dbTypeList = ref([])
function getAllDbType() {
  allDbType().then(res => {
    dbTypeList.value = res.data
  })
}

const languageList = ref([])
function getLanguageList() {
  typeLanguageList({}).then(res => {
    languageList.value = res.data
  })
}

function handleDelete(row) {
  ElMessageBox.confirm($t('common.confirm.deleteOne'), $t('common.confirm.title'), {
    cancelButtonText: $t('common.confirm.cancel'),
    confirmButtonText: $t('common.confirm.confirm'),
    type: 'warning'
  }).then(() => {
    if (row && row.id) {
      projectDelete(row.id).then(res => {
        getList()
      })
    } else {
      const keys = tableListDom.value.getSelectedRowKeys()
      if (keys.length > 0) {
        projectDelete(keys.join(',')).then(res => {
          getList()
        })
      } else {
        ElMessage.error($t('common.confirm.chooseOne'))
      }
    }
  })
}

const router = useRouter()
// 列表操作部分
function handleRun(row) {
  router.push(`/project/detail/${row.id}`)
}

// 表单部分
const dialogFormVisible = ref(false)
const project = ref({
  id: '',
  name: '',
  desc: '',
  author: '',
  organization: '',
  dateFormat: '',
  fileEncode: 'UTF-8',
  httpPrefix: '',
  generateRemark: 1,
  cteTime: '',
  uteTime: ''
})
const projectRules = ref({
  name: [{ required: true, message: $t('project.itemRules.name'), trigger: 'blur' }],
  desc: [{ required: true, message: $t('project.itemRules.desc'), trigger: 'blur' }],
  author: [{ required: true, message: $t('project.itemRules.author'), trigger: 'blur' }],
  organization: [{ required: true, message: $t('project.itemRules.organization'), trigger: 'blur' }],
  dateFormat: [{ required: true, message: $t('project.itemRules.dateFormat'), trigger: 'blur' }],
  fileEncode: [{ required: true, message: $t('project.itemRules.fileEncode'), trigger: 'blur' }],
  httpPrefix: [{ required: true, message: $t('project.itemRules.httpPrefix'), trigger: 'blur' }],
  generateRemark: [{ required: true, message: $t('project.itemRules.generateRemark'), trigger: 'blur' }]
})

function handleDetail(row) {
  project.value = {
    fileEncode: 'UTF-8',
    generateRemark: 1,
  }
  dialogFormVisible.value = true
  if (row && row.id) {
    projectSelect(row.id).then(res => {
      project.value = res.data
    })
  }
}

const projectForm = ref(null)
function handleFormClose() {
  projectForm.value.resetFields()
  dialogFormVisible.value = false
  project.value = {}
}

function handleFormSubmit() {
  projectForm.value.validate(valid => {
    if (valid) {
      if (!project.value.id) {
        projectInsert(project.value).then(res => {
          dialogFormVisible.value = false
          getList()
        })
      } else {
        projectUpdate(project.value).then(res => {
          dialogFormVisible.value = false
          getList()
        })
      }
    }
  })
}

const dialogForm2Visible = ref(false)
const database = ref({
  id: '',
  projectId: '',
  type: null,
  url: '',
  database: '',
  host: '',
  port: '',
  username: '',
  password: '',
  delimitKeyword: '',
  languageId: ''
})
const databaseRules = ref({
  type: [{ required: true, message: $t('project.databaseRules.type'), trigger: 'blur' }],
  url: [{ required: true, message: $t('project.databaseRules.url'), trigger: 'blur' }],
  database: [{ required: true, message: $t('project.databaseRules.database'), trigger: 'blur' }],
  host: [{ required: true, message: $t('project.databaseRules.host'), trigger: 'blur' }],
  port: [{ required: true, message: $t('project.databaseRules.port'), trigger: 'blur' }],
  username: [{ required: true, message: $t('project.databaseRules.username'), trigger: 'blur' }],
  password: [{ required: true, message: $t('project.databaseRules.password'), trigger: 'blur' }],
  delimitKeyword: [{ required: true, message: $t('project.databaseRules.delimitKeyword'), trigger: 'blur' }],
  languageId: [{ required: true, message: $t('project.databaseRules.languageType'), trigger: 'blur' }]
})
function handleDatabase(row) {
  dialogForm2Visible.value = true
  database.value = { projectId: row.id }
  projectDbSelect(row.id).then(res => {
    database.value = res.data
    database.value.projectId = row.id
    if (database.value.type === 0) {
      database.value.type = null
    }
    if (database.value.languageId === 0) {
      database.value.languageId = null
    }
  })
}

const databaseForm = ref(null)
function handleForm2Close() {
  dialogForm2Visible.value = false
  databaseForm.value.resetFields()
}

function handleForm2Submit() {
  databaseForm.value.validate(valid => {
    if (valid) {
      handleDbTypeUrl(database.value.type)
      if (database.value.id) {
        projectDbUpdate(database.value).then(res => {
          dialogForm2Visible.value = false
          getList()
        })
      } else {
        projectDbInsert(database.value).then(res => {
          dialogForm2Visible.value = false
          getList()
        })
      }
    }
  })
}
function handleDbTypeUrl(type) {
  const db = database.value
  if (type === 1) {
    db.url = db.username + ':' + db.password + '@tcp(' + db.host + ':' + db.port + ')/' + db.database + '?charset=utf8'
  } else if (type === 2) {
    db.url = db.username + '/' + db.password + '@' + db.host + ':' + db.port + '/' + db.database + '?charset=utf8'
  } else if (type === 3) {
    db.url = 'user=' + db.username + ' password=' + db.password + ' host=' + db.host + ' port=' + db.port + ' dbname=' + db.database + ' sslmode=disable'
  }
  database.value = db
}
</script>

<style scoped>

</style>
