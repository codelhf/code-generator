<template>
  <div style="width: 100%; height: calc((100vh - 124px) / 2);">
    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="list.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
      height="100%"
      highlight-current-row
      border="border"
    >
      <el-table-column :label="$t('projectTemplate.table.name')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('projectTemplate.table.directory')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.directory }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('projectTemplate.table.packageName')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.packageName }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('projectTemplate.table.fileSuffix')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.fileSuffix }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('projectTemplate.table.isGenerate')" align="center">
        <template slot-scope="scope">
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
        <template slot-scope="scope">
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
        <template slot="header" slot-scope="scope">
          <el-input v-model="search" size="mini" :placeholder="$t('projectTemplate.table.operation')" />
        </template>
        <template slot-scope="scope">
          <el-button circle size="mini" icon="el-icon-edit" @click="handleDetail(scope.row.id)" />
          <el-button circle size="mini" icon="el-icon-delete" @click="handleDelete(scope.row.id)" />
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      :title="projectTemplate.id ? $t('projectTemplate.item.editTitle') : $t('projectTemplate.item.addTitle')"
      :visible.sync="dialogFormVisible"
      @close="handleFormClose('projectTemplateForm')"
    >
      <el-form ref="projectTemplateForm" :model="projectTemplate" :rules="projectTemplateRules()" label-width="130px" label-suffix=":">
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
            <el-button slot="append" icon="el-icon-folder" type="success" @click="btnChange"></el-button>
          </el-input>
        </el-form-item>
        <el-form-item :label="$t('projectTemplate.item.packageName')" prop="packageName">
          <el-input v-model="projectTemplate.packageName" :placeholder="$t('projectTemplate.item.placeholderPackageName')" />
        </el-form-item>
        <el-form-item :label="$t('projectTemplate.item.fileSuffix')" prop="fileSuffix">
          <el-input v-model="projectTemplate.fileSuffix" :placeholder="$t('projectTemplate.item.placeholderFileSuffix')" />
        </el-form-item>
        <el-form-item :label="$t('projectTemplate.item.isGenerate')" prop="isGenerate">
          <el-switch
            v-model="projectTemplate.isGenerate"
            :active-value="1"
            :inactive-value="2"
            :active-text="$t('projectTemplate.switch.activeText')"
            :inactive-text="$t('projectTemplate.switch.inactiveText')"
            active-color="#13ce66"
            inactive-color="#ff4949"
          />
        </el-form-item>
        <el-form-item :label="$t('projectTemplate.item.isOverride')" prop="isOverride">
          <el-switch
            v-model="projectTemplate.isOverride"
            :active-value="1"
            :inactive-value="2"
            :active-text="$t('projectTemplate.switch.activeText')"
            :inactive-text="$t('projectTemplate.switch.inactiveText')"
            active-color="#13ce66"
            inactive-color="#ff4949"
          />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleFormClose('projectTemplateForm')">{{ $t('common.form.cancel') }}</el-button>
        <el-button type="primary" @click="handleFormSubmit('projectTemplateForm')">{{ $t('common.form.confirm') }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { projectTemplateList, projectTemplateSelect, projectTemplateInsert, projectTemplateUpdate, projectTemplateDelete } from '@/api/project-template'
import { templateGroupList } from '@/api/template-group'
import { templateList } from '@/api/template'
export default {
  name: 'ProjectTemplate',
  props: {
    projectId: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      tableKey: 0,
      list: [],
      listLoading: true,
      search: '',
      groupList: [],
      templateAll: [],
      templateList: [],
      dialogFormVisible: false,
      projectTemplate: {
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
      }
    }
  },
  created() {
    this.getList()
    this.getGroupList()
    this.getTemplateList()
  },
  methods: {
    projectTemplateRules() {
      return {
        projectId: [{ required: true, message: this.$t('projectTemplate.itemRules.projectId'), trigger: 'blur' }],
        groupId: [{ required: true, message: this.$t('projectTemplate.itemRules.groupId'), trigger: 'blur' }],
        templateId: [{ required: true, message: this.$t('projectTemplate.itemRules.templateId'), trigger: 'blur' }],
        directory: [{ required: true, message: this.$t('projectTemplate.itemRules.directory'), trigger: 'blur' }],
        packageName: [{ required: true, message: this.$t('projectTemplate.itemRules.packageName'), trigger: 'blur' }],
        fileSuffix: [{ required: true, message: this.$t('projectTemplate.itemRules.fileSuffix'), trigger: 'blur' }],
        isGenerate: [{ required: true, message: this.$t('projectTemplate.itemRules.isGenerate'), trigger: 'blur' }],
        isOverride: [{ required: true, message: this.$t('projectTemplate.itemRules.isOverride'), trigger: 'blur' }]
      }
    },
    getList() {
      const params = { projectId: this.projectId }
      projectTemplateList(params).then(res => {
        this.list = res.data.list
        this.total = res.data.total
        // Just to simulate the time of the request
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    getGroupList() {
      templateGroupList({}).then(res => {
        this.groupList = res.data
      })
    },
    handleGroup() {
      this.templateList = this.templateAll.filter(item => {
        if (item.groupId === this.projectTemplate.groupId) {
          return item
        }
      })
    },
    getTemplateList() {
      templateList({}).then(res => {
        this.templateAll = res.data.list
      })
    },
    handleTableSwitch(row) {
      projectTemplateUpdate(row).then(res => {
        this.getList()
      })
    },
    handleDetail(id) {
      this.dialogFormVisible = true
      this.projectTemplate = {
        projectId: this.projectId,
        isGenerate: 1,
        isOverride: 1
      }
      if (id) {
        projectTemplateSelect(id).then(res => {
          this.templateList = this.templateAll.filter(item => {
            if (item.groupId === res.data.groupId) {
              return item
            }
          })
          this.projectTemplate = res.data
          this.projectTemplate.projectId = this.projectId
        })
      }
    },
    // fileChange(e) {
    //   try {
    //     const directory = document.getElementById('file')
    //     console.log(directory)
    //     if (directory === null) return
    //     console.log(directory.files[0].path)
    //     this.projectTemplate.directory = directory.files[0].path
    //   } catch (error) {
    //     console.debug('choice file err:', error)
    //   }
    // },
    // btnChange() {
    //   const file = document.getElementById('file')
    //   file.click()
    // },
    handleFormSubmit(formName) {
      this.$refs[formName].validate(validate => {
        if (validate) {
          if (this.projectTemplate.id) {
            projectTemplateUpdate(this.projectTemplate).then(res => {
              this.dialogFormVisible = false
              this.getList()
            }, () => {
              let templateName = ''
              this.templateList.map(item => {
                if (item.id === this.projectTemplate.templateId) {
                  templateName = item.name
                }
              })
              this.$message.error(templateName + ' ' + this.$t('common.message.exists'))
            })
          } else {
            projectTemplateInsert(this.projectTemplate).then(res => {
              this.dialogFormVisible = false
              this.getList()
            }, () => {
              let templateName = ''
              this.templateList.map(item => {
                if (item.id === this.projectTemplate.templateId) {
                  templateName = item.name
                }
              })
              this.$message.error(templateName + ' ' + this.$t('common.message.exists'))
            })
          }
        }
      })
    },
    handleFormClose(formName) {
      this.dialogFormVisible = false
      this.$refs[formName].resetFields()
    },
    handleDelete(id) {
      this.$confirm(this.$t('common.confirm.deleteOne'), this.$t('common.confirm.title'), {
        cancelButtonText: this.$t('common.confirm.cancel'),
        confirmButtonText: this.$t('common.confirm.confirm'),
        type: 'warning'
      }).then(() => {
        projectTemplateDelete(id).then(() => {
          this.getList()
        })
      })
    }
  }
}
</script>

<style scoped>

</style>
