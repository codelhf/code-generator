/**
 * @description 所有人可使用的参数配置列表
 * @params hideMenu: 是否隐藏当前路由结点不在导航中展示
 * @params alwaysShow: 只有一个子路由时是否总是展示菜单，默认false
 */
import { reactive } from 'vue'
import { createRouter, createWebHashHistory } from 'vue-router'
import NProgress from '@/utils/nprogress'
import store from '@/store'
import setting from '@/setting'
import { generateTitle } from '@/i18n'

NProgress.configure({ showSpinner: false })

// 引入不需要权限的modules
import System from './modules/system'

/**
 * @name 初始化必须要的路由
 * @description 使用reactive属性使得modules可以在路由菜单里面实时响应，搞定菜单回显的问题
 * @detail 针对modules的任何修改，均会同步至菜单级别，记住，是针对变量名为：moduels的修改
 **/
const modules = reactive([
  ...System
])

/** 引入需要权限的Modules */
// import Dashboard from './modules/pages/dashboard'

/** 登录后需要动态加入的本地路由 */
const asyncRoutes = [
  // ...Dashboard,
  ...setting.router
]

// 创建路由对象
const router = createRouter({
  history: createWebHashHistory(),
  routes: modules
})

// 未授权时可访问的白名单
const whiteList = ['/login', '/update_password', '/reset_password']

// 路由跳转前的监听操作
router.beforeEach((_to, _from, next) => {
  NProgress.start()
  if (store.state.user) {
    _to.meta.title ? (changeTitle(_to.meta.title)) : '' // 动态title
    if (_to.path === '/login') {
      next('/')
    } else {
      next()
    }
  } else if (whiteList.includes(_to.path)) {
    _to.meta.title ? (changeTitle(_to.meta.title)) : '' // 动态title
    next()
  } else {
    const redirect = window.location.href.replace(window.location.origin, '').replace('/#', '')
    next(`/login?redirect=${redirect}`) // 全部重定向到登录页
  }
})

// 路由跳转后的监听操作
router.afterEach((_to, _from) => {
  const keepAliveComponentsName = store.getters['keepAlive/keepAliveComponentsName'] || []
  const name = _to.matched[_to.matched.length - 1].components.default.name
  if (_to.meta && _to.meta.cache && name && !keepAliveComponentsName.includes(name)) {
    store.commit('keepAlive/addKeepAliveComponentsName', name)
  }
  NProgress.done()
})

// 修改页面标签标题
export function changeTitle(name) {
  name = generateTitle(name)
  document.title = `${name}-${setting.systemTitle}`
}

/**
 * @name 动态路由的权限新增，供登录后调用
 * @other 如果需要进行后端接口控制菜单的话，请在此拿到后端的菜单树与asyncRoutes对比，生成一个新的值
 */
export function addRoutes() {
  // 已验证完成，下面代码添加的可以实时同步至菜单中去，可以添加setTimeout(() => {}) 模拟异步代码的操作
  // 利用前端路由表模拟后端数据问题
  // const auths = store.state.user.info ? store.state.user.info.auths : []
  // const authPaths = auths.map(item => { item.targetUrl })
  // asyncRoutes.map(item => {
  //   if (authPaths.includes(item.path)) {
  //     const children = item.children
  //     if (children.length > 0) {
  //       for (let i = 0, iLength = children.length; i < iLength; i++) {
  //         const childrenItem = children[i]
  //         if (!authPaths.includes(item.path + '/' + childrenItem.path)) {
  //           children.splice(i, 1)
  //         }
  //       }
  //     }
  //     modules.push(item)
  //     router.addRoute(item)
  //   }
  // })
  asyncRoutes.map(item => {
    modules.push(item)
    router.addRoute(item)
  })
}

// 刷新页面时动态路由被清空，需重新添加路由
// if (store.state.user.token) {
//   addRoutes()
// }
addRoutes()

export default router
