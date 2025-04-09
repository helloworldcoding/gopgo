// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLog is the golang structure of table hg_sys_log for DAO operations like Where/Data.
type SysLog struct {
	g.Meta     `orm:"table:hg_sys_log, do:true"`
	Id         interface{} //
	ReqId      interface{} //
	AppId      interface{} //
	MerchantId interface{} //
	MemberId   interface{} //
	Method     interface{} //
	Module     interface{} //
	Url        interface{} //
	GetData    *gjson.Json //
	PostData   *gjson.Json //
	HeaderData *gjson.Json //
	Ip         interface{} //
	ProvinceId interface{} //
	CityId     interface{} //
	ErrorCode  interface{} //
	ErrorMsg   interface{} //
	ErrorData  *gjson.Json //
	UserAgent  interface{} //
	TakeUpTime interface{} //
	Timestamp  interface{} //
	Status     interface{} //
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
