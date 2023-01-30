import { viteMockServe } from 'vite-plugin-mock'

export default function createMock(env, isBuild) {
  const { VITE_BUILD_MOCK } = env
  return viteMockServe({
    mockPath: 'mock', // 你的mock文件地址
    localEnabled: !isBuild, // 开发环境
    prodEnabled: isBuild && VITE_BUILD_MOCK === 'true', // 生产环境 //  这样可以控制关闭mock的时候不让mock打包到最终代码内
    // 如果prodEnable设置为true，则在编译打包的时候，会把mock的文件打包进去，如果你不写injectFile，那就是默认注入到main.ts/main.js
    injectCode: `
              // 这个路径需要注意是在你要注入的文件下的引用路径
              import { setupProdMockServer } from './utils/mockProdServer';
              setupProdMockServer();
            `,
    // 在全局中注入代码,不配置的话默认是在src/main.js/main.ts
    // injectFile: resolve("src/main.tsx"),
    watchFiles: true,
    logger: true
  })
}
