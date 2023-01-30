<template>
  <div class="passport-container">
    <!--nav-->
    <!--<div class="passport-nav">-->
    <!--  <img src="@/assets/logo.png" alt="" />-->
    <!--</div>-->
    <!--inner-->
    <div class="passport-container-inner">
      <!--main-->
      <div class="passport-main">
        <div class="login-box">
          <div class="login-box-tabs">
            <h3 class="tab active">修改密码</h3>
          </div>
          <div class="login-box-form">
            <el-form class="form update" v-if="stepForm === 1" @keydown.enter="onUpdatePasswordCheck">
              <el-form-item>
                <el-input
                  ref="username"
                  type="text"
                  v-model="form.username"
                  placeholder="用户名/手机号/邮箱"
                  size="large"
                  maxlength="50"
                  @blur="checkUsername"
                >
                  <template #prepend>
                    <component :is="useRenderIcon('user', 'svg')" />
                  </template>
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-input
                  ref="password"
                  :type="passwordType"
                  v-model="form.password"
                  placeholder="密码"
                  size="large"
                  maxlength="50"
                  @blur="checkPassword"
                >
                  <template #prepend>
                    <component :is="useRenderIcon('lock', 'svg')" />
                  </template>
                  <template #append>
                    <component :is="useRenderIcon(passwordType ? 'eye-close' : 'eye-open', 'svg')" @click="passwordTypeChange" />
                  </template>
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" :loading="form.loading" @click="onUpdatePasswordCheck" style="width: 100%;">下一步</el-button>
                <el-col class="link">
                  <!--<el-link class="link-text" :underline="false" @click="onResetPassword">忘记密码</el-link>-->
                </el-col>
              </el-form-item>
            </el-form>
            <el-form class="form password" v-if="stepForm === 2" @keydown.enter="onUpdatePassword">
              <el-form-item>
                <el-input
                  ref="password"
                  :type="newPasswordType"
                  v-model="form.newPassword"
                  placeholder="请输入新密码"
                  size="large"
                  maxlength="50"
                  @blur="checkNewPassword"
                >
                  <template #prepend>
                    <component :is="useRenderIcon('lock', 'svg')" />
                  </template>
                  <template #append>
                    <component :is="useRenderIcon(newPasswordType ? 'eye-close' : 'eye-open', 'svg')" @click="newPasswordTypeChange" />
                  </template>
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-input
                  ref="password"
                  :type="confirmPasswordType"
                  v-model="form.confirmPassword"
                  placeholder="请再次输入新密码"
                  size="large"
                  maxlength="50"
                  @blur="checkConfirmPassword"
                >
                  <template #prepend>
                    <component :is="useRenderIcon('lock', 'svg')" />
                  </template>
                  <template #append>
                    <component :is="useRenderIcon(confirmPasswordType ? 'eye-close' : 'eye-open', 'svg')" @click="confirmPasswordTypeChange" />
                  </template>
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" :loading="form.loading" @click="onUpdatePassword" style="width: 100%;">提交</el-button>
                <el-col class="link">
                  <!--<el-link class="link-text" :underline="false" @click="onResetPassword">忘记密码</el-link>-->
                </el-col>
              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { isPassword } from '@/utils'
