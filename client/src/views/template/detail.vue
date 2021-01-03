<template>
  <div class="page-container">
    <el-form ref="templateForm" :model="template" :rules="templateRules()" label-width="120px" label-suffix=":">
      <el-row>
        <el-col :span="8">
          <el-form-item :label="$t('template.item.group')" prop="groupId">
            <el-select v-model="template.groupId" :placeholder="$t('template.item.placeholderGroup')" style="width: 100%;">
              <el-option v-for="(item) in groupList" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item :label="$t('template.item.type')" prop="type">
            <el-select v-model="template.type" :placeholder="$t('template.item.placeholderType')" style="width: 100%;">
              <el-option :label="$t('template.item.type1')" :value="1" />
              <el-option :label="$t('template.item.type2')" :value="2" />
              <el-option :label="$t('template.item.type3')" :value="3" />
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="8">
          <el-form-item :label="$t('template.item.name')" prop="name">
            <el-input v-model="template.name" :placeholder="$t('template.item.placeholderName')" @blur="handleCheckName" />
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item :label="$t('template.item.fileType')" prop="fileType">
            <el-input v-model="template.fileType" :placeholder="$t('template.item.placeholderFileType')" />
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item :label="$t('template.item.desc')">
            <el-input v-model="template.desc" :placeholder="$t('template.item.placeholderDesc')" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item :label="$t('template.item.template')" prop="template">
        <code-editor :key="timer" v-model="template.template" :value="template.template" :placeholder="$t('template.item.placeholderTemplate')" @input="handleTemplate" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="handleFormSubmit('templateForm')">{{ $t('common.form.confirm') }}</el-button>
        <el-button @click="handleFormClose('templateForm')">{{ $t('common.form.cancel') }}</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { templateSelect, templateCheck, templateInsert, templateUpdate } from '@/api/template'
import { templateGroupList } from '@/api/template-group'
import CodeEditor from './codeEditor'
export default {
  name: 'Detail',
  components: {
    CodeEditor
  },
  data() {
    return {
      template: {
        id: '',
        groupId: '',
        name: '',
        desc: '',
        type: null,
        fileType: '',
        template: '',
        cteTime: '',
        uteTime: ''
      },
      groupList: [],
      timer: ''
    }
  },
  created() {
    this.id = this.$route.params.id
    this.timer = new Date().getTime()
    this.getDetail()
    this.getGroupList()
  },
  methods: {
    templateRules() {
      return {
        groupId: [{ required: true, message: this.$t('template.itemRules.group'), trigger: 'blur' }],
        name: [{ required: true, message: this.$t('template.itemRules.name'), trigger: 'blur' }],
        desc: [{ required: true, message: this.$t('template.itemRules.desc'), trigger: 'blur' }],
        type: [{ required: true, message: this.$t('template.itemRules.type'), trigger: 'blur' }],
        fileType: [{ required: true, message: this.$t('template.itemRules.fileType'), trigger: 'blur' }],
        template: [{ required: true, message: this.$t('template.itemRules.template'), trigger: 'blur' }]
      }
    },
    getDetail() {
      this.template = {}
      if (this.id !== '0') {
        templateSelect(this.id).then(res => {
          this.template = res.data
          this.timer = new Date().getTime()
        })
      }
    },
    handleTemplate(val) {
      this.template.template = val
    },
    handleCheckName() {
      templateCheck(this.template.id, this.template.name).then(res => {
        // do nothing
      }, () => {
        this.$message.error(this.template.name + ' ' + this.$t('common.message.exists'))
      })
    },
    handleFormClose(formName) {
      this.$refs[formName].resetFields()
      this.$router.push('/template/index')
    },
    handleFormSubmit(formName) {
      this.$refs[formName].validate(validate => {
        if (validate) {
          if (this.template.id) {
            templateUpdate(this.template).then(res => {
              this.handleFormClose(formName)
            })
          } else {
            templateInsert(this.template).then(res => {
              this.handleFormClose(formName)
            })
          }
        }
      })
    },
    getGroupList() {
      templateGroupList({}).then(res => {
        this.groupList = res.data
      })
    }
  }
}
</script>

<style scoped>

</style>
