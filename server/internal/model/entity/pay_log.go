// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayLog is the golang structure for table pay_log.
type PayLog struct {
	Id            int64       `json:"id"            orm:"id"             description:""`
	MemberId      int64       `json:"memberId"      orm:"member_id"      description:""`
	AppId         string      `json:"appId"         orm:"app_id"         description:""`
	AddonsName    string      `json:"addonsName"    orm:"addons_name"    description:""`
	OrderSn       string      `json:"orderSn"       orm:"order_sn"       description:""`
	OrderGroup    string      `json:"orderGroup"    orm:"order_group"    description:""`
	TradeNo       string      `json:"tradeNo"       orm:"trade_no"       description:""`
	PayType       string      `json:"payType"       orm:"pay_type"       description:""`
	PayAmount     float64     `json:"payAmount"     orm:"pay_amount"     description:""`
	OutTradeNo    string      `json:"outTradeNo"    orm:"out_trade_no"   description:""`
	TransactionId string      `json:"transactionId" orm:"transaction_id" description:""`
	Openid        string      `json:"openid"        orm:"openid"         description:""`
	MchId         string      `json:"mchId"         orm:"mch_id"         description:""`
	Subject       string      `json:"subject"       orm:"subject"        description:""`
	Detail        *gjson.Json `json:"detail"        orm:"detail"         description:""`
	ActualAmount  float64     `json:"actualAmount"  orm:"actual_amount"  description:""`
	PayStatus     int         `json:"payStatus"     orm:"pay_status"     description:""`
	PayAt         *gtime.Time `json:"payAt"         orm:"pay_at"         description:""`
	TradeType     string      `json:"tradeType"     orm:"trade_type"     description:""`
	RefundSn      string      `json:"refundSn"      orm:"refund_sn"      description:""`
	IsRefund      int         `json:"isRefund"      orm:"is_refund"      description:""`
	Custom        string      `json:"custom"        orm:"custom"         description:""`
	AuthCode      string      `json:"authCode"      orm:"auth_code"      description:""`
	CreateIp      string      `json:"createIp"      orm:"create_ip"      description:""`
	PayIp         string      `json:"payIp"         orm:"pay_ip"         description:""`
	NotifyUrl     string      `json:"notifyUrl"     orm:"notify_url"     description:""`
	ReturnUrl     string      `json:"returnUrl"     orm:"return_url"     description:""`
	TraceIds      *gjson.Json `json:"traceIds"      orm:"trace_ids"      description:""`
	Status        int         `json:"status"        orm:"status"         description:""`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:""`
	Uuid          string      `json:"uuid"          orm:"uuid"           description:""`
}
