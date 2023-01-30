import request from '@/utils/request'

// 首页数据
export function HomeApi(params) {
  return request({
    url: '/dashboard/home',
    method: 'get',
    baseURL: '/mock',
    params
  })
}
