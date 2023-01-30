import ViteRestart from 'vite-plugin-restart'

export default function createRestart() {
  const pluginsFile = []
  const fs = require('fs')
  const plugins = fs.readdirSync('vite/plugins')
  plugins.map(dirname => {
    if (fs.statSync(`vite/plugins/${dirname}`).isFile()) {
      pluginsFile.push(`vite/plugins/${dirname}`)
    }
  })
  return ViteRestart({
    restart: pluginsFile
  })
}
