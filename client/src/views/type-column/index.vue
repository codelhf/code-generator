<template>
  <div class="page-container">
    <el-form :model="listQuery" :inline="true" label-width="100px" label-suffix=":">
      <el-row>
        <el-col :span="6">
          <el-form-item :label="$t('typeColumn.listQuery.dbType')">
            <el-select v-model="listQuery.dbType" :placeholder="$t('typeColumn.listQuery.placeholderDbType')" style="width: 200px">
              <el-option v-for="(item) in dbTypeList" :key="item.id" :label="item.name" :value="item.id" />
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
            <el-select v-model="listQuery.languageId" :placeholder="$t('typeColumn.listQuery.placeholderLanguageType')" style="width: 200px">
              <el-option v-for="(item) in languageList" :key="item.id" :label="item.name" :value="item.id" />
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
          <span>{{ scope.row.dbName }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('typeColumn.table.columnType')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.columnType }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('typeColumn.table.fieldType')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.fieldType }}</span>
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
            <el-option v-for="(item) in dbTypeList" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('typeColumn.item.columnType')" prop="columnType">
          <el-input v-model="typeColumn.columnType" :placeholder="$t('typeColumn.item.placeholderColumnType')" />
        </el-form-item>
        <el-form-item :label="$t('typeColumn.item.languageType')" prop="fieldType">
          <el-select v-model="typeColumn.languageId" :placeholder="$t('typeColumn.item.placeholderLanguageType')">
            <el-option v-for="(item) in languageList" :key="item.id" :label="item.name" :value="item.id" />
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
import { typeColumnList, typeColumnSelect, typeColumnInsert, typeColumnUpdate, typeColumnDelete } from '@/api/type-column'
import { allDbType } from '@/api/type-column'
import { typeLanguageList } from '@/api/type-language'
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
        languageId: null,
        fieldType: ''
      },
      dbTypeList: [],
      languageList: [],
      dialogFormVisible: false,
      typeColumn: {
        id: '',
        languageId: null,
        dbType: null,
        columnType: '',
        fieldType: ''
      }
    }
  },
  created() {
    this.getList()
    this.getAllDbType()
    this.getLanguageList()
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
        dbType: null,
        columnType: '',
        languageId: null,
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
          if (this.typeColumn.languageId === 0) {
            this.typeColumn.languageId = null
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
