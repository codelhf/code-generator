<template>
  <table-layout title="模板列表">
    <!-- 左侧树 -->
    <template #tree>
      <tree-list
        ref="treeDom"
        :title="'分组列表'"
        :data="groupList"
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
        <el-form-item :label="$t('template.listQuery.name')">
          <el-input v-model="listQuery.name" :placeholder="$t('template.listQuery.placeholderName')" />
        </el-form-item>
        <el-form-item :label="$t('template.listQuery.desc')">
          <el-input v-model="listQuery.desc" :placeholder="$t('template.listQuery.placeholderDesc')" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :icon="useRenderIcon('search')" @click="handleSearch">{{ $t('template.listButton.search') }}</el-button>
          <el-button type="primary" :icon="useRenderIcon('refresh-left')" @click="handleReset">{{ $t('template.listButton.reset') }}</el-button>
        </el-form-item>
      </el-form>
    </template>
    <!--按钮组-->
    <template #button>
      <el-button type="primary" :icon="useRenderIcon('plus')" @click="handleDetail">{{ $t('template.listButton.add') }}</el-button>
      <el-button type="primary" :icon="useRenderIcon('upload')" @click="handleImport">{{ $t('template.listButton.import') }}</el-button>
      <el-button type="primary" :icon="useRenderIcon('download')" @click="handleExport">{{ $t('template.listButton.export') }}</el-button>
      <el-button type="danger" :icon="useRenderIcon('delete')" @click="handleDelete">{{ $t('template.listButton.delete') }}</el-button>
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
        <el-table-column :label="$t('template.table.name')" align="center">
          <template #default="scope">
            <span>{{ scope.row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('template.table.desc')" align="center">
          <template #default="scope">
            <span>{{ scope.row.desc }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('template.table.type')" align="center">
          <template #default="scope">
            <span v-if="scope.row.type === 1">{{ $t('template.table.type1') }}</span>
            <span v-if="scope.row.type === 2">{{ $t('template.table.type2') }}</span>
            <span v-if="scope.row.type === 3">{{ $t('template.table.type3') }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('template.table.fileType')" align="center">
          <template #default="scope">
            <span>{{ scope.row.fileType }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('template.table.operation')" align="center" width="160">
          <template #default="scope">
            <el-button circle size="small" :icon="useRenderIcon('edit')" @click="handleDetail(scope.row)" />
            <el-button circle size="small" :icon="useRenderIcon('delete')" @click="handleDelete(scope.row)" />
          </template>
        </el-table-column>
      </table-list>
      <!--导入文件弹窗-->
      <el-dialog
        :title="$t('components.uploadExcel.importExcel')"
        v-model="dialogUploadVisible"
        @close="handleUploadClose"
      >
        <upload-excel :header="getHeads" @onSuccess="handleUploadSuccess" />
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="handleUploadClose">{{ $t('components.uploadExcel.formCancel') }}</el-button>
            <el-button type="primary" @click="handleUploadSubmit">{{ $t('components.uploadExcel.formConfirm') }}</el-button>
          </span>
        </template>
      </el-dialog>
      <!--添加分组弹窗-->
      <el-dialog
        :title="group.id ? $t('template.group.editTitle') : $t('template.group.addTitle')"
        v-model="dialogFormVisible2"
        @close="handleForm2Close"
      >
        <el-form ref="groupForm" :model="group" :rules="groupRules" label-width="120px" label-suffix=":">
          <el-form-item :label="$t('template.group.name')" prop="name">
            <el-input v-model="group.name" :placeholder="$t('template.group.placeholderName')" @blur="checkGroup" />
          </el-form-item>
          <el-form-item :label="$t('template.group.desc')" prop="desc">
            <el-input v-model="group.desc" :placeholder="$t('template.group.placeholderDesc')" />
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
import UploadExcel from '@/components/UploadExcel/index.vue'
export default {
  name: 'Template',
  components: {
    TreeList,
    TableList,
    TableLayout,
    UploadExcel
  }
}
</script>

<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { $t } from '@/i18n'
import {
  templateList, templateExport,
  templateImport, templateDelete
} from '@/pages/code-generator/api/template'
import {
  templateGroupList, templateGroupSelect, templateGroupCheck,
  templateGroupInsert, templateGroupUpdate, templateGroupDelete
} from '@/pages/code-generator/api/template-group'

// 左侧树
const groupList = ref([])
function getGroupList() {
  templateGroupList({}).then(res => {
    groupList.value = res.data
    getList()
  })
}
function handleNodeClick(data) {
  listQuery.value.groupId = data.id
  getList()
}
const dialogFormVisible2 = ref(false)
const group = ref({
  id: '',
  name: '',
  desc: ''
})

function handleGroupCreate() {
  dialogFormVisible2.value = true
  group.value = {}
}

function handleGroupUpdate(item) {
  dialogFormVisible2.value = true
  group.value = {}
  if (item && item.id) {
    templateGroupSelect(item.id).then(res => {
      group.value = res.data
    })
  }
}

function handleGroupDelete(item) {
  ElMessageBox.confirm($t('common.confirm.deleteOne'), $t('common.confirm.title'), {
    cancelButtonText: $t('common.confirm.cancel'),
    confirmButtonText: $t('common.confirm.confirm'),
    type: 'warning'
  }).then(() => {
    templateGroupDelete(item.id).then(() => {
      getGroupList()
      getList()
    })
  })
}

function checkGroup() {
  templateGroupCheck(group.value).then(res => {
    // do nothing
  }, () => {
    ElMessage.error(group.value.name + ' ' + $t('common.message.exists'))
  })
}

const groupForm = ref(null)
function handleForm2Close() {
  dialogFormVisible2.value = false
  groupForm.value.resetFields()
}

function handleForm2Submit() {
  groupForm.value.validate(validate => {
    if (validate) {
      if (group.value.id) {
        templateGroupUpdate(group.value).then(res => {
          dialogFormVisible2.value = false
          getGroupList()
          getList()
        })
      } else {
        templateGroupInsert(group.value).then(res => {
          dialogFormVisible2.value = false
          getGroupList()
          getList()
        })
      }
    }
  })
}

// 生命周期
onMounted(() => {
  getGroupList()
})

// 搜索部分
const listQuery = ref({
  pageNum: 1,
  pageSize: 10,
  sort: '',
  groupId: '',
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
    groupId: '',
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
  templateList(listQuery.value).then(res => {
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
      templateDelete(row.id).then(res => {
        getList()
      })
    } else {
      const keys = tableListDom.value.getSelectedRowKeys()
      if (keys.length > 0) {
        templateDelete(keys.join(',')).then(res => {
          getList()
        })
      } else {
        ElMessage.error($t('common.confirm.chooseOne'))
      }
    }
  })
}

// 详情部分
const router = useRouter()
function handleDetail(row) {
  router.push(`/template/detail/${row.id}`)
}

const dialogUploadVisible = ref(false)
const timer = ref(null)
const uploadFile = ref(null)
function handleImport() {
  dialogUploadVisible.value = true
  timer.value = new Date().getTime()
}
function handleUploadClose() {
  dialogUploadVisible.value = false
}
function handleUploadSuccess(file) {
  uploadFile.value = file
}
function handleUploadSubmit() {
  if (uploadFile.value) {
    const formData = new FormData()
    formData.append('uploadFile', uploadFile.value)
    formData.append('sheetName', $t('template.excel.sheetName'))
    formData.append('heads', JSON.stringify(getHeads()))
    templateImport(formData).then(() => {
      dialogUploadVisible.value = false
      getList()
    })
  }
}

function handleExport() {
  const keys = tableListDom.value.getSelectedRowKeys()
  if (keys > 0) {
    const params = { fileName: $t('template.excel.fileName'), sheetName: $t('template.excel.sheetName') }
    params.heads = JSON.stringify(getHeads())
    params.ids = keys.join(',')
    templateExport(params).then(res => {
      const blob = new Blob([res.data])
      const downloadElement = document.createElement('a') // 创建a标签
      // 从response的headers中获取filename, 后端response.setHeader("Content-disposition", "attachment; filename=xxxx.docx") 设置的文件名;
      const contentDisposition = res.headers['content-disposition']
      const pattern = new RegExp('filename=([^;]+\\.[^\\.;]+);*')
      const result = pattern.exec(contentDisposition)
      const filename = result[1]
      if ('msSaveOrOpenBlob' in window.navigator) {
        // 兼容ie
        window.navigator.msSaveOrOpenBlob(blob, filename)
      } else {
        const href = window.URL.createObjectURL(blob) // 创建下载的链接
        downloadElement.style.display = 'none'
        downloadElement.href = href
        downloadElement.download = filename // 下载后文件名
        document.body.appendChild(downloadElement) // 添加a标签到body
        downloadElement.click() // 点击下载
        document.body.removeChild(downloadElement) // 下载完成移除元素
        window.URL.revokeObjectURL(href) // 释放掉blob对象
      }
    })
  } else {
    ElMessage.warning($t('template.excel.warning'))
  }
}

function getHeads() {
  const heads = []
  heads.push({ name: $t('template.table.name'), field: 'name' })
  heads.push({ name: $t('template.table.desc'), field: 'desc' })
  heads.push({ name: $t('template.table.type'), field: 'type' })
  heads.push({ name: $t('template.table.fileType'), field: 'fileType' })
  heads.push({ name: $t('template.table.template'), field: 'template' })
  return heads
}
</script>

<style scoped>

</style>
