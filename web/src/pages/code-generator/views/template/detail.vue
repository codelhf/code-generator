<template>
  <div class="page-container">
    <el-form ref="templateForm" :model="template" :rules="templateRules" label-width="120px" label-suffix=":">
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
        <!--<el-input type="textarea" v-model="template.template" :autosize="{ minRows: 12, maxRows: 20 }" :placeholder="$t('template.item.placeholderTemplate')" />-->
        <code-editor :key="timer" :value="template.template" :placeholder="$t('template.item.placeholderTemplate')" @input="handleTemplate" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="handleFormSubmit">{{ $t('common.form.confirm') }}</el-button>
        <el-button @click="handleFormClose">{{ $t('common.form.cancel') }}</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import CodeEditor from './codeEditor.vue'
export default {
  name: 'Detail',
  components: {
    CodeEditor
  }
}
</script>

<script setup>
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { $t } from '@/i18n'
import { templateSelect, templateCheck, templateInsert, templateUpdate } from '@/pages/code-generator/api/template'
import { templateGroupList } from '@/pages/code-generator/api/template-group'

const route = useRoute()
const id = ref('')
const timer = ref(new Date().getTime())
// 生命周期
onMounted(() => {
  id.value = route.params.id
  timer.value = new Date().getTime()
  getDetail()
  getGroupList()
})

const template = ref({
  id: '',
  groupId: '',
  name: '',
  desc: '',
  type: null,
  fileType: '',
  template: '',
  cteTime: '',
  uteTime: ''
})
const templateRules = ref({
  groupId: [{ required: true, message: $t('template.itemRules.group'), trigger: 'blur' }],
  name: [{ required: true, message: $t('template.itemRules.name'), trigger: 'blur' }],
  desc: [{ required: true, message: $t('template.itemRules.desc'), trigger: 'blur' }],
  type: [{ required: true, message: $t('template.itemRules.type'), trigger: 'blur' }],
  fileType: [{ required: true, message: $t('template.itemRules.fileType'), trigger: 'blur' }],
  template: [{ required: true, message: $t('template.itemRules.template'), trigger: 'blur' }]
})

function getDetail() {
  template.value = {}
  if (id.value !== '0') {
    templateSelect(id.value).then(res => {
      template.value = res.data
      timer.value = new Date().getTime()
    })
  }
}

const groupList = ref([])
function getGroupList() {
  templateGroupList({}).then(res => {
    groupList.value = res.data
  })
}

function handleCheckName() {
  templateCheck(template.value).then(res => {
    // do nothing
  }, () => {
    ElMessage.error(template.value.name + ' ' + $t('common.message.exists'))
  })
}

function handleTemplate(val) {
  template.value.template = val
}

const templateForm = ref(null)
const router = useRouter()
function handleFormClose() {
  templateForm.value.resetFields()
  router.push('/template/index')
}

function handleFormSubmit() {
  templateForm.value.validate(validate => {
    if (validate) {
      if (template.value.id) {
        templateUpdate(template.value).then(res => {
          handleFormClose()
        })
      } else {
        templateInsert(this.template).then(res => {
          handleFormClose()
        })
      }
    }
  })
}
</script>

<style scoped>

</style>
