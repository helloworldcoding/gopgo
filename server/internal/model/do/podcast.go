// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Podcast is the golang structure of table hg_podcast for DAO operations like Where/Data.
type Podcast struct {
	g.Meta       `orm:"table:hg_podcast, do:true"`
	Id           interface{} //
	Uuid         interface{} //
	Title        interface{} //
	Description  interface{} //
	CoverUrl     interface{} //
	AuthorId     interface{} //
	Duration     interface{} //
	AudioUrl     interface{} //
	Category     interface{} //
	Platform     interface{} //
	OriginUrl    interface{} //
	Content      interface{} //
	Scripts      []string    //
	Zhubos       *gjson.Json //
	Tags         []string    //
	AuditStatus  interface{} //
	OnlineStatus interface{} //
	AuditRemark  interface{} //
	PlayCount    interface{} //
	LikeCount    interface{} //
	DislikeCount interface{} //
	ShowCount    interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
}
