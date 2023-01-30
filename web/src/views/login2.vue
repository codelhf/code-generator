<template>
  <div class="passport-container">
    <!--nav-->
    <!--<div class="passport-nav">-->
    <!--  <img src="@/assets/logo.png" alt="" />-->
    <!--</div>-->
    <!--inner-->
    <div class="passport-container-inner">
      <!--notes-->
      <div class="passport-notes">
        <img src="@/assets/login/left.jpg" alt="" />
        <div class="notes">
          <div>{{ systemTitle }}</div>
          <div>{{ systemSubTitle }}</div>
        </div>
      </div>
      <!--main-->
      <div class="passport-main">
        <div class="welcome-tips">
          <span>终于等到你~</span>
          <img src="@/assets/login/wel_tips.png" alt="" />
        </div>
        <div class="login-box">
          <div class="login-box-tabs">
            <h3 class="tab" :class="loginType === 'dingding' ? 'active' : ''" @click="onLoginTypeClick('dingding')">钉钉登录</h3>
            <el-divider direction="vertical" />
            <h3 class="tab" :class="loginType === 'phone' ? 'active' : ''" @click="onLoginTypeClick('phone')">免密登录</h3>
            <el-divider direction="vertical" />
            <h3 class="tab" :class="loginType === 'password' ? 'active' : ''" @click="onLoginTypeClick('password')">密码登录</h3>
          </div>
          <div class="login-box-form">
            <el-form class="form dingding" v-if="loginType === 'dingding'">
              <img class="qr-code" src="" alt="" />
              <span class="title">打开钉钉扫一扫，快速登录/注册</span>
            </el-form>
            <el-form class="form phone" v-if="loginType === 'phone'" @keydown.enter="onLoginByPhone">
              <el-form-item>
                <el-input
                  ref="phone"
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
                <el-button type="primary" :loading="form.loading" @click="onLoginByPhone" style="width: 100%;">登录</el-button>
              </el-form-item>
            </el-form>
            <el-form class="form password" v-if="loginType === 'password'" @keydown.enter="onLoginByPassword">
              <el-form-item>
                <el-input
                  ref="username"
                  type="text"
                  v-model="form.username"
                  placeholder="用户名/手机号/邮箱"
                  size="large"
                  maxlength="50"
                  @blur="onCheckUsernameBlur"
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
                  @blur="onCheckPasswordBlur"
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
                <el-input
                  ref="valid2Code"
                  v-model="form.validImgCode"
                  placeholder="验证码"
                  size="large"
                  maxlength="50"
                  class="valid-img-code"
                >
                  <template #prepend>
                    <component :is="useRenderIcon('valid-code', 'svg')" />
                  </template>
                  <template #append>
                    <img class="valid-img" :src="'data:image/jpg;base64,' + form.validImgCodeUrl" alt="验证码" :disabled="form.loadingImgCode" @click="onValidImgCodeClick" />
                  </template>
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" :loading="form.loading" @click="onLoginByPassword" style="width: 100%;">登录</el-button>
                <el-col class="link">
                  <el-link class="link-text" :underline="false" @click="onUpdatePassword">修改密码</el-link>
                  <el-link class="link-text" :underline="false" @click="onResetPassword">忘记密码</el-link>
                </el-col>
              </el-form-item>
            </el-form>
          </div>
          <div class="login-box-bottom">
            <el-divider direction="horizontal">其他登录方式</el-divider>
            <div class="third-icon">
              <a class="icon weixin" href="/#/auth/wechat">
                <img src="@/assets/login/weixin.png" alt="" />
              </a>
              <a class="icon qq" href="/#/auth/qq">
                <img src="@/assets/login/qq.png" alt="" />
              </a>
              <a class="icon weibo" href="/#/auth/weibo">
                <img src="@/assets/login/weibo.png" alt="" />
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import setting from '@/setting'
import { defineComponent, ref, reactive } from 'vue'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router'
import { addRoutes } from '@/router'
import { ElMessage } from 'element-plus'
import { isPhone, isValidSmsCode, isPassword, isValidImgCode } from '@/utils'
import { getPhoneApi } from '@/api/user'
import { checkPhoneApi } from '@/api/user-phone'
import { imageCode, smsCode } from '@/api/valid-code'
export default defineComponent({
  name: 'Login',
  setup() {
    const systemTitle = setting.systemTitle
    const systemSubTitle = setting.systemSubTitle
    const store = useStore()
    const router = useRouter()
    const route = useRoute()
    // 登录方式
    const loginType = ref('dingding')
    const onLoginTypeClick = (tabName) => {
      loginType.value = tabName
      if (tabName === 'password') {
        onValidImgCodeClick()
      }
    }
    // 登录表单
    const form = reactive({
      loginType,
      loading: false,
      phone: '',
      validSmsCode: '',
      username: 'admin',
      password: 'Newmedia@bb106',
      validImgCode: '',
      loadingImgCode: false,
      validImgCodeUrl: ''
    })
    /**
     * 免密登录
     */
    const showCountDown = ref(false)
    const countDown = ref(300)
    const onCheckPhoneBlur = () => {
      checkPhoneApi({ phone: form.phone })
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
        return
      }
      smsCode(form).then(res => {
        // 发送短信请求成功后开始倒计时
        showCountDown.value = true
        if (timer) {
          clearInterval(timer)
        }
        timer = setInterval(() => {
          countDown.value = countDown.value - 1
          if (countDown.value === 0) {
            clearInterval(timer)
            showCountDown.value = false
            countDown.value = 300
          }
        }, 1000)
      })
    }
    const onLoginByPhone = () => {
      if (!checkPhone() || !checkValidCode()) {
        return
      }
      // 请求登录接口
      const params = {
        phone: form.phone,
        validCode: form.validSmsCode
      }
      storeLogin('user/loginByPhone', params)
    }

    /**
     * 密码登录
     */
    const passwordType = ref('password')
    const passwordTypeChange = () => {
      passwordType.value === '' ? passwordType.value = 'password' : passwordType.value = ''
    }
    const onCheckUsernameBlur = () => {
      checkUsernameApi({ username: form.username })
    }
    const onCheckPasswordBlur = () => {
      checkPasswordApi({ password: form.password })
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
    const checkValidImgCode = () => {
      if (!form.validImgCode) {
        ElMessage.warning({
          message: '请输入验证码',
          type: 'warning'
        })
        return false
      }
      if (!isValidImgCode(form.validImgCode)) {
        ElMessage.warning({
          message: '验证码格式不正确',
          type: 'warning'
        })
        return false
      }
      return true
    }
    const onValidImgCodeClick = () => {
      if (!form.loadingImgCode) {
        form.loadingImgCode = true
        imageCode().then(res => {
          form.validImgCodeUrl = res.data
          // 增加延时性
          setTimeout(() => {
            form.loadingImgCode = false
          }, 2000)
        })
      }
    }
    const onLoginByPassword = () => {
      if (!checkUsername() || !checkPassword() || !checkValidImgCode()) {
        return
      }
      const params = {
        username: form.username,
        password: form.password,
        validCode: form.validImgCode
      }
      storeLogin('user/login', params)
    }

    const storeLogin = (type, params) => {
      store.dispatch(type, params).then(() => {
        ElMessage.success({
          message: '登录成功',
          type: 'success',
          showClose: true,
          duration: 1000
        })
        addRoutes()
        router.push(route.query.redirect || '/')
      })
    }
    // 修改密码
    const onUpdatePassword = () => {
      router.push('/update_password')
    }
    // 重置密码
    const onResetPassword = () => {
      router.push('/reset_password')
    }

    return {
      systemTitle,
      systemSubTitle,
      loginType,
      onLoginTypeClick,
      form,
      // 免密登录
      showCountDown,
      countDown,
      onCheckPhoneBlur,
      onValidSmsCodeClick,
      onLoginByPhone,
      // 密码登录
      passwordType,
      passwordTypeChange,
      onCheckUsernameBlur,
      onCheckPasswordBlur,
      onValidImgCodeClick,
      onLoginByPassword,
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
    width: 1160px;
    height: 540px;
    background: white;
    border-radius: 8px;
    transform: translate(-50%, -50%);
    overflow: hidden;
    box-shadow: 0 6px 20px 5px rgba(152, 152, 152, 0.1),
    0 16px 24px 2px rgba(117, 117, 117, 0.14);

    .passport-notes {
      position: relative;
      float: left;
      width: 660px;
      height: 100%;
      img {
        width: 100%;
        height: 100%;
      }
      .notes {
        position: absolute;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        background-image: linear-gradient(rgba(0,204,222,0.8), rgba(51,132,224,0.8));
        text-align: center;
        color: #fff;
        font-size: 1.8rem;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        letter-spacing: 2px;

        div:nth-child(1) {
          font-size: 3.5rem;
          margin-bottom: 1em;
        }
      }
    }

    .passport-main {
      float: right;
      width: 500px;
      height: 100%;
      .welcome-tips {
        position: absolute;
        top: 0;
        right: 10px;
        height: 56px;
        span {
          position: relative;
          top: -12px;
          left: 0;
          height: 52px;
          line-height: 52px;
          font-size: 12px;
          color: #000000;
          text-shadow: 0 2px 2px #555666;
        }
        img {
          width: 56px;
          height: 56px;
        }
      }
      .login-box {
        position: relative;
        top: 56px;
        width: 100%;
        height: calc(100% - 56px);
        border-radius: 2px 2px;
        background: #fff;
        &-tabs {
          padding-top: 20px;
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
          height: 230px;
          .form {
            margin: auto;
            width: 80%;
            .el-input {
              margin-bottom: 20px;
            }
            //.password-icon {
            //  cursor: pointer;
            //  color: #409eff;
            //}
          }
          .form.dingding {
            .qr-code {
              display: block;
              margin: 0 auto;
              width: 180px;
              height: 180px;
            }
            .title {
              display: block;
              margin-top: 20px;
              font-size: 12px;
              text-align: center;
            }
          }
          .form.password {
            .el-input {
              margin-bottom: 5px;
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
        &-bottom {
          padding: 0 20px;
          width: calc(100% - 40px);
          height: 120px;
          .third-icon {
            width: 100%;
            height: 40px;
            .icon {
              display: inline-block;
              width: 33%;
              height: 100%;
              text-align: center;
              img {
                width: 40px;
                height: 40px;
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
