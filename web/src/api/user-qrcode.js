import request from '@/utils/request'

// 扫码登录api
export function loginQRApi(data) {
  return request({
    url: '/user/v1.0/login_qrcode',
    method: 'post',
    data
  })
}
