// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenCurdDemo is the golang structure of table hg_sys_gen_curd_demo for DAO operations like Where/Data.
type SysGenCurdDemo struct {
	g.Meta      `orm:"table:hg_sys_gen_curd_demo, do:true"`
	Id          interface{} //
	Uuid        interface{} //
	CategoryId  interface{} //
	Title       interface{} //
	Description interface{} //
	Content     interface{} //
	Image       interface{} //
	Attachfile  interface{} //
	CityId      interface{} //
	Switch      interface{} //
	Sort        interface{} //
	Status      interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
	DeletedBy   interface{} //
	UpdatedBy   interface{} //
	CreatedBy   interface{} //
}
