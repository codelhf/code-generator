<template>
  <div class="page-container">
    <tree-list
      ref="treeList"
      :data="groupList"
      @itemClick="handleItemClick"
      @addItem="handleGroup"
      @updateItem="handleGroup"
      @deleteItem="handleGroupDelete"
    >
      <el-form :model="listQuery" :inline="true" label-width="120px" label-suffix=":">
        <el-row>
          <el-form-item :label="$t('template.listQuery.name')">
            <el-input v-model="listQuery.name" :placeholder="$t('template.listQuery.placeholderName')" />
          </el-form-item>
          <el-form-item :label="$t('template.listQuery.desc')">
            <el-input v-model="listQuery.desc" :placeholder="$t('template.listQuery.placeholderDesc')" />
          </el-form-item>
        </el-row>
        <el-row style="text-align: center">
          <el-form-item>
            <el-button type="primary" size="mini" icon="el-icon-search" @click="handleFilter">{{ $t('template.listButton.search') }}</el-button>
            <el-button type="primary" size="mini" icon="el-icon-refresh" @click="handleReset">{{ $t('template.listButton.reset') }}</el-button>
            <el-button type="primary" size="mini" icon="el-icon-plus" @click="handleDetail(0)">{{ $t('template.listButton.add') }}</el-button>
            <el-button type="primary" size="mini" icon="el-icon-upload2" @click="handleImport">{{ $t('template.listButton.import') }}</el-button>
            <el-button type="primary" size="mini" icon="el-icon-download" @click="handleExport">{{ $t('template.listButton.export') }}</el-button>
          </el-form-item>
        </el-row>
      </el-form>

      <div style="height: calc(100% - 220px);">
        <el-table
          :key="tableKey"
          v-loading="listLoading"
          :data="list"
          height="100%"
          highlight-current-row
          border="border"
          @selection-change="handleSelectionChange"
        >
          <el-table-column
            type="selection"
            width="40"
          />
          <el-table-column :label="$t('template.table.name')" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.name }}</span>
            </template>
          </el-table-column>
          <el-table-column :label="$t('template.table.desc')" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.desc }}</span>
            </template>
          </el-table-column>
          <el-table-column :label="$t('template.table.type')" align="center">
            <template slot-scope="scope">
              <span v-if="scope.row.type === 1">{{ $t('template.table.type1') }}</span>
              <span v-if="scope.row.type === 2">{{ $t('template.table.type2') }}</span>
              <span v-if="scope.row.type === 3">{{ $t('template.table.type3') }}</span>
            </template>
          </el-table-column>
          <el-table-column :label="$t('template.table.fileType')" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.fileType }}</span>
            </template>
          </el-table-column>
          <el-table-column :label="$t('template.table.operation')" align="center" width="160">
            <template slot-scope="scope">
              <el-button circle size="mini" icon="el-icon-edit" @click="handleDetail(scope.row.id)" />
              <el-button circle size="mini" icon="el-icon-delete" @click="handleDelete(scope.row.id)" />
            </template>
          </el-table-column>
        </el-table>
      </div>

      <pagination v-show="total>0" :total="total" :page.sync="listQuery.pageNum" :limit.sync="listQuery.pageSize" @pagination="getList" />
    </tree-list>

    <el-dialog
      :title="$t('components.uploadExcel.importExcel')"
      :visible.sync="dialogUploadVisible"
      @close="handleUploadClose('uploadExcel')"
    >
      <upload-excel :key="timer" @onSuccess="handleUploadSuccess" />
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleUploadClose('uploadExcel')">{{ $t('components.uploadExcel.formCancel') }}</el-button>
        <el-button type="primary" @click="handleUploadSubmit('uploadExcel')">{{ $t('components.uploadExcel.formConfirm') }}</el-button>
      </span>
    </el-dialog>

    <el-dialog
      :title="group.id ? $t('template.group.editTitle') : $t('template.group.addTitle')"
      :visible.sync="dialogFormVisible2"
      @close="handleForm2Close('groupForm')"
    >
      <el-form ref="groupForm" :model="group" :rules="groupRules()" label-width="120px" label-suffix=":">
        <el-form-item :label="$t('template.group.name')" prop="name">
          <el-input v-model="group.name" :placeholder="$t('template.group.placeholderName')" @blur="checkGroup" />
        </el-form-item>
        <el-form-item :label="$t('template.group.desc')" prop="desc">
          <el-input v-model="group.desc" :placeholder="$t('template.group.placeholderDesc')" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleForm2Close('groupForm')">{{ $t('common.form.cancel') }}</el-button>
        <el-button type="primary" @click="handleForm2Submit('groupForm')">{{ $t('common.form.confirm') }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { templateList, templateExport, templateImport, templateDelete } from '@/api/template'
import { templateGroupList, templateGroupSelect, templateGroupCheck, templateGroupInsert, templateGroupUpdate, templateGroupDelete } from '@/api/template-group'
import Pagination from '@/components/Pagination/index'
import UploadExcel from '@/components/UploadExcel/index'
import TreeList from '@/components/TreeList/index'
export default {
  name: 'Template',
  components: {
    Pagination,
    UploadExcel,
    TreeList
  },
  data() {
    return {
      tableKey: 0,
      list: [],
      total: 0,
      listLoading: true,
      listQuery: {
        pageNum: 1,
        pageSize: 10,
        sort: '-uteTime',
        groupId: '',
        name: '',
        desc: ''
      },
      dialogUploadVisible: false,
      timer: null,
      uploadFile: null,
      multipleSelection: [],
      groupList: [],
      dialogFormVisible2: false,
      group: {
        id: '',
        name: '',
        desc: ''
      }
    }
  },
  created() {
    // this.getList()
    this.getGroupList()
  },
  methods: {
    getList() {
      this.listLoading = true
      templateList(this.listQuery).then(res => {
        this.list = res.data.list
        this.total = res.data.total
        // Just to simulate the time of the request
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    handleFilter() {
      this.listQuery.pageNum = 1
      this.getList()
    },
    handleReset() {
      this.listQuery = {
        pageNum: 1,
        pageSize: 10,
        sort: '-uteTime',
        name: '',
        desc: ''
      }
    },
    handleSelectionChange(val) {
      this.multipleSelection = val.map(item => {
        return item.id
      })
    },
    handleExport() {
      if (this.multipleSelection && this.multipleSelection.length > 0) {
        const params = { fileName: this.$t('template.excel.fileName'), sheetName: this.$t('template.excel.sheetName') }
        params.heads = JSON.stringify(this.getHeads())
        params.ids = this.multipleSelection.join(',')
        templateExport(params).then(res => {
          const url = window.URL.createObjectURL(new Blob([res.data]))
          const link = document.createElement('a')
          link.href = url
          link.setAttribute('download', this.$t('template.excel.fileName') + '.xls')
          document.body.appendChild(link)
          link.click()
        })
      } else {
        this.$message.warning(this.$t('template.excel.warning'))
      }
    },
    getHeads() {
      const heads = []
      heads.push({ name: this.$t('template.table.name'), field: 'name' })
      heads.push({ name: this.$t('template.table.desc'), field: 'desc' })
      heads.push({ name: this.$t('template.table.type'), field: 'type' })
      heads.push({ name: this.$t('template.table.fileType'), field: 'fileType' })
      heads.push({ name: this.$t('template.table.template'), field: 'template' })
      return heads
    },
    handleImport() {
      this.dialogUploadVisible = true
      this.timer = new Date().getTime()
    },
    handleUploadClose() {
      this.dialogUploadVisible = false
    },
    handleUploadSuccess(file) {
      this.uploadFile = file
    },
    handleUploadSubmit() {
      if (this.uploadFile) {
        const formData = new FormData()
        formData.append('uploadFile', this.uploadFile)
        formData.append('sheetName', this.$t('template.excel.sheetName'))
        formData.append('heads', JSON.stringify(this.getHeads()))
        templateImport(formData).then(() => {
          this.dialogUploadVisible = false
          this.getList()
        })
      }
    },
    handleDetail(id) {
      this.$router.push(`/template/detail/${id}`)
    },
    handleDelete(id) {
      this.$confirm(this.$t('common.confirm.deleteOne'), this.$t('common.confirm.title'), {
        cancelButtonText: this.$t('common.confirm.cancel'),
        confirmButtonText: this.$t('common.confirm.confirm'),
        type: 'warning'
      }).then(() => {
        templateDelete(id).then(() => {
          this.getList()
        })
      })
    },
    // TemplateGroup
    groupRules() {
      return {
        name: [{ required: true, message: this.$t('template.groupRules.name'), trigger: 'blur' }]
      }
    },
    getGroupList() {
      templateGroupList({}).then(res => {
        this.groupList = res.data
        this.$refs.treeList.$data.clickedItem = ''
        this.groupList.map(item => {
          item.rightMenu = true
          if (item.id === this.listQuery.groupId) {
            item.active = true
          }
        })
        this.listQuery.groupId = this.groupList[0].id
        this.getList()
      })
    },
    handleItemClick(item) {
      this.groupList = this.groupList.map(group => {
        group.active = false
        if (group.id === item.id) {
          group.active = true
          this.listQuery.groupId = item.id
          this.getList()
        }
        return group
      })
    },
    handleGroup(item) {
      this.dialogFormVisible2 = true
      this.group = {}
      if (item && item.id) {
        templateGroupSelect(item.id).then(res => {
          this.group = res.data
        })
      }
    },
    checkGroup() {
      templateGroupCheck(this.group).then(res => {
        // do nothing
      }, () => {
        this.$message.error(this.group.name + ' ' + this.$t('common.message.exists'))
      })
    },
    handleForm2Submit(formName) {
      this.$refs[formName].validate(validate => {
        if (validate) {
          if (this.group.id) {
            templateGroupUpdate(this.group).then(res => {
              this.dialogFormVisible2 = false
              this.getList()
              this.getGroupList()
            })
          } else {
            templateGroupInsert(this.group).then(res => {
              this.dialogFormVisible2 = false
              this.getList()
              this.getGroupList()
            })
          }
        }
      })
    },
    handleForm2Close(formName) {
      this.dialogFormVisible2 = false
      this.$refs[formName].resetFields()
    },
    handleGroupDelete(item) {
      this.$confirm(this.$t('common.confirm.deleteOne'), this.$t('common.confirm.title'), {
        cancelButtonText: this.$t('common.confirm.cancel'),
        confirmButtonText: this.$t('common.confirm.confirm'),
        type: 'warning'
      }).then(() => {
        templateGroupDelete(item.id).then(() => {
          this.getList()
          this.getLanguageList()
        })
      })
    }
  }
}
</script>

<style scoped>

</style>
