// Package podcast
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package podcast

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询播客管理列表
type ListReq struct {
	g.Meta `path:"/podcast/list" method:"get" tags:"播客管理" summary:"获取播客管理列表"`
	sysin.PodcastListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.PodcastListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出播客管理列表
type ExportReq struct {
	g.Meta `path:"/podcast/export" method:"get" tags:"播客管理" summary:"导出播客管理列表"`
	sysin.PodcastListInp
}

type ExportRes struct{}

// ViewReq 获取播客管理指定信息
type ViewReq struct {
	g.Meta `path:"/podcast/view" method:"get" tags:"播客管理" summary:"获取播客管理指定信息"`
	sysin.PodcastViewInp
}

type ViewRes struct {
	*sysin.PodcastViewModel
}

// EditReq 修改/新增播客管理
type EditReq struct {
	g.Meta `path:"/podcast/edit" method:"post" tags:"播客管理" summary:"修改/新增播客管理"`
	sysin.PodcastEditInp
}

type EditRes struct{}

// DeleteReq 删除播客管理
type DeleteReq struct {
	g.Meta `path:"/podcast/delete" method:"post" tags:"播客管理" summary:"删除播客管理"`
	sysin.PodcastDeleteInp
}

type DeleteRes struct{}