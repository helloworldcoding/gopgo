// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCash is the golang structure for table admin_cash.
type AdminCash struct {
	Id        int64       `json:"id"        orm:"id"         description:""`
	Uuid      string      `json:"uuid"      orm:"uuid"       description:""`
	MemberId  int64       `json:"memberId"  orm:"member_id"  description:""`
	Money     float64     `json:"money"     orm:"money"      description:""`
	Fee       float64     `json:"fee"       orm:"fee"        description:""`
	LastMoney float64     `json:"lastMoney" orm:"last_money" description:""`
	Ip        string      `json:"ip"        orm:"ip"         description:""`
	Status    int         `json:"status"    orm:"status"     description:""`
	Msg       string      `json:"msg"       orm:"msg"        description:""`
	HandleAt  *gtime.Time `json:"handleAt"  orm:"handle_at"  description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
}