import { updatePasswordCheck, updatePassword } from '@/api/user'
export default defineComponent({
  name: 'UpdatePassword',
  setup() {
    const router = useRouter()
    const stepForm = ref(1)
    // 表单
    const form = reactive({
      loading: false,
      username: '',
      password: '',
      newPassword: '',
      confirmPassword: ''
    })
    // 是否显示密码
    const passwordType = ref('password')
    const passwordTypeChange = () => {
      passwordType.value === '' ? passwordType.value = 'password' : passwordType.value = ''
    }
    const newPasswordType = ref('password')
    const newPasswordTypeChange = () => {
      newPasswordType.value === '' ? newPasswordType.value = 'password' : newPasswordType.value = ''
    }
    const confirmPasswordType = ref('password')
    const confirmPasswordTypeChange = () => {
      confirmPasswordType.value === '' ? confirmPasswordType.value = 'password' : confirmPasswordType.value = ''
    }
    const checkUsername = () => {
      if (!form.username) {
        ElMessage.warning({
          message: '请输入用户名',
          type: 'warning'
        })
        return false
      }
      return true
    }
    const checkPassword = () => {
      if (!form.password) {
        ElMessage.warning({
          message: '请输入密码',
          type: 'warning'
        })
        return false
      }
      if (!isPassword(form.password)) {
        ElMessage.warning({
          message: '密码格式不正确',
          type: 'warning'
        })
        return false
      }
      return true
    }
    const checkNewPassword = () => {
      if (!form.newPassword) {
        ElMessage.warning({
          message: '请输入新密码',
          type: 'warning'
        })
        return false
      }
      if (!isPassword(form.newPassword)) {
        ElMessage.warning({
          message: '新密码格式不正确',
          type: 'warning'
        })
        return false
      }
      if (form.newPassword === form.password) {
        ElMessage.warning({
          message: '新密码和原密码不能相同',
          type: 'warning'
        })
        return false
      }
      return true
    }
    const checkConfirmPassword = () => {
      if (!form.confirmPassword) {
        ElMessage.warning({
          message: '请输入确认密码',
          type: 'warning'
        })
        return false
      }
      if (!isPassword(form.confirmPassword)) {
        ElMessage.warning({
          message: '确认密码格式不正确',
          type: 'warning'
        })
        return false
      }
      if (form.confirmPassword !== form.newPassword) {
        ElMessage.warning({
          message: '两次输入密码不一致',
          type: 'warning'
        })
        return false
      }
      return true
    }
    const onUpdatePasswordCheck = () => {
      if (!checkUsername() || !checkPassword()) {
        return
      }
      const params = {
        username: form.username,
        password: form.password,
        validCode: form.validImgCode
      }
      updatePasswordCheck(params).then(res => {
        stepForm.value = 2
      })
    }

    const onUpdatePassword = () => {
      if (!checkNewPassword() || !checkConfirmPassword()) {
        return
      }
      const params = {
        username: form.username,
        password: form.password,
        newPassword: form.newPassword
      }
      updatePassword(params).then(res => {
        // 修改密码成功后跳转到登录页
        router.push('/login')
      })
    }
    // 忘记密码
    const onResetPassword = () => {
      router.push('/reset_password')
    }

    return {
      stepForm,
      form,
      passwordType,
      passwordTypeChange,
      newPasswordType,
      newPasswordTypeChange,
      confirmPasswordType,
      confirmPasswordTypeChange,
      checkUsername,
      checkPassword,
      checkNewPassword,
      checkConfirmPassword,
      onUpdatePasswordCheck,
      onUpdatePassword,
      onResetPassword
    }
  }
})
</script>

<style lang="scss" scoped>
.passport-container {
  position: relative;
  width: 100vw;
  height: 100vh;
  background: #fff url('@/assets/login/bg.png') no-repeat center;
  overflow: hidden;
  background-size: cover;
  cursor: pointer;
  user-select: none;

  .passport-nav {
    position: relative;
    display: inline-block;
    padding: 0 20px;
    width: 100%;
    height: 60px;
    text-align: left;
    background: #343434;
    img {
      height: 60px;
      line-height: 60px;
    }
  }

  .passport-container-inner {
    display: flex;
    position: absolute;
    left: 50%;
    top: 50%;
    width: fit-content;
    height: 440px;
    background: white;
    border-radius: 8px;
    transform: translate(-50%, -50%);
    overflow: hidden;
    box-shadow: 0 6px 20px 5px rgba(152, 152, 152, 0.1),
    0 16px 24px 2px rgba(117, 117, 117, 0.14);

    .passport-main {
      float: right;
      width: 500px;
      height: 100%;
      .login-box {
        position: relative;
        top: 26px;
        width: 100%;
        height: calc(100% - 26px);
        border-radius: 2px 2px;
        background: #fff;
        &-tabs {
          width: 100%;
          height: 40px;
          text-align: center;
          .tab {
            display: inline-block;
            width: 30%;
            color: #ababab;
          }
          .tab.active {
            color: #222222;
          }
        }
        &-form {
          margin-top: 50px;
          width: 100%;
          height: 200px;
          .form {
            margin: auto;
            width: 80%;
            .el-input {
              margin-bottom: 20px;
            }
            .link {
              text-align: right;
              .link-text {
                justify-content: flex-end;
                margin-right: 10px;
              }
            }
          }
        }
      }
    }
  }
}

@media screen and (max-width: 1260px) {
  .passport-notes {
    display: none;
  }
  .passport-container-inner {
    width: 500px !important;
  }
}

@media screen and (max-width: 750px) {
  .passport-container .passport-container-inner, .passport-container .passport-main {
    width: 100vw !important;
    height: 100vh;
    box-shadow: none;
    left: 0;
    top: 0;
    transform: none;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
}
</style>

<style lang="scss">
.valid-img-code {
  .el-input-group__append {
    padding: 0;
  }
  .valid-img {
    width: 160px;
    height: 38px;
    padding: 0 2px;
  }
}
</style>
