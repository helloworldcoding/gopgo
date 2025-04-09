// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminDept is the golang structure for table admin_dept.
type AdminDept struct {
	Id        int64       `json:"id"        orm:"id"         description:""`
	Uuid      string      `json:"uuid"      orm:"uuid"       description:""`
	Pid       int64       `json:"pid"       orm:"pid"        description:""`
	Name      string      `json:"name"      orm:"name"       description:""`
	Code      string      `json:"code"      orm:"code"       description:""`
	Type      string      `json:"type"      orm:"type"       description:""`
	Leader    string      `json:"leader"    orm:"leader"     description:""`
	Phone     string      `json:"phone"     orm:"phone"      description:""`
	Email     string      `json:"email"     orm:"email"      description:""`
	Level     int         `json:"level"     orm:"level"      description:""`
	Tree      string      `json:"tree"      orm:"tree"       description:""`
	Sort      int         `json:"sort"      orm:"sort"       description:""`
	Status    int         `json:"status"    orm:"status"     description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
