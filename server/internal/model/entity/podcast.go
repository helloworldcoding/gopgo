// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// Podcast is the golang structure for table podcast.
type Podcast struct {
	Id           int64       `json:"id"           orm:"id"            description:""`
	Uuid         string      `json:"uuid"         orm:"uuid"          description:""`
	Title        string      `json:"title"        orm:"title"         description:""`
	Description  string      `json:"description"  orm:"description"   description:""`
	CoverUrl     string      `json:"coverUrl"     orm:"cover_url"     description:""`
	AuthorId     string      `json:"authorId"     orm:"author_id"     description:""`
	Duration     int         `json:"duration"     orm:"duration"      description:""`
	AudioUrl     string      `json:"audioUrl"     orm:"audio_url"     description:""`
	Category     string      `json:"category"     orm:"category"      description:""`
	Platform     string      `json:"platform"     orm:"platform"      description:""`
	OriginUrl    string      `json:"originUrl"    orm:"origin_url"    description:""`
	Content      string      `json:"content"      orm:"content"       description:""`
	Scripts      []string    `json:"scripts"      orm:"scripts"       description:""`
	Zhubos       *gjson.Json `json:"zhubos"       orm:"zhubos"        description:""`
	Tags         []string    `json:"tags"         orm:"tags"          description:""`
	AuditStatus  int         `json:"auditStatus"  orm:"audit_status"  description:""`
	OnlineStatus int         `json:"onlineStatus" orm:"online_status" description:""`
	AuditRemark  string      `json:"auditRemark"  orm:"audit_remark"  description:""`
	PlayCount    int         `json:"playCount"    orm:"play_count"    description:""`
	LikeCount    int         `json:"likeCount"    orm:"like_count"    description:""`
	DislikeCount int         `json:"dislikeCount" orm:"dislike_count" description:""`
	ShowCount    int         `json:"showCount"    orm:"show_count"    description:""`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""`
}
