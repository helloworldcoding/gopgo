// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMember is the golang structure for table admin_member.
type AdminMember struct {
	Id                 int64       `json:"id"                 orm:"id"                   description:""`
	Uuid               string      `json:"uuid"               orm:"uuid"                 description:""`
	DeptId             int64       `json:"deptId"             orm:"dept_id"              description:""`
	RoleId             int64       `json:"roleId"             orm:"role_id"              description:""`
	RealName           string      `json:"realName"           orm:"real_name"            description:""`
	Username           string      `json:"username"           orm:"username"             description:""`
	PasswordHash       string      `json:"passwordHash"       orm:"password_hash"        description:""`
	Salt               string      `json:"salt"               orm:"salt"                 description:""`
	PasswordResetToken string      `json:"passwordResetToken" orm:"password_reset_token" description:""`
	Integral           float64     `json:"integral"           orm:"integral"             description:""`
	Balance            float64     `json:"balance"            orm:"balance"              description:""`
	Avatar             string      `json:"avatar"             orm:"avatar"               description:""`
	Sex                int         `json:"sex"                orm:"sex"                  description:""`
	Qq                 string      `json:"qq"                 orm:"qq"                   description:""`
	Email              string      `json:"email"              orm:"email"                description:""`
	Mobile             string      `json:"mobile"             orm:"mobile"               description:""`
	Birthday           *gtime.Time `json:"birthday"           orm:"birthday"             description:""`
	CityId             int64       `json:"cityId"             orm:"city_id"              description:""`
	Address            string      `json:"address"            orm:"address"              description:""`
	Pid                int64       `json:"pid"                orm:"pid"                  description:""`
	Level              int         `json:"level"              orm:"level"                description:""`
	Tree               string      `json:"tree"               orm:"tree"                 description:""`
	InviteCode         string      `json:"inviteCode"         orm:"invite_code"          description:""`
	Cash               *gjson.Json `json:"cash"               orm:"cash"                 description:""`
	LastActiveAt       *gtime.Time `json:"lastActiveAt"       orm:"last_active_at"       description:""`
	Remark             string      `json:"remark"             orm:"remark"               description:""`
	Status             int         `json:"status"             orm:"status"               description:""`
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:""`
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"           description:""`
}
