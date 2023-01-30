<template>
  <el-upload
    ref="uploader"
    class="file-uploader"
    :auto-upload="false"
    :show-file-list="true"
    :file-list="fileList"
    :multiple="multiple"
    :limit="10"
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
        <!--上传文件大小不能超过10MB!-->
      </div>
    </template>
  </el-upload>
</template>

<script>
import { defineComponent, onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'UploadFile',
  props: {
    multiple: {
      type: Boolean,
      default: false
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
      const isExcel = /\.(mp4|mp3)$/.test(uploadFile.name)
      if (!isExcel) {
        ElMessage.error('只能上传mp4/mp3文件')
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
      if (uploadFile.status === 'ready') {
        if (!props.multiple) {
          fileList.value = [uploadFile]
          ctx.emit('onSuccess', [uploadFile])
        } else {
          fileList.value = uploadFiles
          ctx.emit('onSuccess', uploadFiles)
        }
      }
    }

    function handleRemove(uploadFile, uploadFiles) {
      fileList.value = uploadFiles.map(item => {
        if (uploadFile.name !== item.name) {
          return item
        }
      })
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

<style lang="scss" scoped>
  .file-uploader {
    width: 100%;
    ::v-deep .el-upload {
      width: 100%;
    }
    ::v-deep .el-upload-dragger {
      padding: 0;
      //width: 100%;
    }
  }
</style>
