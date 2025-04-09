// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysServeLicense is the golang structure for table sys_serve_license.
type SysServeLicense struct {
	Id           int64       `json:"id"           orm:"id"             description:""`
	Group        string      `json:"group"        orm:"group"          description:""`
	Name         string      `json:"name"         orm:"name"           description:""`
	Appid        string      `json:"appid"        orm:"appid"          description:""`
	SecretKey    string      `json:"secretKey"    orm:"secret_key"     description:""`
	RemoteAddr   string      `json:"remoteAddr"   orm:"remote_addr"    description:""`
	OnlineLimit  int         `json:"onlineLimit"  orm:"online_limit"   description:""`
	LoginTimes   int64       `json:"loginTimes"   orm:"login_times"    description:""`
	LastLoginAt  *gtime.Time `json:"lastLoginAt"  orm:"last_login_at"  description:""`
	LastActiveAt *gtime.Time `json:"lastActiveAt" orm:"last_active_at" description:""`
	Routes       *gjson.Json `json:"routes"       orm:"routes"         description:""`
	AllowedIps   string      `json:"allowedIps"   orm:"allowed_ips"    description:""`
	EndAt        *gtime.Time `json:"endAt"        orm:"end_at"         description:""`
	Remark       string      `json:"remark"       orm:"remark"         description:""`
	Status       int         `json:"status"       orm:"status"         description:""`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"     description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"     description:""`
}
