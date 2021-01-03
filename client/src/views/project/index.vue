<template>
  <div class="page-container">
    <el-form :model="listQuery" :inline="true" label-width="120px" label-suffix=":">
      <el-row>
        <el-form-item :label="$t('project.listQuery.name')">
          <el-input v-model="listQuery.name" :placeholder="$t('project.listQuery.placeholderName')" />
        </el-form-item>
        <el-form-item :label="$t('project.listQuery.desc')">
          <el-input v-model="listQuery.desc" :placeholder="$t('project.listQuery.placeholderDesc')" />
        </el-form-item>
      </el-row>
      <el-row style="text-align: center">
        <el-form-item>
          <el-button type="primary" size="mini" icon="el-icon-search" @click="handleFilter">{{ $t('project.listButton.search') }}</el-button>
          <el-button type="primary" size="mini" icon="el-icon-refresh" @click="handleReset">{{ $t('project.listButton.reset') }}</el-button>
          <el-button type="primary" size="mini" icon="el-icon-plus" @click="handleDetail()">{{ $t('project.listButton.add') }}</el-button>
        </el-form-item>
      </el-row>
    </el-form>

    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      highlight-current-row
      border="border"
    >
      <el-table-column :label="$t('project.table.name')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('project.table.desc')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.desc }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('project.table.cteTime')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.cteTime }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('project.table.uteTime')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.uteTime }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('project.table.operation')" align="center">
        <template slot-scope="scope">
          <el-button circle size="mini" icon="el-icon-position" @click="handleRun(scope.row.id)" />
          <el-button circle size="mini" icon="el-icon-edit" @click="handleDetail(scope.row.id)" />
          <el-button circle size="mini" icon="el-icon-coin" @click="handleDatabase(scope.row.id)" />
          <el-button circle size="mini" icon="el-icon-delete" @click="handleDelete(scope.row.id)" />
        </template>
      </el-table-column>
    </el-table>
    <pagination v-show="total>0" :total="total" :page.sync="listQuery.pageNum" :limit.sync="listQuery.pageSize" @pagination="getList" />

    <el-dialog
      :title="project.id ? $t('project.item.editTitle') : $t('project.item.addTitle')"
      :visible.sync="dialogFormVisible"
      @close="handleFormClose('projectForm')"
    >
      <el-form ref="projectForm" :inline="true" :model="project" :rules="projectRules()" label-width="120px" label-suffix=":">
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
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleFormClose('projectForm')">{{ $t('common.form.cancel') }}</el-button>
        <el-button type="primary" @click="handleFormSubmit('projectForm')">{{ $t('common.form.confirm') }}</el-button>
      </span>
    </el-dialog>

    <el-dialog
      :title="database.id ? $t('project.database.editTitle') : $t('project.database.addTitle')"
      :visible.sync="dialogForm2Visible"
      @close="handleForm2Close('databaseForm')"
    >
      <el-form ref="databaseForm" :model="database" :rules="databaseRules()" label-width="120px" label-suffix=":">
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
            <el-form-item :label="$t('project.database.delimitKeyword')" prop="delimitKeyword">
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
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleForm2Close('databaseForm')">{{ $t('common.form.cancel') }}</el-button>
        <el-button type="primary" @click="handleForm2Submit('databaseForm')">{{ $t('common.form.confirm') }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { projectList, projectSelect, projectInsert, projectUpdate, projectDelete } from '@/api/project'
import { projectDbSelect, projectDbInsert, projectDbUpdate } from '@/api/project-db'
import { allDbType } from '@/api/type-column'
import { typeLanguageList } from '@/api/type-language'
import Pagination from '@/components/Pagination/index'
export default {
  name: 'Project',
  components: {
    Pagination
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
        name: '',
        desc: ''
      },
      dbTypeList: [],
      languageList: [],
      dialogFormVisible: false,
      project: {
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
      },
      dialogForm2Visible: false,
      database: {
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
      }
    }
  },
  created() {
    this.getList()
    this.getAllDbType()
    this.getLanguageList()
  },
  methods: {
    projectRules() {
      return {
        name: [{ required: true, message: this.$t('project.itemRules.name'), trigger: 'blur' }],
        desc: [{ required: true, message: this.$t('project.itemRules.desc'), trigger: 'blur' }],
        author: [{ required: true, message: this.$t('project.itemRules.author'), trigger: 'blur' }],
        organization: [{ required: true, message: this.$t('project.itemRules.organization'), trigger: 'blur' }],
        dateFormat: [{ required: true, message: this.$t('project.itemRules.dateFormat'), trigger: 'blur' }],
        fileEncode: [{ required: true, message: this.$t('project.itemRules.fileEncode'), trigger: 'blur' }],
        httpPrefix: [{ required: true, message: this.$t('project.itemRules.httpPrefix'), trigger: 'blur' }],
        generateRemark: [{ required: true, message: this.$t('project.itemRules.generateRemark'), trigger: 'blur' }]
      }
    },
    databaseRules() {
      return {
        type: [{ required: true, message: this.$t('project.databaseRules.type'), trigger: 'blur' }],
        url: [{ required: true, message: this.$t('project.databaseRules.url'), trigger: 'blur' }],
        database: [{ required: true, message: this.$t('project.databaseRules.database'), trigger: 'blur' }],
        host: [{ required: true, message: this.$t('project.databaseRules.host'), trigger: 'blur' }],
        port: [{ required: true, message: this.$t('project.databaseRules.port'), trigger: 'blur' }],
        username: [{ required: true, message: this.$t('project.databaseRules.username'), trigger: 'blur' }],
        password: [{ required: true, message: this.$t('project.databaseRules.password'), trigger: 'blur' }],
        delimitKeyword: [{ required: true, message: this.$t('project.databaseRules.delimitKeyword'), trigger: 'blur' }],
        languageId: [{ required: true, message: this.$t('project.databaseRules.languageType'), trigger: 'blur' }]
      }
    },
    getList() {
      this.listLoading = true
      projectList(this.listQuery).then(res => {
        this.list = res.data.list
        this.total = res.data.total
        // Just to simulate the time of the request
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    getAllDbType() {
      allDbType().then(res => {
        this.dbTypeList = res.data
      })
    },
    getLanguageList() {
      typeLanguageList({}).then(res => {
        this.languageList = res.data
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
    handleRun(id) {
      this.$router.push(`/project/detail/${id}`)
    },
    handleDetail(id) {
      this.dialogFormVisible = true
      this.project = {
        fileEncode: 'UTF-8',
        generateRemark: 1
      }
      if (id) {
        projectSelect(id).then(res => {
          this.project = res.data
        })
      }
    },
    handleDatabase(id) {
      this.dialogForm2Visible = true
      this.database = { projectId: id }
      projectDbSelect(id).then(res => {
        this.database = res.data
        this.database.projectId = id
        if (this.database.type === 0) {
          this.database.type = null
        }
        if (this.database.languageId === 0) {
          this.database.languageId = null
        }
      })
    },
    handleFormClose(formName) {
      this.dialogFormVisible = false
      this.$refs[formName].resetFields()
    },
    handleFormSubmit(formName) {
      this.$refs[formName].validate(validate => {
        if (validate) {
          if (this.project.id) {
            projectUpdate(this.project).then(res => {
              this.dialogFormVisible = false
              this.getList()
            })
          } else {
            projectInsert(this.project).then(res => {
              this.dialogFormVisible = false
              this.getList()
            })
          }
        }
      })
    },
    handleForm2Close(formName) {
      this.dialogForm2Visible = false
      this.$refs[formName].resetFields()
    },
    handleForm2Submit(formName) {
      this.$refs[formName].validate(validate => {
        if (validate) {
          this.handleDbTypeUrl(this.database.type)
          if (this.database.id) {
            projectDbUpdate(this.database).then(res => {
              this.dialogForm2Visible = false
              this.getList()
            })
          } else {
            projectDbInsert(this.database).then(res => {
              this.dialogForm2Visible = false
              this.getList()
            })
          }
        }
      })
    },
    handleDbTypeUrl(type) {
      const database = this.database
      if (type === 1) {
        database.url = database.username + ':' + database.password + '@tcp(' + database.host + ':' + database.port + ')/' + database.database + '?charset=utf8'
      } else if (type === 2) {
        database.url = database.username + '/' + database.password + '@' + database.host + ':' + database.port + '/' + database.database + '?charset=utf8'
      } else if (type === 3) {
        database.url = 'user=' + database.username + ' password=' + database.password + ' host=' + database.host + ' port=' + database.port + ' dbname=' + database.database + ' sslmode=disable'
      }
      this.database = database
    },
    handleDelete(id) {
      this.$confirm(this.$t('common.confirm.deleteOne'), this.$t('common.confirm.title'), {
        cancelButtonText: this.$t('common.confirm.cancel'),
        confirmButtonText: this.$t('common.confirm.confirm'),
        type: 'warning'
      }).then(() => {
        projectDelete(id).then(() => {
          this.getList()
        })
      })
    }
  }
}
</script>

<style scoped>

</style>
