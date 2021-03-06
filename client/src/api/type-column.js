import request from '@/utils/request'

export function allDbType() {
  return request({
    url: '/typeDb/allDbType',
    method: 'get'
  })
}

export function typeColumnList(params) {
  return request({
    url: '/typeColumn/list',
    method: 'get',
    params
  })
}

export function typeColumnSelect(id) {
  return request({
    url: '/typeColumn/select',
    method: 'get',
    params: { id }
  })
}

export function typeColumnCheck(data) {
  return request({
    url: '/typeColumn/check',
    method: 'post',
    data
  })
}

export function typeColumnInsert(data) {
  return request({
    url: '/typeColumn/insert',
    method: 'post',
    data
  })
}

export function typeColumnUpdate(data) {
  return request({
    url: '/typeColumn/update',
    method: 'put',
    data
  })
}

export function typeColumnDelete(ids) {
  return request({
    url: '/typeColumn/delete',
    method: 'delete',
    params: { ids }
  })
}
