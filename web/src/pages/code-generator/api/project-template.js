import request from '@/utils/request'

export function projectTemplateList(params) {
  return request({
    url: '/projectTemplate/list',
    method: 'get',
    params
  })
}

export function projectTemplateSelect(id) {
  return request({
    url: '/projectTemplate/select',
    method: 'get',
    params: { id }
  })
}

export function projectTemplateInsert(data) {
  return request({
    url: '/projectTemplate/insert',
    method: 'post',
    data
  })
}

export function projectTemplateUpdate(data) {
  return request({
    url: '/projectTemplate/update',
    method: 'put',
    data
  })
}

export function projectTemplateDelete(ids) {
  return request({
    url: '/projectTemplate/delete',
    method: 'delete',
    params: { ids }
  })
}
