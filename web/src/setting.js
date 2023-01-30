/**
 * 应用配置
 */
import lottery from '@/pages/code-generator/setting'

// 打包应用配置，标题默认第一个
const builds = [lottery]

const systemTitle = builds[0].systemTitle
const systemSubTitle = builds[0].systemSubTitle
const baseURL = builds[0].baseURL

const router = []
const store = {}
const i18n = {}

builds.map(item => {
  router.push(...item.router)
  Object.assign(store, item.store)
  Object.assign(i18n, item.i18n)
})

export default {
  systemTitle,
  systemSubTitle,
  baseURL,
  router,
  store,
  i18n
}
