import request from '@/utils/request'

export function templateGroupList(params) {
  return request({
    url: '/templateGroup/list',
    method: 'get',
    params
  })
}

export function templateGroupSelect(id) {
  return request({
    url: '/templateGroup/select',
    method: 'get',
    params: { id }
  })
}

export function templateGroupCheck(data) {
  return request({
    url: '/templateGroup/check',
    method: 'post',
    data
  })
}

export function templateGroupInsert(data) {
  return request({
    url: '/templateGroup/insert',
    method: 'post',
    data
  })
}

export function templateGroupUpdate(data) {
  return request({
    url: '/templateGroup/update',
    method: 'put',
    data
  })
}

export function templateGroupDelete(ids) {
  return request({
    url: '/templateGroup/delete',
    method: 'delete',
    params: { ids }
  })
}
