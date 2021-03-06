import request from '@/utils/request'

export function templateList(params) {
  return request({
    url: '/template/list',
    method: 'get',
    params
  })
}

export function templateExport(params) {
  return request({
    url: '/template/export',
    method: 'get',
    responseType: 'blob', // important,
    params
  })
}

export function templateImport(formData) {
  return request({
    url: '/template/dump',
    method: 'post',
    data: formData
  })
}

export function templateSelect(id) {
  return request({
    url: '/template/select',
    method: 'get',
    params: { id }
  })
}

export function templateCheck(data) {
  return request({
    url: '/template/check',
    method: 'post',
    data
  })
}

export function templateInsert(data) {
  return request({
    url: '/template/insert',
    method: 'post',
    data
  })
}

export function templateUpdate(data) {
  return request({
    url: '/template/update',
    method: 'put',
    data
  })
}

export function templateDelete(ids) {
  return request({
    url: '/template/delete',
    method: 'delete',
    params: { ids }
  })
}
