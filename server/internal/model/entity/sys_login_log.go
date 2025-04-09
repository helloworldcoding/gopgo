// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLoginLog is the golang structure for table sys_login_log.
type SysLoginLog struct {
	Id         int64       `json:"id"         orm:"id"          description:""`
	ReqId      string      `json:"reqId"      orm:"req_id"      description:""`
	MemberId   int64       `json:"memberId"   orm:"member_id"   description:""`
	Username   string      `json:"username"   orm:"username"    description:""`
	Response   *gjson.Json `json:"response"   orm:"response"    description:""`
	LoginAt    *gtime.Time `json:"loginAt"    orm:"login_at"    description:""`
	LoginIp    string      `json:"loginIp"    orm:"login_ip"    description:""`
	ProvinceId int64       `json:"provinceId" orm:"province_id" description:""`
	CityId     int64       `json:"cityId"     orm:"city_id"     description:""`
	UserAgent  string      `json:"userAgent"  orm:"user_agent"  description:""`
	ErrMsg     string      `json:"errMsg"     orm:"err_msg"     description:""`
	Status     int         `json:"status"     orm:"status"      description:""`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
}
