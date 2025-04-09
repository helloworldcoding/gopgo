// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMember is the golang structure of table hg_admin_member for DAO operations like Where/Data.
type AdminMember struct {
	g.Meta             `orm:"table:hg_admin_member, do:true"`
	Id                 interface{} //
	Uuid               interface{} //
	DeptId             interface{} //
	RoleId             interface{} //
	RealName           interface{} //
	Username           interface{} //
	PasswordHash       interface{} //
	Salt               interface{} //
	PasswordResetToken interface{} //
	Integral           interface{} //
	Balance            interface{} //
	Avatar             interface{} //
	Sex                interface{} //
	Qq                 interface{} //
	Email              interface{} //
	Mobile             interface{} //
	Birthday           *gtime.Time //
	CityId             interface{} //
	Address            interface{} //
	Pid                interface{} //
	Level              interface{} //
	Tree               interface{} //
	InviteCode         interface{} //
	Cash               *gjson.Json //
	LastActiveAt       *gtime.Time //
	Remark             interface{} //
	Status             interface{} //
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
}
