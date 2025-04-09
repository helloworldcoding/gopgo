// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenTreeDemo is the golang structure for table sys_gen_tree_demo.
type SysGenTreeDemo struct {
	Id          int64       `json:"id"          orm:"id"          description:""`
	Pid         int64       `json:"pid"         orm:"pid"         description:""`
	Level       int         `json:"level"       orm:"level"       description:""`
	Tree        string      `json:"tree"        orm:"tree"        description:""`
	CategoryId  int64       `json:"categoryId"  orm:"category_id" description:""`
	Title       string      `json:"title"       orm:"title"       description:""`
	Description string      `json:"description" orm:"description" description:""`
	Sort        int         `json:"sort"        orm:"sort"        description:""`
	Status      int         `json:"status"      orm:"status"      description:""`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"  description:""`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"  description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:""`
	Uuid        string      `json:"uuid"        orm:"uuid"        description:""`
}
