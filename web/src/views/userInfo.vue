<template>
  <div class="box">
    <el-row>
      <el-col :span="12">
        <el-form
          ref="userInfoForm"
          :model="userInfo"
          :rules="userInfoRules"
          label-suffix=":"
          label-width="120px"
          @close="handleFormClose"
        >
          <el-form-item label="用户头像" prop="avatar">
            <el-upload
              class="avatar-uploader"
              action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
              :show-file-list="false"
              :on-success="handleAvatarSuccess"
              :before-upload="beforeAvatarUpload"
            >
              <el-avatar :size="50" :src="userInfo.avatar" @click="onAvatarChange" />
            </el-upload>
          </el-form-item>
          <el-form-item label="用户名称" prop="username">
            <el-input v-model="userInfo.username" disabled placeholder="请输入用户名称" />
          </el-form-item>
          <el-form-item label="中文名称" prop="nickName">
            <el-input v-model="userInfo.nickName" placeholder="请输入中文名称" />
          </el-form-item>
          <el-form-item label="手机号" prop="phone">
            <el-input v-model="userInfo.phone" disabled placeholder="请输入手机号" />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="userInfo.email" disabled placeholder="请输入邮箱" />
          </el-form-item>
          <el-form-item label="简介" prop="remark">
            <el-input type="textarea" v-model="userInfo.remark" placeholder="请输入用户简介" />
          </el-form-item>
          <el-form-item>
            <span class="dialog-footer">
              <el-button @click="handleFormClose">取消</el-button>
              <el-button type="primary" @click="handleFormSubmit">确定</el-button>
            </span>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { defineComponent, ref, computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { updateUserInfo } from '@/api/user'
export default defineComponent({
  name: 'UserInfo',
  setup() {
    const store = useStore()
    const router = useRouter()
    const userInfoForm = ref(null)
    const userInfo = computed(() => store.state.user.info)
    const userInfoRules = ref({
      username: [{ required: true, message: '用户名称不能为空', trigger: 'blur' }],
      phone: [{ required: true, message: '手机号不能为空', trigger: 'blur' }],
      // email: [{ required: true, message: '邮箱不能为空', trigger: 'blur' }]
    })
    const onAvatarChange = (e) => {
      alert(e.target)
    }
    const handleFormClose = () => {
      userInfoForm.value.resetFields()
      router.back()
    }
    const handleFormSubmit = () => {
      userInfoForm.value.validate(valid => {
        if (valid) {
          updateUserInfo(userInfo).then(res => {
            router.back()
          })
        }
      })
    }
    return {
      userInfoForm,
      userInfo,
      userInfoRules,
      onAvatarChange,
      handleFormClose,
      handleFormSubmit
    }
  }
})
</script>

<style lang="scss" scoped>
  .box {
    padding: 15px;
  }
</style>
