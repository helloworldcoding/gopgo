// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonHgexampleTableDao is the data access object for the table hg_addon_hgexample_table.
type AddonHgexampleTableDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  AddonHgexampleTableColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// AddonHgexampleTableColumns defines and stores column names for the table hg_addon_hgexample_table.
type AddonHgexampleTableColumns struct {
	Id          string //
	Uuid        string //
	Pid         string //
	Level       string //
	Tree        string //
	CategoryId  string //
	Flag        string //
	Title       string //
	Description string //
	Content     string //
	Image       string //
	Images      string //
	Attachfile  string //
	Attachfiles string //
	Map         string //
	Star        string //
	Price       string //
	Views       string //
	ActivityAt  string //
	StartAt     string //
	EndAt       string //
	Switch      string //
	Sort        string //
	Avatar      string //
	Sex         string //
	Qq          string //
	Email       string //
	Mobile      string //
	Hobby       string //
	Channel     string //
	CityId      string //
	Remark      string //
	Status      string //
	CreatedBy   string //
	UpdatedBy   string //
	CreatedAt   string //
	UpdatedAt   string //
	DeletedAt   string //
}

// addonHgexampleTableColumns holds the columns for the table hg_addon_hgexample_table.
var addonHgexampleTableColumns = AddonHgexampleTableColumns{
	Id:          "id",
	Uuid:        "uuid",
	Pid:         "pid",
	Level:       "level",
	Tree:        "tree",
	CategoryId:  "category_id",
	Flag:        "flag",
	Title:       "title",
	Description: "description",
	Content:     "content",
	Image:       "image",
	Images:      "images",
	Attachfile:  "attachfile",
	Attachfiles: "attachfiles",
	Map:         "map",
	Star:        "star",
	Price:       "price",
	Views:       "views",
	ActivityAt:  "activity_at",
	StartAt:     "start_at",
	EndAt:       "end_at",
	Switch:      "switch",
	Sort:        "sort",
	Avatar:      "avatar",
	Sex:         "sex",
	Qq:          "qq",
	Email:       "email",
	Mobile:      "mobile",
	Hobby:       "hobby",
	Channel:     "channel",
	CityId:      "city_id",
	Remark:      "remark",
	Status:      "status",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewAddonHgexampleTableDao creates and returns a new DAO object for table data access.
func NewAddonHgexampleTableDao(handlers ...gdb.ModelHandler) *AddonHgexampleTableDao {
	return &AddonHgexampleTableDao{
		group:    "default",
		table:    "hg_addon_hgexample_table",
		columns:  addonHgexampleTableColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AddonHgexampleTableDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AddonHgexampleTableDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AddonHgexampleTableDao) Columns() AddonHgexampleTableColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AddonHgexampleTableDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AddonHgexampleTableDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AddonHgexampleTableDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
