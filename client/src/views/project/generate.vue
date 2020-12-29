<template>
  <div class="page-container">
    <el-row>
      <el-col :span="4">
        <table-list ref="tableList" :data="tableList" @tableClick="handleTableClick" />
      </el-col>
      <el-col :span="20" style="padding-left: 20px;">
        <template-list ref="templateList" :key="projectId" :project-id="projectId" />
        <el-form ref="currentTable" v-model="currentTable" :inline="true" label-width="80px" style="margin-top: 20px;">
          <el-form-item :label="$t('projectGenerate.currentTable.domainName')">
            <el-input v-model="currentTable.domainName" :placeholder="$t('projectGenerate.currentTable.placeholderDomainName')" />
          </el-form-item>
          <el-form-item :label="$t('projectGenerate.currentTable.domainDesc')">
            <el-input v-model="currentTable.domainDesc" :placeholder="$t('projectGenerate.currentTable.placeholderDomainDesc')" />
          </el-form-item>
          <el-form-item label=" ">
            <el-button type="primary" icon="el-icon-video-play" :loading="loading" @click="handleGenerate">{{ $t('projectGenerate.button.generateCode') }}</el-button>
            <el-button type="primary" icon="el-icon-plus" @click="handleDetail">{{ $t('projectGenerate.button.addTemplate') }}</el-button>
          </el-form-item>
        </el-form>
        <div style="width: 100%;">
          <el-input v-model="generateResult" type="textarea" :rows="10" />
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { projectLastId } from '../../api/project'
import { projectTableList, projectTableSelect, projectTableGenerate } from '../../api/project-table'
import TableList from './components/table'
import TemplateList from './components/template'
export default {
  name: 'Home',
  components: {
    TableList,
    TemplateList
  },
  data() {
    return {
      projectId: '',
      tableList: [],
      currentTable: {
        id: '',
        projectId: '',
        tableName: '',
        tableType: null,
        domainName: '',
        domainDesc: ''
      },
      loading: false,
      generateResult: ''
    }
  },
  created() {
    if (this.$route.params.id) {
      this.projectId = this.$route.params.id
      this.getTableList()
    } else {
      projectLastId().then(res => {
        this.projectId = res.data
        this.getTableList()
      })
    }
  },
  methods: {
    getTableList() {
      const params = { projectId: this.projectId }
      projectTableList(params).then(res => {
        this.tableList = res.data
      })
    },
    handleTableClick(table) {
      projectTableSelect(this.projectId, table.tableName).then(res => {
        if (res.data.tableName) {
          this.currentTable = res.data
        } else {
          this.currentTable = table
        }
        this.tableList = this.tableList.map(item => {
          item.active = false
          if (item.tableName === this.currentTable.tableName) {
            item.id = this.currentTable.id
            item.domainName = this.currentTable.domainName
            item.domainDesc = this.currentTable.domainDesc
            item.active = this.currentTable.active = true
          }
          return item
        })
      })
    },
    handleGenerate() {
      if (!this.currentTable.tableName) {
        this.$message.warning(this.$t('projectGenerate.button.tableWarning'))
        return
      }
      const templateList = this.$refs.templateList.list
      if (templateList.length === 0) {
        this.$message.warning(this.$t('projectGenerate.button.templateWarning'))
        return
      }
      this.loading = true
      this.generateResult = ''
      projectTableGenerate(this.projectId, [this.currentTable]).then(res => {
        setTimeout(() => {
          this.loading = false
          this.generateResult = res.data
        }, 1.5 * 1000)
      })
    },
    handleDetail() {
      this.$refs.templateList.handleDetail()
    }
  }
}
</script>

<style scoped>

</style>
