import request from '@/utils/request'

export function projectTableList(params) {
  return request({
    url: '/projectTable/all',
    method: 'get',
    params
  })
}

export function projectTableSelect(id, name) {
  return request({
    url: '/projectTable/select',
    method: 'get',
    params: { projectId: id, tableName: name }
  })
}

export function projectTableGenerate(id, tables) {
  return request({
    url: '/projectTable/generate?projectId=' + id,
    method: 'post',
    data: tables
  })
}
