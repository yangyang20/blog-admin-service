package service

import (
	"blog-admin/app/dao"
	"blog-admin/app/define"
	"blog-admin/app/model"
	"blog-admin/library/response"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

type TimelineService struct {
}

var blog = new(model.BlogTimeline)

func (t *TimelineService) TimelineCreate(data *define.TimelineCreate) response.ResultType {
	logger := glog.New()
	logger.Print("输入参数", data)

	if err := gconv.Struct(data, &blog); err != nil {
		return response.ResultType{response.ERROR, "参数转换错误", err}
	}
	blog.Created = gtime.Now()
	res, err := dao.BlogTimeline.Insert(blog)
	if err != nil {
		logger.Print("插入的sql语句是", res)
		return response.ResultType{response.FAIL, "插入失败", ""}
	}
	return response.ResultType{response.SUCCESS, "插入成功", ""}
}
