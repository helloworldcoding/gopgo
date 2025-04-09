// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLoginLog is the golang structure of table hg_sys_login_log for DAO operations like Where/Data.
type SysLoginLog struct {
	g.Meta     `orm:"table:hg_sys_login_log, do:true"`
	Id         interface{} //
	ReqId      interface{} //
	MemberId   interface{} //
	Username   interface{} //
	Response   *gjson.Json //
	LoginAt    *gtime.Time //
	LoginIp    interface{} //
	ProvinceId interface{} //
	CityId     interface{} //
	UserAgent  interface{} //
	ErrMsg     interface{} //
	Status     interface{} //
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
