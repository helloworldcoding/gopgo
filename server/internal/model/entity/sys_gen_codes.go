// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenCodes is the golang structure for table sys_gen_codes.
type SysGenCodes struct {
	Id            int64       `json:"id"            orm:"id"             description:""`
	GenType       int         `json:"genType"       orm:"gen_type"       description:""`
	GenTemplate   int         `json:"genTemplate"   orm:"gen_template"   description:""`
	VarName       string      `json:"varName"       orm:"var_name"       description:""`
	Options       *gjson.Json `json:"options"       orm:"options"        description:""`
	DbName        string      `json:"dbName"        orm:"db_name"        description:""`
	TableName     string      `json:"tableName"     orm:"table_name"     description:""`
	TableComment  string      `json:"tableComment"  orm:"table_comment"  description:""`
	DaoName       string      `json:"daoName"       orm:"dao_name"       description:""`
	MasterColumns *gjson.Json `json:"masterColumns" orm:"master_columns" description:""`
	AddonName     string      `json:"addonName"     orm:"addon_name"     description:""`
	Status        int         `json:"status"        orm:"status"         description:""`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:""`
}
