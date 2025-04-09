// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSmsLog is the golang structure for table sys_sms_log.
type SysSmsLog struct {
	Id        int64       `json:"id"        orm:"id"         description:""`
	Event     string      `json:"event"     orm:"event"      description:""`
	Mobile    string      `json:"mobile"    orm:"mobile"     description:""`
	Code      string      `json:"code"      orm:"code"       description:""`
	Times     int64       `json:"times"     orm:"times"      description:""`
	Ip        string      `json:"ip"        orm:"ip"         description:""`
	Status    int         `json:"status"    orm:"status"     description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
