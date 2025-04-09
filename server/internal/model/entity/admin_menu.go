// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMenu is the golang structure for table admin_menu.
type AdminMenu struct {
	Id             int64       `json:"id"             orm:"id"              description:""`
	Pid            int64       `json:"pid"            orm:"pid"             description:""`
	Level          int         `json:"level"          orm:"level"           description:""`
	Tree           string      `json:"tree"           orm:"tree"            description:""`
	Title          string      `json:"title"          orm:"title"           description:""`
	Name           string      `json:"name"           orm:"name"            description:""`
	Path           string      `json:"path"           orm:"path"            description:""`
	Icon           string      `json:"icon"           orm:"icon"            description:""`
	Type           int         `json:"type"           orm:"type"            description:""`
	Redirect       string      `json:"redirect"       orm:"redirect"        description:""`
	Permissions    string      `json:"permissions"    orm:"permissions"     description:""`
	PermissionName string      `json:"permissionName" orm:"permission_name" description:""`
	Component      string      `json:"component"      orm:"component"       description:""`
	AlwaysShow     int         `json:"alwaysShow"     orm:"always_show"     description:""`
	ActiveMenu     string      `json:"activeMenu"     orm:"active_menu"     description:""`
	IsRoot         int         `json:"isRoot"         orm:"is_root"         description:""`
	IsFrame        int         `json:"isFrame"        orm:"is_frame"        description:""`
	FrameSrc       string      `json:"frameSrc"       orm:"frame_src"       description:""`
	KeepAlive      int         `json:"keepAlive"      orm:"keep_alive"      description:""`
	Hidden         int         `json:"hidden"         orm:"hidden"          description:""`
	Affix          int         `json:"affix"          orm:"affix"           description:""`
	Sort           int         `json:"sort"           orm:"sort"            description:""`
	Remark         string      `json:"remark"         orm:"remark"          description:""`
	Status         int         `json:"status"         orm:"status"          description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""`
}
