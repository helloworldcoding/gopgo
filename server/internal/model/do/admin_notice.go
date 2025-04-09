// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminNotice is the golang structure of table hg_admin_notice for DAO operations like Where/Data.
type AdminNotice struct {
	g.Meta    `orm:"table:hg_admin_notice, do:true"`
	Id        interface{} //
	Title     interface{} //
	Type      interface{} //
	Tag       interface{} //
	Content   interface{} //
	Receiver  *gjson.Json //
	Remark    interface{} //
	Sort      interface{} //
	Status    interface{} //
	CreatedBy interface{} //
	UpdatedBy interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
