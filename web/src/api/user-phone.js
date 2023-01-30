import request from '@/utils/request'

// 手机号登录api
export function checkPhoneApi(data) {
  return request({
    url: '/user/v1.0/phone/check_phone',
    method: 'POST',
    data
  })
}

// 手机号登录api
export function registerApi(data) {
  return request({
    url: '/user/v1.0/phone/register',
    method: 'POST',
    data
  })
}

// 手机号登录api
export function loginPhoneApi(data) {
  return request({
    url: '/user/v1.0/phone/login_phone',
    method: 'POST',
    data
  })
}

// 重置密码校验
export function resetPasswordCheck(data) {
  return request({
    url: '/user/v1.0/phone/reset_password_check',
    method: 'POST',
    data
  })
}

// 重置密码
export function resetPassword(data) {
  return request({
    url: '/user/v1.0/phone/reset_password',
    method: 'POST',
    data
  })
}
