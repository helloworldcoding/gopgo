// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminDept is the golang structure of table hg_admin_dept for DAO operations like Where/Data.
type AdminDept struct {
	g.Meta    `orm:"table:hg_admin_dept, do:true"`
	Id        interface{} //
	Uuid      interface{} //
	Pid       interface{} //
	Name      interface{} //
	Code      interface{} //
	Type      interface{} //
	Leader    interface{} //
	Phone     interface{} //
	Email     interface{} //
	Level     interface{} //
	Tree      interface{} //
	Sort      interface{} //
	Status    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
