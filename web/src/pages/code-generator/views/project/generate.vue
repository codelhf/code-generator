<template>
  <table-layout title="模板列表">
    <!-- 左侧树 -->
    <template #tree>
      <tree-list
        ref="treeDom"
        :title="'数据库表'"
        :data="tableList"
        :disable-right-click="true"
        :show-check-box="false"
        :enable-click-check="true"
        :enable-parent-check="false"
        @node-click="handleTableClick"
      />
    </template>
    <template #table>
      <!--template-->
      <div style="width: 100%; height: calc((100% - 70px) / 2);">
        <el-table
          ref="templateListDom"
          :key="tableKey"
          v-loading="listLoading"
          :data="list.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
          height="100%"
          border="border"
        >
          <el-table-column :label="$t('projectTemplate.table.name')" align="center">
            <template #default="scope">
              <span>{{ scope.row.name }}</span>
            </template>
          </el-table-column>
          <el-table-column :label="$t('projectTemplate.table.directory')" align="center">
            <template #default="scope">
              <span>{{ scope.row.directory }}</span>
            </template>
          </el-table-column>
          <el-table-column :label="$t('projectTemplate.table.packageName')" align="center">
            <template #default="scope">
              <span>{{ scope.row.packageName }}</span>
            </template>
          </el-table-column>
          <el-table-column :label="$t('projectTemplate.table.fileSuffix')" align="center">
            <template #default="scope">
              <span>{{ scope.row.fileSuffix }}</span>
            </template>
          </el-table-column>
          <el-table-column :label="$t('projectTemplate.table.isGenerate')" align="center">
            <template #default="scope">
              <el-switch
                v-model="scope.row.isGenerate"
                :active-value="1"
                :inactive-value="2"
                :active-text="$t('projectTemplate.switch.activeText')"
                :inactive-text="$t('projectTemplate.switch.inactiveText')"
                active-color="#13ce66"
                inactive-color="#ff4949"
                @change="handleTableSwitch(scope.row)"
              />
            </template>
          </el-table-column>
          <el-table-column :label="$t('projectTemplate.table.isOverride')" align="center">
            <template #default="scope">
              <el-switch
                v-model="scope.row.isOverride"
                :active-value="1"
                :inactive-value="2"
                :active-text="$t('projectTemplate.switch.activeText')"
                :inactive-text="$t('projectTemplate.switch.inactiveText')"
                active-color="#13ce66"
                inactive-color="#ff4949"
                @change="handleTableSwitch(scope.row)"
              />
            </template>
          </el-table-column>
          <el-table-column :label="$t('projectTemplate.table.operation')" align="center">
            <template #header #default="scope">
              <el-input v-model="search" size="small" :placeholder="$t('projectTemplate.table.operation')" />
            </template>
            <template #default="scope">
              <el-button circle size="small" :icon="useRenderIcon('edit')" @click="handleDetail(scope.row.id)" />
              <el-button circle size="small" :icon="useRenderIcon('delete')" @click="handleDelete(scope.row.id)" />
            </template>
          </el-table-column>
        </el-table>
      </div>
      <!--button-->
      <el-form ref="currentTableForm" :model="currentTable" :inline="true" label-width="80px" style="margin-top: 20px;">
        <el-form-item :label="$t('projectGenerate.currentTable.domainName')">
          <el-input v-model="currentTable.domainName" :placeholder="$t('projectGenerate.currentTable.placeholderDomainName')" />
        </el-form-item>
        <el-form-item :label="$t('projectGenerate.currentTable.domainDesc')">
          <el-input v-model="currentTable.domainDesc" :placeholder="$t('projectGenerate.currentTable.placeholderDomainDesc')" />
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" :icon="useRenderIcon('video-play')" :loading="loading" @click="handleGenerate">{{ $t('projectGenerate.button.generateCode') }}</el-button>
          <el-button type="primary" :icon="useRenderIcon('plus')" @click="handleDetail">{{ $t('projectGenerate.button.addTemplate') }}</el-button>
        </el-form-item>
      </el-form>
      <!--result-->
      <div style="width: 100%; height: calc((100% - 70px) / 2);">
        <el-input type="textarea" v-model="generateResult" />
      </div>

      <!-- templateDetail -->
      <el-dialog
        :title="projectTemplate.id ? $t('projectTemplate.item.editTitle') : $t('projectTemplate.item.addTitle')"
        v-model="dialogFormVisible"
        @close="handleFormClose"
      >
        <el-form ref="projectTemplateForm" :model="projectTemplate" :rules="projectTemplateRules" label-width="130px" label-suffix=":">
          <el-form-item :label="$t('projectTemplate.item.groupId')" prop="groupId">
            <el-select v-model="projectTemplate.groupId" filterable :placeholder="$t('projectTemplate.item.placeholderGroup')" @change="handleGroup">
              <el-option v-for="item in groupList" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('projectTemplate.item.templateId')" prop="templateId">
            <el-select v-model="projectTemplate.templateId" filterable :placeholder="$t('projectTemplate.item.placeholderName')">
              <el-option v-for="item in templateList" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('projectTemplate.item.directory')" prop="directory">
            <!--<input id="file" type="file" hidden webkitdirectory @change="fileChange">-->
            <el-input v-model="projectTemplate.directory" :placeholder="$t('projectTemplate.item.placeholderDirectory')">
              <el-button slot="append" icon="el-icon-folder" type="success" @click="" />
            </el-input>
          </el-form-item>
          <el-form-item :label="$t('projectTemplate.item.packageName')" prop="packageName">
            <el-input v-model="projectTemplate.packageName" :placeholder="$t('projectTemplate.item.placeholderPackageName')" />
          </el-form-item>
          <el-form-item :label="$t('projectTemplate.item.fileSuffix')" prop="fileSuffix">
            <el-input v-model="projectTemplate.fileSuffix" :placeholder="$t('projectTemplate.item.placeholderFileSuffix')" />
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
import TableLayout from '@/components/TableList/layout.vue'
export default {
  name: 'Generate',
  components: {
    TreeList,
    TableList,
    TableLayout
  }
}
</script>

