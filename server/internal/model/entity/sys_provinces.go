// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProvinces is the golang structure for table sys_provinces.
type SysProvinces struct {
	Id        int64       `json:"id"        orm:"id"         description:""`
	Title     string      `json:"title"     orm:"title"      description:""`
	Pinyin    string      `json:"pinyin"    orm:"pinyin"     description:""`
	Lng       string      `json:"lng"       orm:"lng"        description:""`
	Lat       string      `json:"lat"       orm:"lat"        description:""`
	Pid       int64       `json:"pid"       orm:"pid"        description:""`
	Level     int         `json:"level"     orm:"level"      description:""`
	Tree      string      `json:"tree"      orm:"tree"       description:""`
	Sort      int         `json:"sort"      orm:"sort"       description:""`
	Status    int         `json:"status"    orm:"status"     description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
