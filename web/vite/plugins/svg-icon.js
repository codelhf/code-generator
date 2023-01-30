import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'

export default function createSvgIcon(isBuild) {
  const iconDirs = []
  const path = require('path')
  iconDirs.push(path.resolve(process.cwd(), 'src/icons'))
  // 自定义icons
  const fs = require('fs')
  const pages = fs.readdirSync(path.resolve(process.cwd(), 'src/pages'))
  pages.map(dirname => {
    iconDirs.push(path.resolve(process.cwd(), `src/pages/${dirname}/icons`))
  })
  return createSvgIconsPlugin({
    iconDirs: iconDirs,
    symbolId: 'icon-[dir]-[name]',
    svgoOptions: isBuild
  })
}
