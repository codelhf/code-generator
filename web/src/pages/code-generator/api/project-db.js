import request from '@/utils/request'

export function projectDbSelect(id) {
  return request({
    url: '/projectDb/select',
    method: 'get',
    params: { projectId: id }
  })
}

export function projectDbInsert(data) {
  return request({
    url: '/projectDb/insert',
    method: 'post',
    data
  })
}

export function projectDbUpdate(data) {
  return request({
    url: '/projectDb/update',
    method: 'put',
    data
  })
}
