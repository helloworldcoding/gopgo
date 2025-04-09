// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminNotice is the golang structure for table admin_notice.
type AdminNotice struct {
	Id        int64       `json:"id"        orm:"id"         description:""`
	Title     string      `json:"title"     orm:"title"      description:""`
	Type      int64       `json:"type"      orm:"type"       description:""`
	Tag       int         `json:"tag"       orm:"tag"        description:""`
	Content   string      `json:"content"   orm:"content"    description:""`
	Receiver  *gjson.Json `json:"receiver"  orm:"receiver"   description:""`
	Remark    string      `json:"remark"    orm:"remark"     description:""`
	Sort      int         `json:"sort"      orm:"sort"       description:""`
	Status    int         `json:"status"    orm:"status"     description:""`
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:""`
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
