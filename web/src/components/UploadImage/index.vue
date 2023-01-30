<template>
  <el-upload
    ref="uploader"
    class="avatar-uploader"
    :auto-upload="false"
    :show-file-list="false"
    :before-upload="beforeAvatarUpload"
    :on-change="onAvatarChange"
    :limit="1"
  >
    <el-image v-if="imageUrl" :src="imageUrl" class="avatar" alt="avatar" />
    <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
  </el-upload>
</template>

<script>
import { ElMessage } from 'element-plus'
import { uploadImage } from '@/pages/lottery/api/upload'

export default {
  name: 'UploadImage',
  props: {
    imageUrl: {
      type: String,
      default: ''
    }
  },
  setup(props, ctx) {
    function beforeAvatarUpload(uploadFile) {
      console.log(uploadFile)
      const isImage = /\.(jpg|png|gif)$/.test(uploadFile.name)
      if (!isImage) {
        ElMessage.error('只能上传jpg、png、gif格式图片')
        return false
      }
      const isLt10M = uploadFile.size / 1024 / 1024 < 10
      if (!isLt10M) {
        ElMessage.error('上传头像图片大小不能超过10MB!')
        return false
      }
      return true
    }
    function onAvatarChange(uploadFile, uploadFiles) {
      if (uploadFile.status === 'ready' && beforeAvatarUpload(uploadFile)) {
        const formData = new FormData()
        // uploadFile is JSON Object, not Binary Data
        formData.append('file', uploadFile.raw)
        uploadImage(formData).then(res => {
          ctx.emit('upload-path', res.data)
        })
      }
    }
    return {
      beforeAvatarUpload,
      onAvatarChange
    }
  }
}
</script>

<style lang="scss" scoped>
.avatar-uploader {
  ::v-deep {
    .el-upload {
      border: 1px dashed #d9d9d9;
      border-radius: 6px;
      cursor: pointer;
      position: relative;
      overflow: hidden;
    }
    .el-upload:hover {
      border-color: #409EFF;
    }
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
  }
}
</style>
