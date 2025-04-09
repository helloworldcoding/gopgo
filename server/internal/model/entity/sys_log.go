// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLog is the golang structure for table sys_log.
type SysLog struct {
	Id         int64       `json:"id"         orm:"id"           description:""`
	ReqId      string      `json:"reqId"      orm:"req_id"       description:""`
	AppId      string      `json:"appId"      orm:"app_id"       description:""`
	MerchantId int64       `json:"merchantId" orm:"merchant_id"  description:""`
	MemberId   int64       `json:"memberId"   orm:"member_id"    description:""`
	Method     string      `json:"method"     orm:"method"       description:""`
	Module     string      `json:"module"     orm:"module"       description:""`
	Url        string      `json:"url"        orm:"url"          description:""`
	GetData    *gjson.Json `json:"getData"    orm:"get_data"     description:""`
	PostData   *gjson.Json `json:"postData"   orm:"post_data"    description:""`
	HeaderData *gjson.Json `json:"headerData" orm:"header_data"  description:""`
	Ip         string      `json:"ip"         orm:"ip"           description:""`
	ProvinceId int64       `json:"provinceId" orm:"province_id"  description:""`
	CityId     int64       `json:"cityId"     orm:"city_id"      description:""`
	ErrorCode  int         `json:"errorCode"  orm:"error_code"   description:""`
	ErrorMsg   string      `json:"errorMsg"   orm:"error_msg"    description:""`
	ErrorData  *gjson.Json `json:"errorData"  orm:"error_data"   description:""`
	UserAgent  string      `json:"userAgent"  orm:"user_agent"   description:""`
	TakeUpTime int64       `json:"takeUpTime" orm:"take_up_time" description:""`
	Timestamp  int64       `json:"timestamp"  orm:"timestamp"    description:""`
	Status     int         `json:"status"     orm:"status"       description:""`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"   description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"   description:""`
}
