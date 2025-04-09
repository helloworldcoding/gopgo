// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysServeLog is the golang structure for table sys_serve_log.
type SysServeLog struct {
	Id          int64       `json:"id"          orm:"id"           description:""`
	TraceId     string      `json:"traceId"     orm:"trace_id"     description:""`
	LevelFormat string      `json:"levelFormat" orm:"level_format" description:""`
	Content     string      `json:"content"     orm:"content"      description:""`
	Stack       *gjson.Json `json:"stack"       orm:"stack"        description:""`
	Line        string      `json:"line"        orm:"line"         description:""`
	TriggerNs   int64       `json:"triggerNs"   orm:"trigger_ns"   description:""`
	Status      int         `json:"status"      orm:"status"       description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""`
}
