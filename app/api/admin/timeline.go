package admin

import (
	"blog-admin/app/define"
	"blog-admin/app/service"
	"blog-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
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

func (timeline *timelineApi) List(r *ghttp.Request) {
	res := timeline.timelineService.TimelineList()
	response.JsonExit(r, res.Code, res.Msg, res.Data)
}

func (timeline *timelineApi) PutInfo(r *ghttp.Request) {
	logger := glog.New()
	var (
		data *define.TimelinePut
	)
	err := r.Parse(&data)
	if err != nil {
		response.JsonExit(r, response.FAIL, "参数有误:"+gconv.String(err), err)
	}
	logger.Print("data", data)
	update := gconv.Map(data)
	res := timeline.timelineService.TimelineUpdate(update, "id="+gconv.String(data.Id))
	response.JsonExit(r, res.Code, res.Msg, res.Data)
}

func (timeline *timelineApi) PutShow(r *ghttp.Request) {
	logger := glog.New()
	id := r.GetRequest("id")
	logger.Print("id:", id)
	isShow := r.GetRequest("isShow")
	logger.Print("isshow:", isShow)
	data := make(map[string]interface{})
	data["is_show"] = gconv.Int(isShow)
	res := timeline.timelineService.TimelineUpdate(data, "id="+gconv.String(id))
	response.JsonExit(r, res.Code, res.Msg, res.Data)
}

func (timeline *timelineApi) GetInfo(r *ghttp.Request) {
	id := r.Get("id")
	logger := glog.New()
	logger.Print("id", id)
	where := make(map[string]interface{})
	where["id"] = id
	logger.Print("where", where)
	res := timeline.timelineService.TimelineInfo(where)
	response.JsonExit(r, res.Code, res.Msg, res.Data)
}
