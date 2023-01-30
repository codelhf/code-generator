import axios from 'axios'
import { ElMessage } from 'element-plus'
import store from '@/store'
import setting from '@/setting'
// const baseURL = import.meta.env.VITE_BASE_URL

const service = axios.create({
  baseURL: setting.baseURL,
  timeout: 50000
})

// 请求前的统一处理
service.interceptors.request.use(
  (config) => {
    // 请求头
    if (!config.headers['Content-Type']) {
      config.headers['Content-Type'] = 'application/json'
    }
    // JWT鉴权处理
    if (store.getters['user/token']) {
      config.headers['X-Auth-Token'] = store.state.user.token
    }
    return config
  },
  (error) => {
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

service.interceptors.response.use(
  (response) => {
    const res = response.data
    if (res.code === 200) {
      return res
    } else {
      showError(res)
      return Promise.reject(res)
    }
  },
  (error) => {
    console.log(error) // for debug
    const badMessage = error.message || error
    const code = parseInt(badMessage.toString().replace('Error: Request failed with status code ', ''))
    showError({ code, message: badMessage })
    return Promise.reject(error)
  }
)

function showError(error) {
  if (error.code === 401 || error.code === 2001) {
    // to re-login
    store.dispatch('user/logout')
  } else {
    ElMessage({
      message: error.msg || error.message || '服务异常',
      type: 'error',
      duration: 3 * 1000
    })
  }
}

export default service
