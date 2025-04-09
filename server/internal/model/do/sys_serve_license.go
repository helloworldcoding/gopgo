// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysServeLicense is the golang structure of table hg_sys_serve_license for DAO operations like Where/Data.
type SysServeLicense struct {
	g.Meta       `orm:"table:hg_sys_serve_license, do:true"`
	Id           interface{} //
	Group        interface{} //
	Name         interface{} //
	Appid        interface{} //
	SecretKey    interface{} //
	RemoteAddr   interface{} //
	OnlineLimit  interface{} //
	LoginTimes   interface{} //
	LastLoginAt  *gtime.Time //
	LastActiveAt *gtime.Time //
	Routes       *gjson.Json //
	AllowedIps   interface{} //
	EndAt        *gtime.Time //
	Remark       interface{} //
	Status       interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
