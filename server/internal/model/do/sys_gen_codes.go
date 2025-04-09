// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenCodes is the golang structure of table hg_sys_gen_codes for DAO operations like Where/Data.
type SysGenCodes struct {
	g.Meta        `orm:"table:hg_sys_gen_codes, do:true"`
	Id            interface{} //
	GenType       interface{} //
	GenTemplate   interface{} //
	VarName       interface{} //
	Options       *gjson.Json //
	DbName        interface{} //
	TableName     interface{} //
	TableComment  interface{} //
	DaoName       interface{} //
	MasterColumns *gjson.Json //
	AddonName     interface{} //
	Status        interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
