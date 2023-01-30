<template>
  <el-upload
    ref="uploader"
    class="excel-uploader"
    :auto-upload="false"
    :show-file-list="true"
    :file-list="fileList"
    :before-upload="beforeUpload"
    :limit="1"
    :drag="true"
    :on-change="handleChange"
    :on-remove="handleRemove"
  >
    <el-icon class="el-icon--upload"><upload-filled /></el-icon>
    <div class="el-upload__text">
      将文件拖到此处，或 <em>点击上传</em>
    </div>
    <template #tip>
      <div class="el-upload__tip">
        上传文件大小不能超过10MB!
      </div>
    </template>
  </el-upload>
</template>

<script>
import { defineComponent, onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { readExcel } from '@/utils/components/excel'

export default defineComponent({
  name: 'UploadExcel',
  props: {
    header: {
      type: Array,
      required: true
    }
  },
  emits: ['onSuccess'],
  setup(props, ctx) {
    const fileList = ref([])
    onMounted(() => {
      fileList.value = []
    })

    function beforeUpload(uploadFile) {
      console.log(uploadFile)
      const isExcel = /\.(xlsx|xls)$/.test(uploadFile.name)
      if (!isExcel) {
        ElMessage.error('只能上传xlsx/xls文件')
        return false
      }
      const isLt10M = uploadFile.size / 1024 / 1024 < 10
      if (!isLt10M) {
        ElMessage.error('上传文件大小不能超过10MB!')
        return false
      }
      return true
    }

    function handleChange(uploadFile, uploadFiles) {
      if (window.FormData === undefined || window.XMLHttpRequest === undefined) {
        ElMessage.error('浏览器版本过低，暂不支持，请使用IE10及以上的版本或Chrome、Firefox等最新的浏览器')
      }
      if (uploadFile.status === 'ready' && beforeUpload(uploadFile)) {
        readExcel(uploadFile.raw, props.header, (data) => {
          console.log(data)
          ctx.emit('onSuccess', data)
          fileList.value = []
          fileList.value.push({ name: uploadFile.name })
        })
      }
    }

    function handleRemove(uploadFile, uploadFiles) {
      fileList.value = []
    }

    return {
      fileList,
      beforeUpload,
      handleChange,
      handleRemove
    }
  }
})
</script>

<style scoped>
  .excel-uploader >>> .el-upload,
  .excel-uploader >>> .el-upload-dragger{
    width: 100%;
  }
</style>