<script setup>
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { $t } from '@/i18n'
import { projectLastId } from '@/pages/code-generator/api/project'
import { projectTableList, projectTableSelect, projectTableGenerate } from '@/pages/code-generator/api/project-table'
import { projectTemplateList, projectTemplateSelect, projectTemplateInsert, projectTemplateUpdate, projectTemplateDelete } from '@/pages/code-generator/api/project-template'
import { templateGroupList } from '@/pages/code-generator/api/template-group'
import { templateList as fetchTemplateList } from '@/pages/code-generator/api/template'

// 生命周期
const route = useRoute()
const router = useRouter()
const projectId = ref('')
onMounted(() => {
  if (route.params.id) {
    projectId.value = route.params.id
    getTableList()
    getList()
  } else {
    projectLastId().then(res => {
      if (!res.data) {
        router.push('/project/index')
      }
      projectId.value = res.data
      getTableList()
      getList()
    })
  }
  getGroupList()
  getTemplateList()
})

// 数据库表
const tableList = ref([])
function getTableList() {
  const params = { projectId: projectId.value }
  projectTableList(params).then(res => {
    tableList.value = res.data
    tableList.value.map(item => {
      item.name = item.tableName
      item.table = item.tableType === 1
    })
  })
}

const currentTable = ref({
  id: '',
  projectId: '',
  tableName: '',
  tableType: null,
  domainName: '',
  domainDesc: ''
})
function handleTableClick(table) {
  projectTableSelect(projectId.value, table.tableName).then(res => {
    if (res.data.tableName) {
      currentTable.value = res.data
    } else {
      currentTable.value = table
    }
  })
}
const loading = ref(false)
const generateResult = ref('')
const templateListDom = ref(null)
function handleGenerate() {
  if (!currentTable.value.tableName) {
    ElMessage.warning($t('projectGenerate.button.tableWarning'))
    return
  }
  const templateList = templateListDom.value.data
  if (templateList.length === 0) {
    ElMessage.warning($t('projectGenerate.button.templateWarning'))
    return
  }
  loading.value = true
  generateResult.value = ''
  projectTableGenerate(projectId.value, [currentTable.value]).then(res => {
    setTimeout(() => {
      loading.value = false
      generateResult.value = res.data
    }, 1.5 * 1000)
  })
}

// template 部分
const tableKey = ref(0)
const list = ref([])
const total = ref(0)
const listLoading = ref(false)
const search = ref('')
function getList() {
  const params = { projectId: projectId.value }
  projectTemplateList(params).then(res => {
    list.value = res.data.list
    total.value = res.data.total
    // Just to simulate the time of the request
    setTimeout(() => {
      listLoading.value = false
    }, 1.5 * 1000)
  })
}

