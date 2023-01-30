import { createStore, createLogger } from 'vuex'
import Persistent from './plugins/persistent'
import setting from '@/setting'

const files = import.meta.globEager('./modules/*.js')
Object.assign(files, setting.store)

const modules = {}
Object.keys(files).forEach((c) => {
  const module = files[c].default
  const moduleName = c.replace(/^\.\/(.*)\/(.*)\.\w+$/, '$2')
  modules[moduleName] = module
})

const debug = process.env.NODE_ENV !== 'production'

const persistent = Persistent({ key: 'vuex', modules, modulesKeys: {
  local: Object.keys(modules),
  session: []
}})

export default createStore({
  modules: {
    ...modules
  },
  strict: debug,
  plugins: debug ? [createLogger(), persistent] : [persistent]
})
