<template>
  <div class="page-container">
    <el-form :model="listQuery" :inline="true" label-width="100px" label-suffix=":">
      <el-row>
        <el-col :span="6">
          <el-form-item :label="$t('typeColumn.listQuery.dbType')">
            <el-select v-model="listQuery.dbType" :placeholder="$t('typeColumn.listQuery.placeholderDbType')" style="width: 200px">
              <el-option :label="$t('project.database.type1')" :value="1" />
              <el-option :label="$t('project.database.type2')" :value="2" />
              <el-option :label="$t('project.database.type3')" :value="3" />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item :label="$t('typeColumn.listQuery.columnType')">
            <el-input v-model="listQuery.columnType" :placeholder="$t('typeColumn.listQuery.placeholderColumnType')" />
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item :label="$t('typeColumn.listQuery.languageType')">
            <el-select v-model="listQuery.languageType" :placeholder="$t('typeColumn.listQuery.placeholderLanguageType')" style="width: 200px">
              <el-option :label="$t('project.database.languageType1')" :value="1" />
              <el-option :label="$t('project.database.languageType2')" :value="2" />
              <el-option :label="$t('project.database.languageType3')" :value="3" />
              <el-option :label="$t('project.database.languageType4')" :value="4" />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item :label="$t('typeColumn.listQuery.fieldType')">
            <el-input v-model="listQuery.fieldType" :placeholder="$t('typeColumn.listQuery.placeholderFieldType')" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row style="text-align: center">
        <el-form-item>
          <el-button type="primary" icon="el-icon-search" @click="handleFilter">{{ $t('typeColumn.listButton.search') }}</el-button>
          <el-button type="primary" icon="el-icon-refresh" @click="handleReset">{{ $t('typeColumn.listButton.reset') }}</el-button>
          <el-button type="primary" icon="el-icon-plus" @click="handleDetail()">{{ $t('typeColumn.listButton.add') }}</el-button>
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
      <el-table-column :label="$t('typeColumn.table.dbType')" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.dbType === 1">{{ $t('project.database.type1') }}</span>
          <span v-if="scope.row.dbType === 2">{{ $t('project.database.type2') }}</span>
          <span v-if="scope.row.dbType === 3">{{ $t('project.database.type3') }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('typeColumn.table.columnType')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.columnType }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('typeColumn.table.languageType')" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.languageType === 1">{{ $t('project.database.languageType1') }}</span>
          <span v-if="scope.row.languageType === 2">{{ $t('project.database.languageType2') }}</span>
          <span v-if="scope.row.languageType === 3">{{ $t('project.database.languageType3') }}</span>
          <span v-if="scope.row.languageType === 4">{{ $t('project.database.languageType4') }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('typeColumn.table.fieldType')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.fieldType }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('typeColumn.table.cteTime')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.cteTime }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('typeColumn.table.uteTime')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.uteTime }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('typeColumn.table.operation')" align="center" width="160">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handleDetail(scope.row.id)">{{ $t('typeColumn.table.edit') }}</el-button>
          <el-button type="danger" size="mini" @click="handleDelete(scope.row.id)">{{ $t('typeColumn.table.delete') }}</el-button>
        </template>
      </el-table-column>
    </el-table>
    <pagination v-show="total>0" :total="total" :page.sync="listQuery.pageNum" :limit.sync="listQuery.pageSize" @pagination="getList" />

    <el-dialog
      :title="typeColumn.id ? $t('typeColumn.item.editTitle') : $t('typeColumn.item.addTitle')"
      :visible.sync="dialogFormVisible"
      @close="handleFormClose('typeColumnForm')"
    >
      <el-form ref="typeColumnForm" :model="typeColumn" :rules="typeColumnRules()" label-width="120px" label-suffix=":">
        <el-form-item :label="$t('typeColumn.item.dbType')" prop="dbType">
          <el-select v-model="typeColumn.dbType" :placeholder="$t('typeColumn.item.placeholderDbType')">
            <el-option :label="$t('project.database.type1')" :value="1" />
            <el-option :label="$t('project.database.type2')" :value="2" />
            <el-option :label="$t('project.database.type3')" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('typeColumn.item.columnType')" prop="columnType">
          <el-input v-model="typeColumn.columnType" :placeholder="$t('typeColumn.item.placeholderColumnType')" />
        </el-form-item>
        <el-form-item :label="$t('typeColumn.item.languageType')" prop="fieldType">
          <el-select v-model="typeColumn.languageType" :placeholder="$t('typeColumn.item.placeholderLanguageType')">
            <el-option :label="$t('project.database.languageType1')" :value="1" />
            <el-option :label="$t('project.database.languageType2')" :value="2" />
            <el-option :label="$t('project.database.languageType3')" :value="3" />
            <el-option :label="$t('project.database.languageType4')" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('typeColumn.item.fieldType')" prop="fieldType">
          <el-input v-model="typeColumn.fieldType" :placeholder="$t('typeColumn.item.placeholderFieldType')" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleFormClose('typeColumnForm')">{{ $t('typeColumn.item.formCancel') }}</el-button>
        <el-button type="primary" @click="handleFormSubmit('typeColumnForm')">{{ $t('typeColumn.item.formConfirm') }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { typeColumnList, typeColumnSelect, typeColumnInsert, typeColumnUpdate, typeColumnDelete } from '../../api/type-column'
import Pagination from '@/components/Pagination/index'
export default {
  name: 'TypeColumn',
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
        dbType: null,
        columnType: '',
        languageType: null,
        fieldType: ''
      },
      dialogFormVisible: false,
      typeColumn: {
        id: '',
        dbType: null,
        columnType: '',
        languageType: null,
        fieldType: '',
        cteTime: '',
        uteTime: ''
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    typeColumnRules() {
      return {
        dbType: [{ required: true, message: this.$t('typeColumn.itemRules.dbType'), trigger: 'blur' }],
        columnType: [{ required: true, message: this.$t('typeColumn.itemRules.columnType'), trigger: 'blur' }],
        languageType: [{ required: true, message: this.$t('typeColumn.itemRules.languageType'), trigger: 'blur' }],
        fieldType: [{ required: true, message: this.$t('typeColumn.itemRules.fieldType'), trigger: 'blur' }]
      }
    },
    getList() {
      this.listLoading = true
      typeColumnList(this.listQuery).then(res => {
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
        dbType: null,
        columnType: '',
        languageType: null,
        fieldType: ''
      }
    },
    handleDetail(id) {
      this.dialogFormVisible = true
      this.typeColumn = {}
      if (id) {
        typeColumnSelect(id).then(res => {
          this.typeColumn = res.data
          if (this.typeColumn.dbType === 0) {
            this.typeColumn.dbType = null
          }
          if (this.typeColumn.languageType === 0) {
            this.typeColumn.languageType = null
          }
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
          if (this.typeColumn.id) {
            typeColumnUpdate(this.typeColumn).then(res => {
              this.dialogFormVisible = false
              this.getList()
            })
          } else {
            typeColumnInsert(this.typeColumn).then(res => {
              this.dialogFormVisible = false
              this.getList()
            })
          }
        }
      })
    },
    handleDelete(id) {
      this.$confirm(this.$t('typeColumn.confirm.deleteOne'), this.$t('typeColumn.confirm.title'), {
        cancelButtonText: this.$t('typeColumn.confirm.cancel'),
        confirmButtonText: this.$t('typeColumn.confirm.confirm'),
        type: 'warning'
      }).then(() => {
        typeColumnDelete(id).then(() => {
          this.getList()
        })
      })
    }
  }
}
</script>

<style scoped>

</style>
