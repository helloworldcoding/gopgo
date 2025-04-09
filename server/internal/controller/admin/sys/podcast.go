// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package sys

import (
	"context"
	"hotgo/api/admin/podcast"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	Podcast = cPodcast{}
)

type cPodcast struct{}

// List 查看播客管理列表
func (c *cPodcast) List(ctx context.Context, req *podcast.ListReq) (res *podcast.ListRes, err error) {
	list, totalCount, err := service.SysPodcast().List(ctx, &req.PodcastListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.PodcastListModel{}
	}

	res = new(podcast.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出播客管理列表
func (c *cPodcast) Export(ctx context.Context, req *podcast.ExportReq) (res *podcast.ExportRes, err error) {
	err = service.SysPodcast().Export(ctx, &req.PodcastListInp)
	return
}

// Edit 更新播客管理
func (c *cPodcast) Edit(ctx context.Context, req *podcast.EditReq) (res *podcast.EditRes, err error) {
	err = service.SysPodcast().Edit(ctx, &req.PodcastEditInp)
	return
}

// View 获取指定播客管理信息
func (c *cPodcast) View(ctx context.Context, req *podcast.ViewReq) (res *podcast.ViewRes, err error) {
	data, err := service.SysPodcast().View(ctx, &req.PodcastViewInp)
	if err != nil {
		return
	}

	res = new(podcast.ViewRes)
	res.PodcastViewModel = data
	return
}

// Delete 删除播客管理
func (c *cPodcast) Delete(ctx context.Context, req *podcast.DeleteReq) (res *podcast.DeleteRes, err error) {
	err = service.SysPodcast().Delete(ctx, &req.PodcastDeleteInp)
	return
}