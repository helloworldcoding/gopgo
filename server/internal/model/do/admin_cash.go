// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCash is the golang structure of table hg_admin_cash for DAO operations like Where/Data.
type AdminCash struct {
	g.Meta    `orm:"table:hg_admin_cash, do:true"`
	Id        interface{} //
	Uuid      interface{} //
	MemberId  interface{} //
	Money     interface{} //
	Fee       interface{} //
	LastMoney interface{} //
	Ip        interface{} //
	Status    interface{} //
	Msg       interface{} //
	HandleAt  *gtime.Time //
	CreatedAt *gtime.Time //
}
