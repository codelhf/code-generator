import vite from '@vitejs/plugin-vue'

import createRestart from './plugins/restart'
import createMock from './plugins/mock'
import createSvgIcon from './plugins/svg-icon'
// import createAutoImport from './plugins/auto-import'
import monacoEditor from './plugins/monaco-editor'

export default function createVitePlugins(viteEnv, isBuild = false) {
  const vitePlugins = [vite()]
  !isBuild && vitePlugins.push(createRestart())
  vitePlugins.push(createMock(viteEnv, isBuild))
  vitePlugins.push(createSvgIcon(isBuild))
  vitePlugins.push(monacoEditor(isBuild))
  return vitePlugins
}
