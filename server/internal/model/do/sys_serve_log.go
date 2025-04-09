// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysServeLog is the golang structure of table hg_sys_serve_log for DAO operations like Where/Data.
type SysServeLog struct {
	g.Meta      `orm:"table:hg_sys_serve_log, do:true"`
	Id          interface{} //
	TraceId     interface{} //
	LevelFormat interface{} //
	Content     interface{} //
	Stack       *gjson.Json //
	Line        interface{} //
	TriggerNs   interface{} //
	Status      interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
