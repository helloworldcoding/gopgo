// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminPost is the golang structure for table admin_post.
type AdminPost struct {
	Id        int64       `json:"id"        orm:"id"         description:""`
	Code      string      `json:"code"      orm:"code"       description:""`
	Name      string      `json:"name"      orm:"name"       description:""`
	Remark    string      `json:"remark"    orm:"remark"     description:""`
	Sort      int         `json:"sort"      orm:"sort"       description:""`
	Status    int         `json:"status"    orm:"status"     description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
