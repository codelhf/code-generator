<template>
  <div class="page-container">
    <el-form ref="templateForm" :model="template" :rules="templateRules()" label-width="120px" label-suffix=":">
      <el-row>
        <el-col :span="8">
          <el-form-item :label="$t('template.item.name')" prop="name">
            <el-input v-model="template.name" :placeholder="$t('template.item.placeholderName')" @blur="handleCheckName" />
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item :label="$t('template.item.desc')">
            <el-input v-model="template.desc" :placeholder="$t('template.item.placeholderDesc')" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="8">
          <el-form-item :label="$t('template.item.type')" prop="type">
            <el-select v-model="template.type" :placeholder="$t('template.item.placeholderType')" style="width: 100%;">
              <el-option :label="$t('template.item.type1')" :value="1" />
              <el-option :label="$t('template.item.type2')" :value="2" />
              <el-option :label="$t('template.item.type3')" :value="3" />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item :label="$t('template.item.fileType')" prop="fileType">
            <el-input v-model="template.fileType" :placeholder="$t('template.item.placeholderFileType')" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item :label="$t('template.item.template')" prop="template">
        <code-editor :key="timer" v-model="template.template" :value="template.template" :placeholder="$t('template.item.placeholderTemplate')" @input="handleTemplate" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="handleFormSubmit('templateForm')">{{ $t('template.item.formConfirm') }}</el-button>
        <el-button @click="handleFormClose('templateForm')">{{ $t('template.item.formCancel') }}</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { templateSelect, templateExists, templateInsert, templateUpdate } from '../../api/template'
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
        name: '',
        desc: '',
        type: null,
        fileType: '',
        template: '',
        cteTime: '',
        uteTime: ''
      },
      timer: ''
    }
  },
  created() {
    this.id = this.$route.params.id
    this.getDetail()
    this.timer = new Date().getTime()
  },
  methods: {
    templateRules() {
      return {
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
      templateExists(this.template.id, this.template.name).then(res => {
      }, () => {
        this.$message.error(this.$t('template.itemRules.exists'))
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
    }
  }
}
</script>

<style scoped>

</style>