const groupList = ref([])
function getGroupList() {
  templateGroupList({}).then(res => {
    groupList.value = res.data
  })
}

const templateAll = ref([])
function getTemplateList() {
  fetchTemplateList({}).then(res => {
    templateAll.value = res.data.list
  })
}

const templateList = ref([])
function handleGroup() {
  templateList.value = templateAll.value.filter(item => {
    if (item.groupId === projectTemplate.value.groupId) {
      return item
    }
  })
}

function handleTableSwitch(row) {
  projectTemplateUpdate(row).then(res => {
    getList()
  })
}

const dialogFormVisible = ref(false)
const projectTemplate = ref({
  id: '',
  projectId: '',
  groupId: '',
  templateId: '',
  directory: '',
  packageName: '',
  fileSuffix: '',
  isGenerate: 1,
  isOverride: 1,
  name: '',
  type: null,
  fileType: ''
})
const projectTemplateRules = ref({
  projectId: [{ required: true, message: $t('projectTemplate.itemRules.projectId'), trigger: 'blur' }],
  groupId: [{ required: true, message: $t('projectTemplate.itemRules.groupId'), trigger: 'blur' }],
  templateId: [{ required: true, message: $t('projectTemplate.itemRules.templateId'), trigger: 'blur' }],
  directory: [{ required: true, message: $t('projectTemplate.itemRules.directory'), trigger: 'blur' }],
  packageName: [{ required: true, message: $t('projectTemplate.itemRules.packageName'), trigger: 'blur' }],
  fileSuffix: [{ required: true, message: $t('projectTemplate.itemRules.fileSuffix'), trigger: 'blur' }],
  isGenerate: [{ required: true, message: $t('projectTemplate.itemRules.isGenerate'), trigger: 'blur' }],
  isOverride: [{ required: true, message: $t('projectTemplate.itemRules.isOverride'), trigger: 'blur' }]
})
function handleDetail(id) {
  dialogFormVisible.value = true
  // this.projectTemplate = {
  //   projectId: this.projectId,
  //   isGenerate: 1,
  //   isOverride: 1
  // }
  // if (id) {
  projectTemplateSelect(id).then(res => {
    templateList.value = templateAll.value.filter(item => {
      if (item.groupId === res.data.groupId) {
        return item
      }
    })
    projectTemplate.value = res.data
    if (!res.data || !res.data.id) {
      projectTemplate.value.projectId = projectId.value
      projectTemplate.value.isGenerate = 1
      projectTemplate.value.isOverride = 1
    }
  })
  // }
}

const projectTemplateForm = ref(null)
function handleFormSubmit() {
  projectTemplateForm.value.validate(validate => {
    if (validate) {
      if (projectTemplate.value.id) {
        projectTemplateUpdate(projectTemplate.value).then(res => {
          dialogFormVisible.value = false
          getList()
        }, () => {
          let templateName = ''
          templateList.value.map(item => {
            if (item.id === projectTemplate.value.templateId) {
              templateName = item.name
            }
          })
          ElMessage.error(templateName + ' ' + $t('common.message.exists'))
        })
      } else {
        projectTemplateInsert(projectTemplate.value).then(res => {
          dialogFormVisible.value = false
          getList()
        }, () => {
          let templateName = ''
          templateList.value.map(item => {
            if (item.id === projectTemplate.value.templateId) {
              templateName = item.name
            }
          })
          ElMessage.error(templateName + ' ' + $t('common.message.exists'))
        })
      }
    }
  })
}
function handleFormClose() {
  dialogFormVisible.value = false
  projectTemplateForm.value.resetFields()
}
function handleDelete(id) {
  ElMessageBox.confirm($t('common.confirm.deleteOne'), $t('common.confirm.title'), {
    cancelButtonText: $t('common.confirm.cancel'),
    confirmButtonText: $t('common.confirm.confirm'),
    type: 'warning'
  }).then(() => {
    projectTemplateDelete(id).then(() => {
      getList()
    })
  })
}
</script>

<style lang="scss" scoped>
::v-deep {
  .el-textarea {
    height: 100%;

    .el-textarea__inner {
      height: 100%;
    }
  }
}
</style>
