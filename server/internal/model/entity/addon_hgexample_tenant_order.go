// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonHgexampleTenantOrder is the golang structure for table addon_hgexample_tenant_order.
type AddonHgexampleTenantOrder struct {
	Id          int64       `json:"id"          orm:"id"           description:""`
	Uuid        string      `json:"uuid"        orm:"uuid"         description:""`
	TenantId    int         `json:"tenantId"    orm:"tenant_id"    description:""`
	MerchantId  int64       `json:"merchantId"  orm:"merchant_id"  description:""`
	UserId      int64       `json:"userId"      orm:"user_id"      description:""`
	ProductName string      `json:"productName" orm:"product_name" description:""`
	OrderSn     string      `json:"orderSn"     orm:"order_sn"     description:""`
	Money       float64     `json:"money"       orm:"money"        description:""`
	Remark      string      `json:"remark"      orm:"remark"       description:""`
	Status      int         `json:"status"      orm:"status"       description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""`
}
