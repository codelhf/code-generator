import Layout from '@/layout/index.vue'
import { createNameComponent } from '@/router/createNode'

const route = [
  {
    path: '/',
    component: Layout,
    redirect: '/home',
    meta: { title: 'dashboard', icon: 'home', hideClose: true },
    children: [
      {
        path: 'home',
        component: createNameComponent(() => import('@/pages/code-generator/views/project/generate.vue')),
        meta: { title: 'dashboard', icon: 'home', noCache: false, affix: true }
      }
    ]
  },

  {
    path: '/project',
    component: Layout,
    redirect: '/project/index',
    meta: { title: 'project', icon: 'project' },
    children: [
      {
        path: 'index',
        component: createNameComponent(() => import('@/pages/code-generator/views/project/index.vue')),
        meta: { title: 'project', icon: 'project', noCache: false }
      },
      {
        path: 'detail/:id',
        component: createNameComponent(() => import('@/pages/code-generator/views/project/generate.vue')),
        meta: { title: 'projectGenerate', icon: 'project', noCache: false },
        hidden: true
      }
    ]
  },

  {
    path: '/template',
    component: Layout,
    redirect: '/template/index',
    meta: { title: 'template', icon: 'template' },
    children: [
      {
        path: 'index',
        component: createNameComponent(() => import('@/pages/code-generator/views/template/index.vue')),
        meta: { title: 'template', icon: 'template', noCache: false }
      },
      {
        path: 'detail/:id',
        component: createNameComponent(() => import('@/pages/code-generator/views/template/detail.vue')),
        meta: { title: 'templateDetail', icon: 'template', noCache: false },
        hidden: true
      }
    ]
  },

  {
    path: '/templateField',
    component: Layout,
    redirect: '/templateField/index',
    meta: { title: 'templateField', icon: 'templateField' },
    children: [
      {
        path: 'index',
        component: createNameComponent(() => import('@/pages/code-generator/views/template-field/index.vue')),
        meta: { title: 'templateField', icon: 'templateField', noCache: false }
      }
    ]
  },

  {
    path: '/typeColumn',
    component: Layout,
    redirect: '/typeColumn/index',
    meta: { title: 'typeColumn', icon: 'typeColumn' },
    children: [
      {
        path: 'index',
        component: createNameComponent(() => import('@/pages/code-generator/views/type-column/index.vue')),
        meta: { title: 'typeColumn', icon: 'typeColumn', noCache: false }
      }
    ]
  }
]

export default route
