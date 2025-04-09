// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminNoticeRead is the golang structure for table admin_notice_read.
type AdminNoticeRead struct {
	Id        int64       `json:"id"        orm:"id"         description:""`
	NoticeId  int64       `json:"noticeId"  orm:"notice_id"  description:""`
	MemberId  int64       `json:"memberId"  orm:"member_id"  description:""`
	Clicks    int         `json:"clicks"    orm:"clicks"     description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
}
