// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayLog is the golang structure of table hg_pay_log for DAO operations like Where/Data.
type PayLog struct {
	g.Meta        `orm:"table:hg_pay_log, do:true"`
	Id            interface{} //
	MemberId      interface{} //
	AppId         interface{} //
	AddonsName    interface{} //
	OrderSn       interface{} //
	OrderGroup    interface{} //
	TradeNo       interface{} //
	PayType       interface{} //
	PayAmount     interface{} //
	OutTradeNo    interface{} //
	TransactionId interface{} //
	Openid        interface{} //
	MchId         interface{} //
	Subject       interface{} //
	Detail        *gjson.Json //
	ActualAmount  interface{} //
	PayStatus     interface{} //
	PayAt         *gtime.Time //
	TradeType     interface{} //
	RefundSn      interface{} //
	IsRefund      interface{} //
	Custom        interface{} //
	AuthCode      interface{} //
	CreateIp      interface{} //
	PayIp         interface{} //
	NotifyUrl     interface{} //
	ReturnUrl     interface{} //
	TraceIds      *gjson.Json //
	Status        interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
	Uuid          interface{} //
}
