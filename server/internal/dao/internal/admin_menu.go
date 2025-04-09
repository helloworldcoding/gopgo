// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminMenuDao is the data access object for the table hg_admin_menu.
type AdminMenuDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminMenuColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminMenuColumns defines and stores column names for the table hg_admin_menu.
type AdminMenuColumns struct {
	Id             string //
	Pid            string //
	Level          string //
	Tree           string //
	Title          string //
	Name           string //
	Path           string //
	Icon           string //
	Type           string //
	Redirect       string //
	Permissions    string //
	PermissionName string //
	Component      string //
	AlwaysShow     string //
	ActiveMenu     string //
	IsRoot         string //
	IsFrame        string //
	FrameSrc       string //
	KeepAlive      string //
	Hidden         string //
	Affix          string //
	Sort           string //
	Remark         string //
	Status         string //
	UpdatedAt      string //
	CreatedAt      string //
}

// adminMenuColumns holds the columns for the table hg_admin_menu.
var adminMenuColumns = AdminMenuColumns{
	Id:             "id",
	Pid:            "pid",
	Level:          "level",
	Tree:           "tree",
	Title:          "title",
	Name:           "name",
	Path:           "path",
	Icon:           "icon",
	Type:           "type",
	Redirect:       "redirect",
	Permissions:    "permissions",
	PermissionName: "permission_name",
	Component:      "component",
	AlwaysShow:     "always_show",
	ActiveMenu:     "active_menu",
	IsRoot:         "is_root",
	IsFrame:        "is_frame",
	FrameSrc:       "frame_src",
	KeepAlive:      "keep_alive",
	Hidden:         "hidden",
	Affix:          "affix",
	Sort:           "sort",
	Remark:         "remark",
	Status:         "status",
	UpdatedAt:      "updated_at",
	CreatedAt:      "created_at",
}

// NewAdminMenuDao creates and returns a new DAO object for table data access.
func NewAdminMenuDao(handlers ...gdb.ModelHandler) *AdminMenuDao {
	return &AdminMenuDao{
		group:    "default",
		table:    "hg_admin_menu",
		columns:  adminMenuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminMenuDao) Columns() AdminMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminMenuDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
