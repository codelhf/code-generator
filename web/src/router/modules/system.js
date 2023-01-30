import Layout from '@/layout/index.vue'
import { createNameComponent } from '@/router/createNode'
const route = [
  {
    path: '/system',
    component: Layout,
    redirect: '/404',
    meta: { title: '系统目录' },
    hidden: true,
    children: [
      {
        path: '/403',
        component: createNameComponent(() => import('@/views/error/403.vue')),
        meta: { title: '403', hideTabs: true }
      },
      {
        path: '/404',
        component: createNameComponent(() => import('@/views/error/404.vue')),
        meta: { title: '404', hideTabs: true }
      },
      {
        path: '/500',
        component: createNameComponent(() => import('@/views/error/500.vue')),
        meta: { title: '500', hideTabs: true }
      },
      {
        path: '/redirect/:path(.*)',
        component: createNameComponent(() => import('@/views/redirect.vue')),
        meta: { title: 'redirect', hideTabs: true }
      },
      {
        path: '/user_info',
        component: createNameComponent(() => import('@/views/userInfo.vue')),
        meta: { title: '用户信息', hideTabs: true }
      }
    ]
  },
  // 没有子节点不展示
  {
    path: '/login',
    component: createNameComponent(() => import('@/views/login.vue')),
    meta: { title: '登录', hideTabs: true },
    hidden: true
  },
  {
    path: '/update_password',
    component: createNameComponent(() => import('@/views/updatePassword.vue')),
    meta: { title: '修改密码', hideTabs: true },
    hidden: true
  },
  {
    path: '/reset_password',
    component: createNameComponent(() => import('@/views/resetPassword.vue')),
    meta: { title: '重置密码', hideTabs: true },
    hidden: true
  },
  {
    // 找不到路由重定向到404页面
    path: '/:pathMatch(.*)',
    component: Layout,
    redirect: '/404',
    hidden: true
  }
]

export default route
