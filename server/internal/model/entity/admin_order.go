// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminOrder is the golang structure for table admin_order.
type AdminOrder struct {
	Id                 int64       `json:"id"                 orm:"id"                   description:""`
	MemberId           int64       `json:"memberId"           orm:"member_id"            description:""`
	OrderType          string      `json:"orderType"          orm:"order_type"           description:""`
	ProductId          int64       `json:"productId"          orm:"product_id"           description:""`
	OrderSn            string      `json:"orderSn"            orm:"order_sn"             description:""`
	Money              float64     `json:"money"              orm:"money"                description:""`
	Remark             string      `json:"remark"             orm:"remark"               description:""`
	RefundReason       string      `json:"refundReason"       orm:"refund_reason"        description:""`
	RejectRefundReason string      `json:"rejectRefundReason" orm:"reject_refund_reason" description:""`
	Status             int         `json:"status"             orm:"status"               description:""`
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:""`
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"           description:""`
}
