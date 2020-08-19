<template>
  <div class="page-container">
    <el-form :model="listQuery" :inline="true" label-width="120px" label-suffix=":">
      <el-row>
        <el-form-item :label="$t('templateField.listQuery.name')">
          <el-input v-model="listQuery.name" :placeholder="$t('templateField.listQuery.placeholderName')" />
        </el-form-item>
        <el-form-item :label="$t('templateField.listQuery.desc')">
          <el-input v-model="listQuery.desc" :placeholder="$t('templateField.listQuery.placeholderDesc')" />
        </el-form-item>
      </el-row>
      <el-row style="text-align: center">
        <el-form-item>
          <el-button type="primary" icon="el-icon-search" @click="handleFilter">{{ $t('templateField.listButton.search') }}</el-button>
          <el-button type="primary" icon="el-icon-refresh" @click="handleReset">{{ $t('templateField.listButton.reset') }}</el-button>
          <el-button type="primary" icon="el-icon-plus" @click="handleDetail()">{{ $t('templateField.listButton.add') }}</el-button>
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
      <el-table-column :label="$t('templateField.table.name')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('templateField.table.desc')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.desc }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('templateField.table.type')" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.type === 1">{{ $t('templateField.table.type1') }}</span>
          <span v-if="scope.row.type === 2">{{ $t('templateField.table.type2') }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('templateField.table.value')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.value }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('templateField.table.cteTime')" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.type === 2">{{ scope.row.cteTime }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('templateField.table.uteTime')" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.type === 2">{{ scope.row.uteTime }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('templateField.table.operation')" align="center" width="160">
        <template slot-scope="scope">
          <el-button v-if="scope.row.type === 2" type="primary" size="mini" @click="handleDetail(scope.row.id)">{{ $t('templateField.table.edit') }}</el-button>
          <el-button v-if="scope.row.type === 2" type="warning" size="mini" @click="handleDelete(scope.row.id)">{{ $t('templateField.table.delete') }}</el-button>
        </template>
      </el-table-column>
    </el-table>
    <pagination v-show="total>0" :total="total" :page.sync="listQuery.pageNum" :limit.sync="listQuery.pageSize" @pagination="getList" />

    <el-dialog
      :title="templateField.id ? $t('templateField.item.editTitle') : $t('templateField.item.addTitle')"
      :visible.sync="dialogFormVisible"
      @close="handleFormClose('templateFieldForm')"
    >
      <el-form ref="templateFieldForm" :model="templateField" :rules="templateFieldRules()" label-width="120px" label-suffix=":">
        <el-form-item :label="$t('templateField.item.name')" prop="name">
          <el-input v-model="templateField.name" :placeholder="$t('templateField.item.placeholderName')" />
        </el-form-item>
        <el-form-item :label="$t('templateField.item.desc')">
          <el-input v-model="templateField.desc" :placeholder="$t('templateField.item.placeholderDesc')" />
        </el-form-item>
        <el-form-item :label="$t('templateField.item.value')" prop="value">
          <el-input v-model="templateField.value" :placeholder="$t('templateField.item.placeholderValue')" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleFormClose('templateFieldForm')">{{ $t('templateField.item.formCancel') }}</el-button>
        <el-button type="primary" @click="handleFormSubmit('templateFieldForm')">{{ $t('templateField.item.formConfirm') }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { templateFieldList, templateFieldSelect, templateFieldInsert, templateFieldUpdate, templateFieldDelete } from '../../api/template-field'
import Pagination from '@/components/Pagination/index'
export default {
  name: 'TemplateField',
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
      dialogFormVisible: false,
      templateField: {
        id: '',
        name: '',
        desc: '',
        type: null,
        value: '',
        cteTime: '',
        uteTime: ''
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    templateFieldRules() {
      return {
        name: [{ required: true, message: this.$t('templateField.itemRules.name'), trigger: 'blur' }],
        desc: [{ required: true, message: this.$t('templateField.itemRules.desc'), trigger: 'blur' }],
        value: [{ required: true, message: this.$t('templateField.itemRules.value'), trigger: 'blur' }]
      }
    },
    getList() {
      this.listLoading = true
      templateFieldList(this.listQuery).then(res => {
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
    handleDetail(id) {
      this.dialogFormVisible = true
      this.templateField = {}
      if (id) {
        templateFieldSelect(id).then(res => {
          this.templateField = res.data
        })
      }
    },
    handleFormClose(formName) {
      this.dialogFormVisible = false
      this.$refs[formName].resetFields()
    },
    handleFormSubmit(formName) {
      this.$refs[formName].validate(validate => {
        if (validate) {
          if (this.templateField.id) {
            templateFieldUpdate(this.templateField).then(res => {
              this.dialogFormVisible = false
              this.getList()
            })
          } else {
            templateFieldInsert(this.templateField).then(res => {
              this.dialogFormVisible = false
              this.getList()
            })
          }
        }
      })
    },
    handleDelete(id) {
      this.$confirm(this.$t('templateField.confirm.deleteOne'), this.$t('templateField.confirm.title'), {
        cancelButtonText: this.$t('templateField.confirm.cancel'),
        confirmButtonText: this.$t('templateField.confirm.confirm'),
        type: 'warning'
      }).then(() => {
        templateFieldDelete(id).then(() => {
          this.getList()
        })
      })
    }
  }
}
</script>

<style scoped>

</style>
