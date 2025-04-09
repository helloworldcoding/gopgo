// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package sys

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysPodcast struct{}

func NewSysPodcast() *sSysPodcast {
	return &sSysPodcast{}
}

func init() {
	service.RegisterSysPodcast(NewSysPodcast())
}

// Model 播客管理ORM模型
func (s *sSysPodcast) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Podcast.Ctx(ctx), option...)
}

// List 获取播客管理列表
func (s *sSysPodcast) List(ctx context.Context, in *sysin.PodcastListInp) (list []*sysin.PodcastListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.PodcastListModel{})

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.Podcast.Columns().Id, in.Id)
	}

	// 查询uuid
	if in.Uuid != "" {
		mod = mod.Where(dao.Podcast.Columns().Uuid, in.Uuid)
	}

	// 查询审核状态
	if in.AuditStatus > 0 {
		mod = mod.Where(dao.Podcast.Columns().AuditStatus, in.AuditStatus)
	}

	// 查询上架状态
	if in.OnlineStatus > 0 {
		mod = mod.Where(dao.Podcast.Columns().OnlineStatus, in.OnlineStatus)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.Podcast.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.Podcast.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取播客管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出播客管理
func (s *sSysPodcast) Export(ctx context.Context, in *sysin.PodcastListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.PodcastExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出播客管理-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.PodcastExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增播客管理
func (s *sSysPodcast) Edit(ctx context.Context, in *sysin.PodcastEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.PodcastUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改播客管理失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.PodcastInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增播客管理失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除播客管理
func (s *sSysPodcast) Delete(ctx context.Context, in *sysin.PodcastDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除播客管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取播客管理指定信息
func (s *sSysPodcast) View(ctx context.Context, in *sysin.PodcastViewInp) (res *sysin.PodcastViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取播客管理信息，请稍后重试！")
		return
	}
	return
}