<template>
  <div class="page-container">
    <tree-list
      ref="treeList"
      :data="typeList"
      :show-search="false"
      @itemClick="handleItemClick"
    >
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
            <el-button type="primary" size="mini" icon="el-icon-search" @click="handleFilter">{{ $t('templateField.listButton.search') }}</el-button>
            <el-button type="primary" size="mini" icon="el-icon-refresh" @click="handleReset">{{ $t('templateField.listButton.reset') }}</el-button>
            <el-button type="primary" size="mini" icon="el-icon-plus" @click="handleDetail()">{{ $t('templateField.listButton.add') }}</el-button>
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
          <el-table-column :label="$t('templateField.table.value')" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.value }}</span>
            </template>
          </el-table-column>
          <el-table-column v-if="listQuery.type === 2" :label="$t('templateField.table.operation')" align="center">
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
      :title="templateField.id ? $t('templateField.item.editTitle') : $t('templateField.item.addTitle')"
      :visible.sync="dialogFormVisible"
      @close="handleFormClose('templateFieldForm')"
    >
      <el-form ref="templateFieldForm" :model="templateField" :rules="templateFieldRules()" label-width="120px" label-suffix=":">
        <el-form-item :label="$t('templateField.item.name')" prop="name">
          <el-input v-model="templateField.name" :placeholder="$t('templateField.item.placeholderName')" @blur="checkField" />
        </el-form-item>
        <el-form-item :label="$t('templateField.item.desc')">
          <el-input v-model="templateField.desc" :placeholder="$t('templateField.item.placeholderDesc')" />
        </el-form-item>
        <el-form-item :label="$t('templateField.item.value')" prop="value">
          <el-input v-model="templateField.value" :placeholder="$t('templateField.item.placeholderValue')" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleFormClose('templateFieldForm')">{{ $t('common.form.cancel') }}</el-button>
        <el-button type="primary" @click="handleFormSubmit('templateFieldForm')">{{ $t('common.form.confirm') }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { templateFieldList, templateFieldSelect, templateFieldInsert, templateFieldUpdate, templateFieldDelete, templateFieldCheck } from '@/api/template-field'
import Pagination from '@/components/Pagination/index'
import TreeList from '@/components/TreeList/index'
export default {
  name: 'TemplateField',
  components: {
    Pagination,
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
        name: '',
        desc: '',
        type: null
      },
      typeList: [],
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
    this.getTypeList()
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
        desc: '',
        type: null
      }
    },
    getTypeList() {
      setTimeout(() => {
        this.typeList = [
          { id: 1, name: this.$t('templateField.table.type1') },
          { id: 2, name: this.$t('templateField.table.type2') }
        ]
      }, 100)
    },
    handleItemClick(item) {
      this.typeList = this.typeList.map(type => {
        type.active = false
        if (type.name === item.name) {
          type.active = true
          this.listQuery.type = type.id
          this.getList()
        }
        return type
      })
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
    checkField() {
      templateFieldCheck(this.templateField).then(res => {
        // do nothing
      }, () => {
        this.$message.error(this.templateField.name + ' ' + this.$t('common.message.exists'))
      })
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
