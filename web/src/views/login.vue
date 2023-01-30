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
            <h3 class="tab" :class="loginType === 'dingtalk' ? 'active' : ''" @click="onLoginTypeClick('dingtalk')">钉钉扫码</h3>
            <el-divider direction="vertical" />
            <h3 class="tab" :class="loginType === 'password' ? 'active' : ''" @click="onLoginTypeClick('password')">密码登录</h3>
          </div>
          <div class="login-box-form">
            <el-form class="form dingtalk" v-if="loginType === 'dingtalk'">
              <div id="qr-code" class="qr-code"></div>
              <span class="title">打开钉钉扫一扫，快速登录/注册</span>
            </el-form>
            <el-form class="form password" v-if="loginType === 'password'" @keydown.enter="onLoginGetPhone">
              <el-form-item>
                <el-input
                  ref="username"
                  type="text"
                  v-model="form.username"
                  placeholder="用户名/手机号/邮箱"
                  size="large"
                  maxlength="50"
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
                <el-button type="primary" :loading="form.loading" @click="onLoginGetPhone" style="width: 100%;">登录</el-button>
                <el-col class="link">
                  <el-link class="link-text" :underline="false" @click="onUpdatePassword">修改密码</el-link>
                  <!--<el-link class="link-text" :underline="false" @click="onResetPassword">忘记密码</el-link>-->
                </el-col>
              </el-form-item>
            </el-form>
          </div>
          <div class="login-box-bottom">
            <!--<el-divider direction="horizontal">其他登录方式</el-divider>-->
            <!--<div class="third-icon">-->
            <!--  <a class="icon weixin" href="/#/auth/wechat">-->
            <!--    <img src="@/assets/login/weixin.png" alt="" />-->
            <!--  </a>-->
            <!--  <a class="icon qq" href="/#/auth/qq">-->
            <!--    <img src="@/assets/login/qq.png" alt="" />-->
            <!--  </a>-->
            <!--  <a class="icon weibo" href="/#/auth/weibo">-->
            <!--    <img src="@/assets/login/weibo.png" alt="" />-->
            <!--  </a>-->
            <!--</div>-->
          </div>
        </div>
      </div>
    </div>
    <el-dialog
      top="30vh"
      title="短信验证"
      width="30%"
      v-model="dialogFormVisible"
      @close="handleFormClose"
    >
      <el-form class="dialog-phone-form" @keydown.enter="onLoginByPasswordPhone">
        <el-form-item>
          <el-input
            ref="phone"
            type="text"
            v-model="form.showPhone"
            placeholder="手机号"
            size="large"
            maxlength="50"
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
          <el-button type="primary" :loading="form.loading" @click="onLoginByPasswordPhone" style="width: 100%;">确认</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import setting from '@/setting'
import { defineComponent, ref, reactive, nextTick, onMounted } from 'vue'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router'
import { addRoutes } from '@/router'
import { ElMessage } from 'element-plus'
import { isPhone, isValidSmsCode, isPassword } from '@/utils'
import { getPhoneApi } from '@/api/user'
import { smsCode } from '@/api/valid-code'
export default defineComponent({
  name: 'Login',
  setup() {
    const systemTitle = setting.systemTitle
    const systemSubTitle = setting.systemSubTitle
    const baseURL = setting.baseURL
    const store = useStore()
    const router = useRouter()
    const route = useRoute()
    onMounted(() => {
      dingLogin()
    })
    // 登录方式
    const loginType = ref('dingtalk')
    const onLoginTypeClick = (tabName) => {
      loginType.value = tabName
      if (tabName === 'dingtalk') {
        dingLogin()
      }
    }
    // 钉钉扫码
    const dingLogin = () => {
      nextTick(() => {
        // STEP3：在需要的时候，调用 window.DTFrameLogin 方法构造登录二维码，并处理登录成功或失败的回调。
        window.DTFrameLogin(
          {
            id: 'qr-code',
            width: 300,
            height: 300
          },
          {
            redirect_uri: encodeURIComponent(`${window.location.origin + baseURL}/user/v1.0/login_qrcode`),
            client_id: 'dingmag2ff0st32hsc8s',
            scope: 'openid',
            response_type: 'code',
            state: 'STATE',
            prompt: 'consent'
          },
          (loginResult) => {
            const { authCode } = loginResult
            // 这里可以直接进行重定向
            // window.location.href = redirectUrl
            // 也可以在不跳转页面的情况下，使用code进行授权
            console.log(loginResult)
            console.log(authCode)
            const params = {
              authCode: authCode
            }
            storeLogin('user/loginByQRCode', params)
          },
          (errorMsg) => {
            // 这里一般需要展示登录失败的具体原因
            alert(`Login Error: ${errorMsg}`)
          },
        )
      })
    }
    // 登录表单
    const form = reactive({
      loginType,
      loading: false,
      username: '',
      password: '',
      phone: '',
      validSmsCode: ''
    })

    /**
     * 密码登录
     */
    const passwordType = ref('password')
    const passwordTypeChange = () => {
      passwordType.value === '' ? passwordType.value = 'password' : passwordType.value = ''
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
    const dialogFormVisible = ref(false)
    const showCountDown = ref(false)
    const countDown = ref(300)
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
    // 关闭短信认证弹窗
    const handleFormClose = () => {
      dialogFormVisible.value = true
    }
    // 获取手机号
    const onLoginGetPhone = () => {
      if (!checkUsername() || !checkPassword()) {
        return
      }
      const params = {
        username: form.username,
        password: form.password
      }
      getPhoneApi(params).then(res => {
        form.phone = res.data
        form.showPhone = res.data.replace(/^(\d{3})\d{4}(\d{4})$/, '$1****$2')
        dialogFormVisible.value = true
      })
    }
    // 登录（双因子认证：密码加短信验证码）
    const onLoginByPasswordPhone = () => {
      if (!checkUsername() || !checkPassword() || !checkPhone() || !checkValidCode()) {
        return
      }
      const params = {
        username: form.username,
        password: form.password,
        phone: form.phone,
        validCode: form.validSmsCode
      }
      storeLogin('user/login', params)
    }
    // 调用store登录方法
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
      passwordType,
      passwordTypeChange,
      dialogFormVisible,
      showCountDown,
      countDown,
      handleFormClose,
      onValidSmsCodeClick,
      onLoginGetPhone,
      onLoginByPasswordPhone,
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
          .form.dingtalk {
            margin-top: 0;
            .qr-code {
              display: block;
              margin: 0 auto;
              width: 300px;
              height: 300px;
            }
            .title {
              display: block;
              margin-top: 20px;
              font-size: 12px;
              text-align: center;
            }
          }
          .form.password {
            margin-top: 50px;
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
