// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSmsLog is the golang structure of table hg_sys_sms_log for DAO operations like Where/Data.
type SysSmsLog struct {
	g.Meta    `orm:"table:hg_sys_sms_log, do:true"`
	Id        interface{} //
	Event     interface{} //
	Mobile    interface{} //
	Code      interface{} //
	Times     interface{} //
	Ip        interface{} //
	Status    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
