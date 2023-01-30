/**
 * 代码生成器
 */
const systemTitle = 'Code Generator' // 系统名称，用于显示在左上角模块，登录模块、以及浏览器标题上使用,使用配置项
const systemSubTitle = '时间不在于你拥有多少,而在于你怎样使用。' // 系统提示信息，用于登录模块使用
const baseURL = import.meta.env.MODE === 'development' ? '/dev-api' : ''

import router from './router'
import store from './store'
import i18n from './i18n'

export default {
  systemTitle,
  systemSubTitle,
  baseURL,
  router,
  store,
  i18n
}
