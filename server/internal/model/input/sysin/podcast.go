// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package sysin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PodcastUpdateFields 修改播客管理字段过滤
type PodcastUpdateFields struct {
	Title        string      `json:"title"        dc:"标题"`
	Description  string      `json:"description"  dc:"摘要描述"`
	CoverUrl     string      `json:"coverUrl"     dc:"封面图"`
	AuthorId     string      `json:"authorId"     dc:"作者"`
	Duration     int         `json:"duration"     dc:"时长(s)"`
	AudioUrl     string      `json:"audioUrl"     dc:"音频"`
	Category     string      `json:"category"     dc:"category"`
	Platform     string      `json:"platform"     dc:"platform"`
	OriginUrl    string      `json:"originUrl"    dc:"来源地址"`
	Content      string      `json:"content"      dc:"内容"`
	Scripts      string      `json:"scripts"      dc:"脚本"`
	Zhubos       *gjson.Json `json:"zhubos"       dc:"主播配置"`
	Tags         string      `json:"tags"         dc:"tags"`
	AuditStatus  int         `json:"auditStatus"  dc:"审核状态"`
	OnlineStatus int         `json:"onlineStatus" dc:"上架状态"`
	AuditRemark  string      `json:"auditRemark"  dc:"审核备注"`
	PlayCount    int         `json:"playCount"    dc:"播放次数"`
	LikeCount    int         `json:"likeCount"    dc:"点赞次数"`
	DislikeCount int         `json:"dislikeCount" dc:"不喜欢次数"`
	ShowCount    int         `json:"showCount"    dc:"展示次数"`
}

// PodcastInsertFields 新增播客管理字段过滤
type PodcastInsertFields struct {
	Title        string      `json:"title"        dc:"标题"`
	Description  string      `json:"description"  dc:"摘要描述"`
	CoverUrl     string      `json:"coverUrl"     dc:"封面图"`
	AuthorId     string      `json:"authorId"     dc:"作者"`
	Duration     int         `json:"duration"     dc:"时长(s)"`
	AudioUrl     string      `json:"audioUrl"     dc:"音频"`
	Category     string      `json:"category"     dc:"category"`
	Platform     string      `json:"platform"     dc:"platform"`
	OriginUrl    string      `json:"originUrl"    dc:"来源地址"`
	Content      string      `json:"content"      dc:"内容"`
	Scripts      string      `json:"scripts"      dc:"脚本"`
	Zhubos       *gjson.Json `json:"zhubos"       dc:"主播配置"`
	Tags         string      `json:"tags"         dc:"tags"`
	AuditStatus  int         `json:"auditStatus"  dc:"审核状态"`
	OnlineStatus int         `json:"onlineStatus" dc:"上架状态"`
	AuditRemark  string      `json:"auditRemark"  dc:"审核备注"`
	PlayCount    int         `json:"playCount"    dc:"播放次数"`
	LikeCount    int         `json:"likeCount"    dc:"点赞次数"`
	DislikeCount int         `json:"dislikeCount" dc:"不喜欢次数"`
	ShowCount    int         `json:"showCount"    dc:"展示次数"`
}

// PodcastEditInp 修改/新增播客管理
type PodcastEditInp struct {
	entity.Podcast
}

