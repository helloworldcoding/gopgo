// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMenu is the golang structure of table hg_admin_menu for DAO operations like Where/Data.
type AdminMenu struct {
	g.Meta         `orm:"table:hg_admin_menu, do:true"`
	Id             interface{} //
	Pid            interface{} //
	Level          interface{} //
	Tree           interface{} //
	Title          interface{} //
	Name           interface{} //
	Path           interface{} //
	Icon           interface{} //
	Type           interface{} //
	Redirect       interface{} //
	Permissions    interface{} //
	PermissionName interface{} //
	Component      interface{} //
	AlwaysShow     interface{} //
	ActiveMenu     interface{} //
	IsRoot         interface{} //
	IsFrame        interface{} //
	FrameSrc       interface{} //
	KeepAlive      interface{} //
	Hidden         interface{} //
	Affix          interface{} //
	Sort           interface{} //
	Remark         interface{} //
	Status         interface{} //
	UpdatedAt      *gtime.Time //
	CreatedAt      *gtime.Time //
}
