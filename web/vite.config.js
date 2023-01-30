import { loadEnv } from 'vite'
import { resolve } from 'path'
import createVitePlugins from './vite'

// https://vitejs.dev/config/
export default ({ command, mode }) => {
  const env = loadEnv(mode, process.cwd())
  return {
    base: './',
    resolve: {
      alias: {
        '@': resolve(__dirname, '.', 'src')
      }
    },
    server: {
      port: 3002,
      open: false,
      proxy: { // 代理配置
        '/dev': {
          target: 'http://localhost:8888',
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/dev-api/, '')
        }
      }
    },
    build: {
      outDir: '../dist',
      rollupOptions: {
        output: {
          manualChunks: {
            'echarts': ['echarts']
          }
        }
      }
    },
    plugins: [
      ...createVitePlugins(env, command === 'build')
    ],
    css: {
      postcss: {
        plugins: [
          {
            postcssPlugin: 'internal:charset-removal',
            AtRule: {
              charset: (atRule) => {
                if (atRule.name === 'charset') {
                  atRule.remove()
                }
              }
            }
          }
        ]
      }
    }
  }
}
