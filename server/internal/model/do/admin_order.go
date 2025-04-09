// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminOrder is the golang structure of table hg_admin_order for DAO operations like Where/Data.
type AdminOrder struct {
	g.Meta             `orm:"table:hg_admin_order, do:true"`
	Id                 interface{} //
	MemberId           interface{} //
	OrderType          interface{} //
	ProductId          interface{} //
	OrderSn            interface{} //
	Money              interface{} //
	Remark             interface{} //
	RefundReason       interface{} //
	RejectRefundReason interface{} //
	Status             interface{} //
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
}
