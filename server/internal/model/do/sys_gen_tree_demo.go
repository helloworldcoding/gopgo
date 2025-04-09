// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenTreeDemo is the golang structure of table hg_sys_gen_tree_demo for DAO operations like Where/Data.
type SysGenTreeDemo struct {
	g.Meta      `orm:"table:hg_sys_gen_tree_demo, do:true"`
	Id          interface{} //
	Pid         interface{} //
	Level       interface{} //
	Tree        interface{} //
	CategoryId  interface{} //
	Title       interface{} //
	Description interface{} //
	Sort        interface{} //
	Status      interface{} //
	CreatedBy   interface{} //
	UpdatedBy   interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
	Uuid        interface{} //
}
