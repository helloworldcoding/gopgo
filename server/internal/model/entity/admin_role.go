// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRole is the golang structure for table admin_role.
type AdminRole struct {
	Id            int64       `json:"id"            orm:"id"             description:"主键"`
	Key           string      `json:"key"           orm:"key"            description:"角色唯一标识"`
	Name          string      `json:"name"          orm:"name"           description:"角色名称"`
	DefaultRouter string      `json:"defaultRouter" orm:"default_router" description:"默认路由"`
	DataScope     int         `json:"dataScope"     orm:"data_scope"     description:"数据范围（1：全部数据权限 2：自定义数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：本人数据权限）"`
	CustomDept    *gjson.Json `json:"customDept"    orm:"custom_dept"    description:"自定义部门权限"`
	Sort          int         `json:"sort"          orm:"sort"           description:"排序"`
	Pid           int64       `json:"pid"           orm:"pid"            description:"父ID"`
	Level         int         `json:"level"         orm:"level"          description:"级别"`
	Tree          string      `json:"tree"          orm:"tree"           description:"树结构"`
	Remark        string      `json:"remark"        orm:"remark"         description:"备注"`
	Status        int         `json:"status"        orm:"status"         description:"状态"`
	CreatedBy     int         `json:"createdBy"     orm:"created_by"     description:"创建者"`
	UpdatedBy     int64       `json:"updatedBy"     orm:"updated_by"     description:"更新者"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"修改时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     description:"删除时间"`
	Uuid          string      `json:"uuid"          orm:"uuid"           description:"唯一标识"`
}
