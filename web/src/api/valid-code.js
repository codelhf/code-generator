import request from '@/utils/request'

export function imageCode() {
  return request({
    url: '/validate/code/img',
    method: 'GET'
  })
}

export function smsCode(query) {
  return request({
    url: '/validate/code/sms/' + query.phone,
    method: 'GET'
  })
}
