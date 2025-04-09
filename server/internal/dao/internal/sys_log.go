// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysLogDao is the data access object for the table hg_sys_log.
type SysLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysLogColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysLogColumns defines and stores column names for the table hg_sys_log.
type SysLogColumns struct {
	Id         string //
	ReqId      string //
	AppId      string //
	MerchantId string //
	MemberId   string //
	Method     string //
	Module     string //
	Url        string //
	GetData    string //
	PostData   string //
	HeaderData string //
	Ip         string //
	ProvinceId string //
	CityId     string //
	ErrorCode  string //
	ErrorMsg   string //
	ErrorData  string //
	UserAgent  string //
	TakeUpTime string //
	Timestamp  string //
	Status     string //
	CreatedAt  string //
	UpdatedAt  string //
}

// sysLogColumns holds the columns for the table hg_sys_log.
var sysLogColumns = SysLogColumns{
	Id:         "id",
	ReqId:      "req_id",
	AppId:      "app_id",
	MerchantId: "merchant_id",
	MemberId:   "member_id",
	Method:     "method",
	Module:     "module",
	Url:        "url",
	GetData:    "get_data",
	PostData:   "post_data",
	HeaderData: "header_data",
	Ip:         "ip",
	ProvinceId: "province_id",
	CityId:     "city_id",
	ErrorCode:  "error_code",
	ErrorMsg:   "error_msg",
	ErrorData:  "error_data",
	UserAgent:  "user_agent",
	TakeUpTime: "take_up_time",
	Timestamp:  "timestamp",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewSysLogDao creates and returns a new DAO object for table data access.
func NewSysLogDao(handlers ...gdb.ModelHandler) *SysLogDao {
	return &SysLogDao{
		group:    "default",
		table:    "hg_sys_log",
		columns:  sysLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysLogDao) Columns() SysLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
