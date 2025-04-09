// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonHgexampleTable is the golang structure for table addon_hgexample_table.
type AddonHgexampleTable struct {
	Id          int64       `json:"id"          orm:"id"          description:""`
	Uuid        string      `json:"uuid"        orm:"uuid"        description:""`
	Pid         int64       `json:"pid"         orm:"pid"         description:""`
	Level       int         `json:"level"       orm:"level"       description:""`
	Tree        string      `json:"tree"        orm:"tree"        description:""`
	CategoryId  int         `json:"categoryId"  orm:"category_id" description:""`
	Flag        *gjson.Json `json:"flag"        orm:"flag"        description:""`
	Title       string      `json:"title"       orm:"title"       description:""`
	Description string      `json:"description" orm:"description" description:""`
	Content     string      `json:"content"     orm:"content"     description:""`
	Image       string      `json:"image"       orm:"image"       description:""`
	Images      *gjson.Json `json:"images"      orm:"images"      description:""`
	Attachfile  string      `json:"attachfile"  orm:"attachfile"  description:""`
	Attachfiles *gjson.Json `json:"attachfiles" orm:"attachfiles" description:""`
	Map         *gjson.Json `json:"map"         orm:"map"         description:""`
	Star        float64     `json:"star"        orm:"star"        description:""`
	Price       float64     `json:"price"       orm:"price"       description:""`
	Views       int         `json:"views"       orm:"views"       description:""`
	ActivityAt  *gtime.Time `json:"activityAt"  orm:"activity_at" description:""`
	StartAt     *gtime.Time `json:"startAt"     orm:"start_at"    description:""`
	EndAt       *gtime.Time `json:"endAt"       orm:"end_at"      description:""`
	Switch      int         `json:"switch"      orm:"switch"      description:""`
	Sort        int         `json:"sort"        orm:"sort"        description:""`
	Avatar      string      `json:"avatar"      orm:"avatar"      description:""`
	Sex         int         `json:"sex"         orm:"sex"         description:""`
	Qq          string      `json:"qq"          orm:"qq"          description:""`
	Email       string      `json:"email"       orm:"email"       description:""`
	Mobile      string      `json:"mobile"      orm:"mobile"      description:""`
	Hobby       *gjson.Json `json:"hobby"       orm:"hobby"       description:""`
	Channel     int         `json:"channel"     orm:"channel"     description:""`
	CityId      int64       `json:"cityId"      orm:"city_id"     description:""`
	Remark      string      `json:"remark"      orm:"remark"      description:""`
	Status      int         `json:"status"      orm:"status"      description:""`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"  description:""`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"  description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:""`
}
