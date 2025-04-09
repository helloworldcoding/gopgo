// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonHgexampleTable is the golang structure of table hg_addon_hgexample_table for DAO operations like Where/Data.
type AddonHgexampleTable struct {
	g.Meta      `orm:"table:hg_addon_hgexample_table, do:true"`
	Id          interface{} //
	Uuid        interface{} //
	Pid         interface{} //
	Level       interface{} //
	Tree        interface{} //
	CategoryId  interface{} //
	Flag        *gjson.Json //
	Title       interface{} //
	Description interface{} //
	Content     interface{} //
	Image       interface{} //
	Images      *gjson.Json //
	Attachfile  interface{} //
	Attachfiles *gjson.Json //
	Map         *gjson.Json //
	Star        interface{} //
	Price       interface{} //
	Views       interface{} //
	ActivityAt  *gtime.Time //
	StartAt     *gtime.Time //
	EndAt       *gtime.Time //
	Switch      interface{} //
	Sort        interface{} //
	Avatar      interface{} //
	Sex         interface{} //
	Qq          interface{} //
	Email       interface{} //
	Mobile      interface{} //
	Hobby       *gjson.Json //
	Channel     interface{} //
	CityId      interface{} //
	Remark      interface{} //
	Status      interface{} //
	CreatedBy   interface{} //
	UpdatedBy   interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
}
