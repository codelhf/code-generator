import request from '@/utils/request'

export function typeLanguageList(params) {
  return request({
    url: '/typeLanguage/list',
    method: 'get',
    params
  })
}

export function typeLanguageSelect(id) {
  return request({
    url: '/typeLanguage/select',
    method: 'get',
    params: { id }
  })
}

export function typeLanguageInsert(data) {
  return request({
    url: '/typeLanguage/insert',
    method: 'post',
    data
  })
}

export function typeLanguageUpdate(data) {
  return request({
    url: '/typeLanguage/update',
    method: 'put',
    data
  })
}

export function typeLanguageDelete(ids) {
  return request({
    url: '/typeLanguage/delete',
    method: 'delete',
    params: { ids }
  })
}
