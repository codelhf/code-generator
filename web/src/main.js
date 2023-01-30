import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/theme-chalk/display.css' // 引入基于断点的隐藏类
import 'element-plus/dist/index.css'
import 'normalize.css' // css初始化
import './style/common.scss' // 公共css
import App from './App.vue'
import router from './router'
import store from './store'
import i18n from './i18n'

import 'virtual:svg-icons-register'
import * as icons from '@element-plus/icons-vue/dist/index'
import useRenderIcon from './utils/components/render-icon'

const app = createApp(App)

for (const [key, component] of Object.entries(icons)) {
  app.component(key, component)
}
app.config.globalProperties.useRenderIcon = useRenderIcon
// app.config.performance = true

app.use(ElementPlus)
app.use(router)
app.use(store)
app.use(i18n)
app.mount('#app')
