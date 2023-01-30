import { createProdMockServer } from 'vite-plugin-mock/es/createProdMockServer'
import homeModule from './modules/pages/dashboard'
import userModule from './modules/system/user'

export function setupProdMockServer() {
  createProdMockServer([
    ...homeModule,
    ...userModule
  ])
}
