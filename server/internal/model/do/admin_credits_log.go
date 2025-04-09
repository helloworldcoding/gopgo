// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCreditsLog is the golang structure of table hg_admin_credits_log for DAO operations like Where/Data.
type AdminCreditsLog struct {
	g.Meta      `orm:"table:hg_admin_credits_log, do:true"`
	Id          interface{} //
	Uuid        interface{} //
	MemberId    interface{} //
	AppId       interface{} //
	AddonsName  interface{} //
	CreditType  interface{} //
	CreditGroup interface{} //
	BeforeNum   interface{} //
	Num         interface{} //
	AfterNum    interface{} //
	Remark      interface{} //
	Ip          interface{} //
	MapId       interface{} //
	Status      interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
