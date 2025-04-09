// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminPost is the golang structure of table hg_admin_post for DAO operations like Where/Data.
type AdminPost struct {
	g.Meta    `orm:"table:hg_admin_post, do:true"`
	Id        interface{} //
	Code      interface{} //
	Name      interface{} //
	Remark    interface{} //
	Sort      interface{} //
	Status    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
