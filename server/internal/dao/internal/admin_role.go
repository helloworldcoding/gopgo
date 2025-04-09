// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleDao is the data access object for the table hg_admin_role.
type AdminRoleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminRoleColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminRoleColumns defines and stores column names for the table hg_admin_role.
type AdminRoleColumns struct {
	Id            string // 主键
	Key           string // 角色唯一标识
	Name          string // 角色名称
	DefaultRouter string // 默认路由
	DataScope     string // 数据范围（1：全部数据权限 2：自定义数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：本人数据权限）
	CustomDept    string // 自定义部门权限
	Sort          string // 排序
	Pid           string // 父ID
	Level         string // 级别
	Tree          string // 树结构
	Remark        string // 备注
	Status        string // 状态
	CreatedBy     string // 创建者
	UpdatedBy     string // 更新者
	CreatedAt     string // 创建时间
	UpdatedAt     string // 修改时间
	DeletedAt     string // 删除时间
	Uuid          string // 唯一标识
}

// adminRoleColumns holds the columns for the table hg_admin_role.
var adminRoleColumns = AdminRoleColumns{
	Id:            "id",
	Key:           "key",
	Name:          "name",
	DefaultRouter: "default_router",
	DataScope:     "data_scope",
	CustomDept:    "custom_dept",
	Sort:          "sort",
	Pid:           "pid",
	Level:         "level",
	Tree:          "tree",
	Remark:        "remark",
	Status:        "status",
	CreatedBy:     "created_by",
	UpdatedBy:     "updated_by",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	Uuid:          "uuid",
}

// NewAdminRoleDao creates and returns a new DAO object for table data access.
func NewAdminRoleDao(handlers ...gdb.ModelHandler) *AdminRoleDao {
	return &AdminRoleDao{
		group:    "default",
		table:    "hg_admin_role",
		columns:  adminRoleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminRoleDao) Columns() AdminRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminRoleDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AdminRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
