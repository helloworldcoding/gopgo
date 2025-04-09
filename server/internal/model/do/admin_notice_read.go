// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminNoticeRead is the golang structure of table hg_admin_notice_read for DAO operations like Where/Data.
type AdminNoticeRead struct {
	g.Meta    `orm:"table:hg_admin_notice_read, do:true"`
	Id        interface{} //
	NoticeId  interface{} //
	MemberId  interface{} //
	Clicks    interface{} //
	UpdatedAt *gtime.Time //
	CreatedAt *gtime.Time //
}
