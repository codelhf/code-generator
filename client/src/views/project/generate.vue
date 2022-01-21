<template>
  <div class="page-container">
    <table-layout>
      <template v-slot:left>
        <tree-list
          ref="treeList"
          :data="tableList"
          :node-key="'tableName'"
          :show-icon="true"
          :disable-right-click="true"
          @node-click="handleTableClick"
        />
      </template>
      <template v-slot:right>
        <template-list
          ref="templateList"
          :key="projectId"
          :project-id="projectId"
        />

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

        <div style="width: 100%; height: calc(((100vh - 124px) / 2) - 84px);">
          <el-input v-model="generateResult" type="textarea" :rows="10" />
        </div>
      </template>
    </table-layout>
  </div>
</template>

<script>
import { projectLastId } from '@/api/project'
import { projectTableList, projectTableSelect, projectTableGenerate } from '@/api/project-table'
import TreeList from '@/components/TreeList/index'
import TableLayout from '@/components/TreeList/layout'
import TemplateList from './components/template'
export default {
  name: 'Home',
  components: {
    TreeList,
    TableLayout,
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
        if (!res.data) {
          this.$router.push('/project/index')
        }
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
        this.tableList.map(item => {
          item.name = item.tableName
          item.table = item.tableType === 1
        })
      })
    },
    handleTableClick(table) {
      projectTableSelect(this.projectId, table.tableName).then(res => {
        if (res.data.tableName) {
          this.currentTable = res.data
        } else {
          this.currentTable = table
        }
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
