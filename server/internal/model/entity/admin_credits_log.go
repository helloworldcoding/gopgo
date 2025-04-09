// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCreditsLog is the golang structure for table admin_credits_log.
type AdminCreditsLog struct {
	Id          int64       `json:"id"          orm:"id"           description:""`
	Uuid        string      `json:"uuid"        orm:"uuid"         description:""`
	MemberId    int64       `json:"memberId"    orm:"member_id"    description:""`
	AppId       string      `json:"appId"       orm:"app_id"       description:""`
	AddonsName  string      `json:"addonsName"  orm:"addons_name"  description:""`
	CreditType  string      `json:"creditType"  orm:"credit_type"  description:""`
	CreditGroup string      `json:"creditGroup" orm:"credit_group" description:""`
	BeforeNum   float64     `json:"beforeNum"   orm:"before_num"   description:""`
	Num         float64     `json:"num"         orm:"num"          description:""`
	AfterNum    float64     `json:"afterNum"    orm:"after_num"    description:""`
	Remark      string      `json:"remark"      orm:"remark"       description:""`
	Ip          string      `json:"ip"          orm:"ip"           description:""`
	MapId       int64       `json:"mapId"       orm:"map_id"       description:""`
	Status      int         `json:"status"      orm:"status"       description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""`
}
