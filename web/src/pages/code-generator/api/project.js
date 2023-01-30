import request from '@/utils/request'

export function projectList(params) {
  return request({
    url: '/project/list',
    method: 'get',
    params
  })
}

export function projectSelect(id) {
  return request({
    url: '/project/select',
    method: 'get',
    params: { id }
  })
}

export function projectInsert(data) {
  return request({
    url: '/project/insert',
    method: 'post',
    data
  })
}

export function projectUpdate(data) {
  return request({
    url: '/project/update',
    method: 'put',
    data
  })
}

export function projectDelete(ids) {
  return request({
    url: '/project/delete',
    method: 'delete',
    params: { ids }
  })
}

export function projectLastId() {
  return request({
    url: '/project/lastId',
    method: 'get'
  })
}
