import request from '@/utils/request'

export function templateFieldList(params) {
  return request({
    url: '/templateField/list',
    method: 'get',
    params
  })
}

export function templateFieldSelect(id) {
  return request({
    url: '/templateField/select',
    method: 'get',
    params: { id }
  })
}

export function templateFieldCheck(data) {
  return request({
    url: '/templateField/check',
    method: 'post',
    data
  })
}

export function templateFieldInsert(data) {
  return request({
    url: '/templateField/insert',
    method: 'post',
    data
  })
}

export function templateFieldUpdate(data) {
  return request({
    url: '/templateField/update',
    method: 'put',
    data
  })
}

export function templateFieldDelete(ids) {
  return request({
    url: '/templateField/delete',
    method: 'delete',
    params: { ids }
  })
}
