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
var logger = glog.New()

func (t *TimelineService) TimelineCreate(data *define.TimelineCreate) response.ResultType {

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

func (t *TimelineService) TimelineList() response.ResultType {

	res, err := dao.BlogTimeline.All()
	if err != nil {
		logger.Print("有错误")
		return response.ResultType{response.FAIL, "查询错误", err}
	}
	return response.ResultType{response.SUCCESS, "查询成功", res}
}

func (t *TimelineService) TimelineUpdate(data map[string]interface{}, where string) response.ResultType {
	logger.Print("查询条件", where)
	res, err := dao.BlogTimeline.Where(where).FindOne()
	if err != nil || res == nil {
		logger.Print("查询结果", res)
		logger.Print("查询条件有错误", err)
		return response.ResultType{response.FAIL, "查询错误", err}
	}

	_, upt_error := dao.BlogTimeline.Update(data, where)
	if upt_error != nil {
		logger.Print("更新错误")
		return response.ResultType{response.ERROR, "更新错误", upt_error}
	}
	return response.ResultType{response.SUCCESS, "更新成功", nil}
}

func (t *TimelineService) TimelineInfo(where map[string]interface{}) response.ResultType {
	logger.Print("查询参数", where)
	res, err := dao.BlogTimeline.Where(where).FindOne()
	if err != nil || res == nil {
		logger.Panic("查询错误:", err)
		return response.ResultType{response.FAIL, "查询错误", err}
	}
	return response.ResultType{response.SUCCESS, "查询成功", res}
}
