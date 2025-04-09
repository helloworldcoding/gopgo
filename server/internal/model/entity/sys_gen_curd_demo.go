// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenCurdDemo is the golang structure for table sys_gen_curd_demo.
type SysGenCurdDemo struct {
	Id          int64       `json:"id"          orm:"id"          description:""`
	Uuid        string      `json:"uuid"        orm:"uuid"        description:""`
	CategoryId  int64       `json:"categoryId"  orm:"category_id" description:""`
	Title       string      `json:"title"       orm:"title"       description:""`
	Description string      `json:"description" orm:"description" description:""`
	Content     string      `json:"content"     orm:"content"     description:""`
	Image       string      `json:"image"       orm:"image"       description:""`
	Attachfile  string      `json:"attachfile"  orm:"attachfile"  description:""`
	CityId      int64       `json:"cityId"      orm:"city_id"     description:""`
	Switch      int         `json:"switch"      orm:"switch"      description:""`
	Sort        int         `json:"sort"        orm:"sort"        description:""`
	Status      int         `json:"status"      orm:"status"      description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:""`
	DeletedBy   int64       `json:"deletedBy"   orm:"deleted_by"  description:""`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"  description:""`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"  description:""`
}
