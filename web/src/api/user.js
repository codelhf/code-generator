import request from '@/utils/request'

// 获取手机号api
export function getPhoneApi(data) {
  return request({
    url: '/user/v1.0/get_phone',
    method: 'POST',
    data
  })
}

// 注册api
export function registerApi(data) {
  return request({
    url: '/user/v1.0/register',
    method: 'POST',
    data
  })
}

// 登录api
export function loginApi(data) {
  return request({
    url: '/user/v1.0/login_password',
    method: 'post',
    data
  })
}

// 获取用户信息Api
export function getInfoApi(params) {
  return request({
    url: '/user/v1.0/get_userInfo',
    method: 'post',
    params
  })
}

// 退出登录Api
export function logoutApi() {
  return request({
    url: '/user/v1.0/logout',
    method: 'post'
  })
}

// 修改密码校验
export function updatePasswordCheck(data) {
  return request({
    url: '/user/v1.0/update_password_check',
    method: 'post',
    data
  })
}

// 修改密码
export function updatePassword(data) {
  return request({
    url: '/user/v1.0/update_password',
    method: 'post',
    data
  })
}

// 修改用户信息
export function updateUserInfo(data) {
  return request({
    url: '/user/v1.0/update_userInfo',
    method: 'post',
    data
  })
}
