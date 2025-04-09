// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRole is the golang structure of table hg_admin_role for DAO operations like Where/Data.
type AdminRole struct {
	g.Meta        `orm:"table:hg_admin_role, do:true"`
	Id            interface{} // 主键
	Key           interface{} // 角色唯一标识
	Name          interface{} // 角色名称
	DefaultRouter interface{} // 默认路由
	DataScope     interface{} // 数据范围（1：全部数据权限 2：自定义数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：本人数据权限）
	CustomDept    *gjson.Json // 自定义部门权限
	Sort          interface{} // 排序
	Pid           interface{} // 父ID
	Level         interface{} // 级别
	Tree          interface{} // 树结构
	Remark        interface{} // 备注
	Status        interface{} // 状态
	CreatedBy     interface{} // 创建者
	UpdatedBy     interface{} // 更新者
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 修改时间
	DeletedAt     *gtime.Time // 删除时间
	Uuid          interface{} // 唯一标识
}