func (in *PodcastEditInp) Filter(ctx context.Context) (err error) {
	// 验证标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证摘要描述
	if err := g.Validator().Rules("required").Data(in.Description).Messages("摘要描述不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证封面图
	if err := g.Validator().Rules("required").Data(in.CoverUrl).Messages("封面图不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证作者
	if err := g.Validator().Rules("required").Data(in.AuthorId).Messages("作者不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证时长(s)
	if err := g.Validator().Rules("required").Data(in.Duration).Messages("时长(s)不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证音频
	if err := g.Validator().Rules("required").Data(in.AudioUrl).Messages("音频不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证category
	if err := g.Validator().Rules("required").Data(in.Category).Messages("category不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证tags
	if err := g.Validator().Rules("required").Data(in.Tags).Messages("tags不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PodcastEditModel struct{}

// PodcastDeleteInp 删除播客管理
type PodcastDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PodcastDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PodcastDeleteModel struct{}

// PodcastViewInp 获取指定播客管理信息
type PodcastViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PodcastViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PodcastViewModel struct {
	entity.Podcast
}

// PodcastListInp 获取播客管理列表
type PodcastListInp struct {
	form.PageReq
	Id           int64         `json:"id"           dc:"id"`
	Uuid         string        `json:"uuid"         dc:"uuid"`
	AuditStatus  int           `json:"auditStatus"  dc:"审核状态"`
	OnlineStatus int           `json:"onlineStatus" dc:"上架状态"`
	CreatedAt    []*gtime.Time `json:"createdAt"    dc:"创建时间"`
}

func (in *PodcastListInp) Filter(ctx context.Context) (err error) {
	return
}

type PodcastListModel struct {
	Id           int64       `json:"id"           dc:"id"`
	Title        string      `json:"title"        dc:"标题"`
	Description  string      `json:"description"  dc:"摘要描述"`
	CoverUrl     string      `json:"coverUrl"     dc:"封面图"`
	AuthorId     string      `json:"authorId"     dc:"作者"`
	Duration     int         `json:"duration"     dc:"时长(s)"`
	AudioUrl     string      `json:"audioUrl"     dc:"音频"`
	Category     string      `json:"category"     dc:"category"`
	Platform     string      `json:"platform"     dc:"platform"`
	OriginUrl    string      `json:"originUrl"    dc:"来源地址"`
	Content      string      `json:"content"      dc:"内容"`
	Scripts      string      `json:"scripts"      dc:"脚本"`
	Tags         string      `json:"tags"         dc:"tags"`
	AuditStatus  int         `json:"auditStatus"  dc:"审核状态"`
	OnlineStatus int         `json:"onlineStatus" dc:"上架状态"`
	AuditRemark  string      `json:"auditRemark"  dc:"审核备注"`
	PlayCount    int         `json:"playCount"    dc:"播放次数"`
	LikeCount    int         `json:"likeCount"    dc:"点赞次数"`
	DislikeCount int         `json:"dislikeCount" dc:"不喜欢次数"`
	ShowCount    int         `json:"showCount"    dc:"展示次数"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"更新时间"`
}

// PodcastExportModel 导出播客管理
type PodcastExportModel struct {
	Id           int64       `json:"id"           dc:"id"`
	Uuid         string      `json:"uuid"         dc:"uuid"`
	Title        string      `json:"title"        dc:"标题"`
	Description  string      `json:"description"  dc:"摘要描述"`
	CoverUrl     string      `json:"coverUrl"     dc:"封面图"`
	AuthorId     string      `json:"authorId"     dc:"作者"`
	Duration     int         `json:"duration"     dc:"时长(s)"`
	AudioUrl     string      `json:"audioUrl"     dc:"音频"`
	Category     string      `json:"category"     dc:"category"`
	Platform     string      `json:"platform"     dc:"platform"`
	OriginUrl    string      `json:"originUrl"    dc:"来源地址"`
	Content      string      `json:"content"      dc:"内容"`
	Scripts      string      `json:"scripts"      dc:"脚本"`
	Tags         string      `json:"tags"         dc:"tags"`
	AuditStatus  int         `json:"auditStatus"  dc:"审核状态"`
	OnlineStatus int         `json:"onlineStatus" dc:"上架状态"`
	AuditRemark  string      `json:"auditRemark"  dc:"审核备注"`
	PlayCount    int         `json:"playCount"    dc:"播放次数"`
	LikeCount    int         `json:"likeCount"    dc:"点赞次数"`
	DislikeCount int         `json:"dislikeCount" dc:"不喜欢次数"`
	ShowCount    int         `json:"showCount"    dc:"展示次数"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"更新时间"`
}