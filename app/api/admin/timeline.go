package admin

import (
	"blog-admin/app/define"
	"blog-admin/app/service"
	"blog-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type timelineApi struct {
	timelineService *service.TimelineService
}

var Timeline = new(timelineApi)

func (timeline *timelineApi) Create(r *ghttp.Request) {
	//log :=g.Log()
	var (
		data *define.TimelineCreate
	)
	err := r.Parse(&data)
	if err != nil {
		response.JsonExit(r, response.FAIL, "参数有误:"+gconv.String(err), err)
	}
	res := timeline.timelineService.TimelineCreate(data)
	response.JsonExit(r, res.Code, res.Msg, res.Data)
}
