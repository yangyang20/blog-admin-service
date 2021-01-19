// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// BlogTimeline is the golang structure for table blog_timeline.
type BlogTimeline struct {
	Id      uint        `orm:"id,primary" json:"id"`      //
	Title   string      `orm:"title"      json:"title"`   //
	Content string      `orm:"content"    json:"content"` //
	Image   string      `orm:"image"      json:"image"`   //
	Created *gtime.Time `orm:"created"    json:"created"` //
	IsShow  int         `orm:"is_show"    json:"isShow"`  // 1显示  0 不显示
}
