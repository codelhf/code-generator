import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/404',
    component: () => import('@/views/error/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/home',
    children: [
      {
        path: 'home',
        name: 'Dashboard',
        component: () => import('@/views/project/generate'),
        meta: { title: 'dashboard', icon: 'home', noCache: false, affix: true }
      }
    ]
  },

  {
    path: '/project',
    component: Layout,
    redirect: '/project/index',
    children: [
      {
        path: 'index',
        name: 'Project',
        component: () => import('@/views/project/index'),
        meta: { title: 'project', icon: 'project', noCache: false }
      },
      {
        path: 'detail/:id',
        name: 'detail',
        component: () => import('@/views/project/generate'),
        meta: { title: 'projectGenerate', icon: '', noCache: false },
        hidden: true
      }
    ]
  },

  {
    path: '/template',
    component: Layout,
    redirect: '/template/index',
    children: [
      {
        path: 'index',
        name: 'Template',
        component: () => import('@/views/template/index'),
        meta: { title: 'template', icon: 'template', noCache: false }
      },
      {
        path: 'detail/:id',
        name: 'TemplateDetail',
        component: () => import('@/views/template/detail'),
        meta: { title: 'templateDetail', icon: '', noCache: false },
        hidden: true
      }
    ]
  },

  {
    path: '/templateField',
    component: Layout,
    redirect: '/templateField/index',
    children: [
      {
        path: 'index',
        name: 'TemplateField',
        component: () => import('@/views/template-field/index'),
        meta: { title: 'templateField', icon: 'templateField', noCache: false }
      }
    ]
  },

  {
    path: '/typeColumn',
    component: Layout,
    redirect: '/typeColumn/index',
    children: [
      {
        path: 'index',
        name: 'TypeColumn',
        component: () => import('@/views/type-column/index'),
        meta: { title: 'typeColumn', icon: 'typeColumn', noCache: false }
      }
    ]
  },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

export const asyncRoutes = []

const createRouter = () => new Router({
  mode: 'hash',
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

// 解决ElementUI导航栏中的vue-router在3.0版本以上重复点菜单报错问题
const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}
const originalReplace = Router.prototype.replace
Router.prototype.replace = function replace(location) {
  return originalReplace.call(this, location).catch(err => err)
}

export default router
