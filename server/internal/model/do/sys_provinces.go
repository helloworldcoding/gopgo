// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProvinces is the golang structure of table hg_sys_provinces for DAO operations like Where/Data.
type SysProvinces struct {
	g.Meta    `orm:"table:hg_sys_provinces, do:true"`
	Id        interface{} //
	Title     interface{} //
	Pinyin    interface{} //
	Lng       interface{} //
	Lat       interface{} //
	Pid       interface{} //
	Level     interface{} //
	Tree      interface{} //
	Sort      interface{} //
	Status    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
