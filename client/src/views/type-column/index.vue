<template>
  <div class="page-container">
    <tree-list
      ref="treeList"
      :data="languageList"
      @itemClick="handleItemClick"
      @addItem="handleLanguage"
      @updateItem="handleLanguage"
      @deleteItem="handleLanguageDelete"
    >
      <el-form :model="listQuery" :inline="true" label-width="100px" label-suffix=":">
        <el-row>
          <el-form-item :label="$t('typeColumn.listQuery.dbType')">
            <el-select v-model="listQuery.dbType" :placeholder="$t('typeColumn.listQuery.placeholderDbType')" style="width: 200px">
              <el-option v-for="(item) in dbTypeList" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('typeColumn.listQuery.columnType')">
            <el-input v-model="listQuery.columnType" :placeholder="$t('typeColumn.listQuery.placeholderColumnType')" />
          </el-form-item>
          <el-form-item :label="$t('typeColumn.listQuery.fieldType')">
            <el-input v-model="listQuery.fieldType" :placeholder="$t('typeColumn.listQuery.placeholderFieldType')" />
          </el-form-item>
        </el-row>
        <el-row style="text-align: center">
          <el-form-item>
            <el-button type="primary" size="mini" icon="el-icon-search" @click="handleFilter">{{ $t('typeColumn.listButton.search') }}</el-button>
            <el-button type="primary" size="mini" icon="el-icon-refresh" @click="handleReset">{{ $t('typeColumn.listButton.reset') }}</el-button>
            <el-button type="primary" size="mini" icon="el-icon-plus" @click="handleDetail()">{{ $t('typeColumn.listButton.add') }}</el-button>
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
              <el-button circle size="mini" icon="el-icon-edit" @click="handleDetail(scope.row.id)" />
              <el-button circle size="mini" icon="el-icon-delete" @click="handleDelete(scope.row.id)" />
            </template>
          </el-table-column>
        </el-table>
      </div>

      <pagination v-show="total>0" :total="total" :page.sync="listQuery.pageNum" :limit.sync="listQuery.pageSize" @pagination="getList" />
    </tree-list>

    <el-dialog
      :title="language.id ? $t('typeColumn.language.editTitle') : $t('typeColumn.language.addTitle')"
      :visible.sync="dialogFormVisible"
      @close="handleFormClose('languageForm')"
    >
      <el-form ref="languageForm" :model="language" :rules="languageRules()" label-width="120px" label-suffix=":">
        <el-form-item :label="$t('typeColumn.language.name')" prop="name">
          <el-input v-model="language.name" :placeholder="$t('typeColumn.language.placeholderName')" @blur="checkLanguage" />
        </el-form-item>
        <el-form-item :label="$t('typeColumn.language.desc')" prop="desc">
          <el-input v-model="language.desc" :placeholder="$t('typeColumn.language.placeholderDesc')" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleFormClose('languageForm')">{{ $t('typeColumn.language.formCancel') }}</el-button>
        <el-button type="primary" @click="handleFormSubmit('languageForm')">{{ $t('typeColumn.language.formConfirm') }}</el-button>
      </span>
    </el-dialog>

    <el-dialog
      :title="typeColumn.id ? $t('typeColumn.item.editTitle') : $t('typeColumn.item.addTitle')"
      :visible.sync="dialogFormVisible2"
      @close="handleForm2Close('typeColumnForm')"
    >
      <el-form ref="typeColumnForm" :model="typeColumn" :rules="typeColumnRules()" label-width="120px" label-suffix=":">
        <el-form-item :label="$t('typeColumn.item.languageType')" prop="languageId">
          <el-select v-model="typeColumn.languageId" :placeholder="$t('typeColumn.item.placeholderLanguageType')">
            <el-option v-for="(item) in languageList" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('typeColumn.item.dbType')" prop="dbType">
          <el-select v-model="typeColumn.dbType" :placeholder="$t('typeColumn.item.placeholderDbType')">
            <el-option v-for="(item) in dbTypeList" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('typeColumn.item.columnType')" prop="columnType">
          <el-input v-model="typeColumn.columnType" :placeholder="$t('typeColumn.item.placeholderColumnType')" @blur="checkTypeColumn" />
        </el-form-item>
        <el-form-item :label="$t('typeColumn.item.fieldType')" prop="fieldType">
          <el-input v-model="typeColumn.fieldType" :placeholder="$t('typeColumn.item.placeholderFieldType')" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleForm2Close('typeColumnForm')">{{ $t('typeColumn.item.formCancel') }}</el-button>
        <el-button type="primary" @click="handleForm2Submit('typeColumnForm')">{{ $t('typeColumn.item.formConfirm') }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { typeColumnList, typeColumnSelect, typeColumnCheck, typeColumnInsert, typeColumnUpdate, typeColumnDelete } from '@/api/type-column'
import { typeLanguageList, typeLanguageSelect, typeLanguageCheck, typeLanguageInsert, typeLanguageUpdate, typeLanguageDelete } from '@/api/type-language'
import { allDbType } from '@/api/type-column'
import Pagination from '@/components/Pagination/index'
import TreeList from '@/components/TreeList/index'
export default {
  name: 'TypeColumn',
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
        languageId: '',
        dbType: null,
        columnType: '',
        fieldType: ''
      },
      dbTypeList: [],
      languageList: [],
      dialogFormVisible: false,
      language: {
        id: '',
        name: '',
        desc: ''
      },
      dialogFormVisible2: false,
      typeColumn: {
        id: '',
        languageId: '',
        dbType: null,
        columnType: '',
        fieldType: ''
      }
    }
  },
  created() {
    this.getList()
    this.getLanguageList()
    this.getAllDbType()
  },
  methods: {
    languageRules() {
      return {
        name: [{ required: true, message: this.$t('typeColumn.languageRules.name'), trigger: 'blur' }]
      }
    },
    typeColumnRules() {
      return {
        dbType: [{ required: true, message: this.$t('typeColumn.itemRules.dbType'), trigger: 'blur' }],
        columnType: [{ required: true, message: this.$t('typeColumn.itemRules.columnType'), trigger: 'blur' }],
        languageId: [{ required: true, message: this.$t('typeColumn.itemRules.languageType'), trigger: 'blur' }],
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
        languageId: '',
        dbType: null,
        columnType: '',
        fieldType: ''
      }
    },
    getLanguageList() {
      typeLanguageList({}).then(res => {
        this.languageList = res.data
        this.$refs.treeList.$data.clickedItem = ''
        this.languageList.map(item => {
          item.rightMenu = true
          if (item.id === this.listQuery.languageId) {
            item.active = true
          }
        })
      })
    },
    getAllDbType() {
      allDbType().then(res => {
        this.dbTypeList = res.data
      })
    },
    handleItemClick(item) {
      this.languageList = this.languageList.map(language => {
        language.active = false
        if (language.id === item.id) {
          language.active = true
          this.listQuery.languageId = item.id
          this.getList()
        }
        return language
      })
    },
    handleLanguage(item) {
      this.dialogFormVisible = true
      this.language = {}
      if (item && item.id) {
        typeLanguageSelect(item.id).then(res => {
          this.language = res.data
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
          if (this.language.id) {
            typeLanguageUpdate(this.language).then(res => {
              this.dialogFormVisible = false
              this.getList()
              this.getLanguageList()
            })
          } else {
            typeLanguageInsert(this.language).then(res => {
              this.dialogFormVisible = false
              this.getList()
              this.getLanguageList()
            })
          }
        }
      })
    },
    checkLanguage() {
      typeLanguageCheck(this.language).then(res => {
        if (res && res.msg) {
          this.$message({
            message: this.language.name + this.$t('system.message.exists'),
            type: 'error'
          })
        }
      })
    },
    handleLanguageDelete(item) {
      this.$confirm(this.$t('typeColumn.confirm.deleteOne'), this.$t('typeColumn.confirm.title'), {
        cancelButtonText: this.$t('typeColumn.confirm.cancel'),
        confirmButtonText: this.$t('typeColumn.confirm.confirm'),
        type: 'warning'
      }).then(() => {
        typeLanguageDelete(item.id).then(() => {
          this.getList()
          this.getLanguageList()
        })
      })
    },
    handleDetail(id) {
      this.dialogFormVisible2 = true
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
    handleForm2Close(formName) {
      this.dialogFormVisible2 = false
      this.$refs[formName].resetFields()
    },
    handleForm2Submit(formName) {
      this.$refs[formName].validate(validate => {
        if (validate) {
          if (this.typeColumn.id) {
            typeColumnUpdate(this.typeColumn).then(res => {
              this.dialogFormVisible2 = false
              this.getList()
            })
          } else {
            typeColumnInsert(this.typeColumn).then(res => {
              this.dialogFormVisible2 = false
              this.getList()
            })
          }
        }
      })
    },
    checkTypeColumn() {
      typeColumnCheck(this.typeColumn).then(res => {
        if (res && res.msg) {
          this.$message({
            message: this.typeColumn.columnType + this.$t('system.message.exists'),
            type: 'error'
          })
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
