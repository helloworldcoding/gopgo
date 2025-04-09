// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AdminRoleMenu is the golang structure for table admin_role_menu.
type AdminRoleMenu struct {
	Id     int64 `json:"id"     orm:"id"      description:""`
	RoleId int64 `json:"roleId" orm:"role_id" description:""`
	MenuId int64 `json:"menuId" orm:"menu_id" description:""`
}
