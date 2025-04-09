// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PodcastDao is the data access object for the table hg_podcast.
type PodcastDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PodcastColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PodcastColumns defines and stores column names for the table hg_podcast.
type PodcastColumns struct {
	Id           string //
	Uuid         string //
	Title        string //
	Description  string //
	CoverUrl     string //
	AuthorId     string //
	Duration     string //
	AudioUrl     string //
	Category     string //
	Platform     string //
	OriginUrl    string //
	Content      string //
	Scripts      string //
	Zhubos       string //
	Tags         string //
	AuditStatus  string //
	OnlineStatus string //
	AuditRemark  string //
	PlayCount    string //
	LikeCount    string //
	DislikeCount string //
	ShowCount    string //
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
}

// podcastColumns holds the columns for the table hg_podcast.
var podcastColumns = PodcastColumns{
	Id:           "id",
	Uuid:         "uuid",
	Title:        "title",
	Description:  "description",
	CoverUrl:     "cover_url",
	AuthorId:     "author_id",
	Duration:     "duration",
	AudioUrl:     "audio_url",
	Category:     "category",
	Platform:     "platform",
	OriginUrl:    "origin_url",
	Content:      "content",
	Scripts:      "scripts",
	Zhubos:       "zhubos",
	Tags:         "tags",
	AuditStatus:  "audit_status",
	OnlineStatus: "online_status",
	AuditRemark:  "audit_remark",
	PlayCount:    "play_count",
	LikeCount:    "like_count",
	DislikeCount: "dislike_count",
	ShowCount:    "show_count",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewPodcastDao creates and returns a new DAO object for table data access.
func NewPodcastDao(handlers ...gdb.ModelHandler) *PodcastDao {
	return &PodcastDao{
		group:    "default",
		table:    "hg_podcast",
		columns:  podcastColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PodcastDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PodcastDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PodcastDao) Columns() PodcastColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PodcastDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PodcastDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PodcastDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
