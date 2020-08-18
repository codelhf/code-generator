<template>
  <el-upload
    class="excel-uploader"
    action="/"
    :auto-upload="true"
    :show-file-list="true"
    :file-list="handleFileList"
    :before-upload="beforeAvatarUpload"
    :drag="true"
    :limit="1"
    :on-remove="handleRemove"
  >
    <i class="el-icon-upload" />
    <div class="el-upload__text">{{ $t('components.uploadExcel.drag') }} <em>{{ $t('components.uploadExcel.browse') }}</em></div>
    <div slot="tip" class="el-upload__tip">{{ $t('components.uploadExcel.accept') }}</div>
  </el-upload>
</template>

<script>
export default {
  name: 'UploadExcel',
  props: {
    fileList: {
      type: Array,
      required: false
    }
  },
  data() {
    return {
      handleFileList: []
    }
  },
  watch: {
    fileList: function(newVal, oldVla) {
      this.handleFileList = newVal
    }
  },
  methods: {
    beforeAvatarUpload(file) {
      if (!this.isExcel(file)) {
        this.$message.error(this.$t('components.uploadExcel.accept'))
        return false
      }
      const isLt10M = file.size / 1024 / 1024 < 10
      if (!isLt10M) {
        this.$message.error(this.$t('components.uploadExcel.error'))
      }
      console.log(file)
      this.handleFileList = []
      this.handleFileList.push({ name: file.name })
      this.$emit('onSuccess', file)
      return isLt10M
    },
    handleRemove(file, fileList) {
      this.handleFileList = []
      this.$emit('onSuccess', file)
    },
    isExcel(file) {
      return /\.(xlsx|xls)$/.test(file.name)
    }
  }
}
</script>

<style>
  .excel-uploader .el-upload,
  .excel-uploader .el-upload-dragger{
    width: 100%;
  }
</style>
