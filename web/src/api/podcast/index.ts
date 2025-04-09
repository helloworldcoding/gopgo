import { http, jumpExport } from '@/utils/http/axios';

// 获取播客管理列表
export function List(params) {
  return http.request({
    url: '/podcast/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除播客管理
export function Delete(params) {
  return http.request({
    url: '/podcast/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑播客管理
export function Edit(params) {
  return http.request({
    url: '/podcast/edit',
    method: 'POST',
    params,
  });
}

// 获取播客管理指定详情
export function View(params) {
  return http.request({
    url: '/podcast/view',
    method: 'GET',
    params,
  });
}

// 导出播客管理
export function Export(params) {
  jumpExport('/podcast/export', params);
}