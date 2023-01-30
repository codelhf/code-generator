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
            <h3 class="tab active">重置密码</h3>
          </div>
          <div class="login-box-form">
            <el-form class="form reset" v-if="stepForm === 1" @keydown.enter="onResetPasswordCheck">
              <el-form-item>
                <el-input
                  ref="username"
                  type="text"
                  v-model="form.phone"
                  placeholder="手机号"
                  size="large"
                  maxlength="50"
                  @blur="onCheckPhoneBlur"
                >
                  <template #prepend>
                    <component :is="useRenderIcon('phone', 'svg')" />
                  </template>
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-input
                  ref="validCode"
                  type="text"
                  v-model="form.validSmsCode"
                  placeholder="6位数字验证码"
                  size="large"
                  maxlength="50"
                >
                  <template #prepend>
                    <component :is="useRenderIcon('message', 'svg')" />
                  </template>
                  <template #append>
                    <span class="valid-code send" v-if="!showCountDown" @click="onValidSmsCodeClick">发送验证码</span>
                    <span class="valid-code count-down" v-if="showCountDown">{{ countDown }}s后重新获取</span>
                  </template>
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" :loading="form.loading" @click="onResetPasswordCheck" style="width: 100%;">下一步</el-button>
                <el-col class="link">
                  <el-link class="link-text" :underline="false" @click="onUpdatePassword">修改密码</el-link>
                </el-col>
              </el-form-item>
            </el-form>
            <el-form class="form password" v-if="stepForm === 2" @keydown.enter="onResetPassword">
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
                <el-button type="primary" :loading="form.loading" @click="onResetPassword" style="width: 100%;">提交</el-button>
                <el-col class="link">
                  <el-link class="link-text" :underline="false" @click="onUpdatePassword">修改密码</el-link>
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
import { isPhone, isPassword, isValidSmsCode } from '@/utils'
import { checkPhoneApi, resetPasswordCheck, resetPassword } from '@/api/user-phone'
import { smsCode } from '@/api/valid-code'
export default defineComponent({
  name: 'ResetPassword',
  setup() {
    const router = useRouter()
    const stepForm = ref(1)
    // 表单
    const form = reactive({
      loading: false,
      phone: '',
      validSmsCode: '',
      newPassword: '',
      confirmPassword: ''
    })
    const showCountDown = ref(false)
    const countDown = ref(300)
    const onCheckPhoneBlur = () => {
      checkPhoneApi({ phone: form.phone })
    }
    // 是否显示密码
    const newPasswordType = ref('password')
    const newPasswordTypeChange = () => {
      newPasswordType.value === '' ? newPasswordType.value = 'password' : newPasswordType.value = ''
    }
    const confirmPasswordType = ref('password')
    const confirmPasswordTypeChange = () => {
      confirmPasswordType.value === '' ? confirmPasswordType.value = 'password' : confirmPasswordType.value = ''
    }
    const checkPhone = () => {
      if (!form.phone) {
        ElMessage.warning({
          message: '请先输入手机号',
          type: 'warning'
        })
        return false
      }
      if (!isPhone(form.phone)) {
        ElMessage.warning({
          message: '手机号格式不正确',
          type: 'warning'
        })
        return false
      }
      return true
    }
    const checkValidCode = () => {
      if (!form.validSmsCode || !isValidSmsCode(form.validSmsCode)) {
        ElMessage.warning({
          message: '请输入6位数字验证码',
          type: 'warning'
        })
        return false
      }
      return true
    }
    let timer = null
    const onValidSmsCodeClick = () => {
      if (!checkPhone()) {
        console.log('1')
        return
      }
      console.log('2')
      smsCode(form).then(res => {
        // 发送短信请求成功后开始倒计时
        showCountDown.value = true
        if (timer) {
          clearInterval(timer)
        }
        timer = setInterval(() => {
          countDown.value = countDown.value - 1
          if (countDown.value === 0) {
            showCountDown.value = false
            clearInterval(timer)
          }
        }, 1000)
      })
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
    const onResetPasswordCheck = () => {
      if (!checkPhone() || !checkValidCode()) {
        return
      }
      const params = {
        phone: form.phone,
        validCode: form.validSmsCode
      }
      resetPasswordCheck(params).then(res => {
        stepForm.value = 2
      })
    }

    const onResetPassword = () => {
      if (!checkNewPassword() || !checkConfirmPassword()) {
        return
      }
      const params = {
        phone: form.phone,
        validCode: form.validSmsCode,
        newPassword: form.newPassword
      }
      resetPassword(params).then(res => {
        // 修改密码成功后跳转到登录页
        router.push('/login')
      })
    }
    // 修改密码
    const onUpdatePassword = () => {
      router.push('/update_password')
    }

    return {
      stepForm,
      form,
      onCheckPhoneBlur,
      showCountDown,
      countDown,
      onValidSmsCodeClick,
      newPasswordType,
      newPasswordTypeChange,
      confirmPasswordType,
      confirmPasswordTypeChange,
      checkNewPassword,
      checkConfirmPassword,
      onResetPasswordCheck,
      onResetPassword,
      onUpdatePassword
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
