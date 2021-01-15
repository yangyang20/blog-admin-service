package admin

import (
	"blog-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

type UploadApi struct {}

var Upload = new(UploadApi)


func (upload *UploadApi)UploadImg(r *ghttp.Request)  {
	log :=g.Log()
	files := r.GetUploadFiles("image")
	visitpath :=g.Cfg().GetString("upload.visitpath")
	uploadPath := g.Cfg().GetString("upload.path")
	log.Print("时间",gtime.Date())
	err :=gfile.Mkdir(uploadPath+"/"+gtime.Date())
	if err !=nil {
		log.Print("有错误",err)
	}
	names, err := files.Save(uploadPath+"/"+gtime.Date(),true)
	if err != nil {
		response.JsonExit(r,response.ERROR,"上传失败",err)
	}
	url := "http://"+r.Host + visitpath +"/"+gtime.Date()+"/"+gconv.String(names[0]);
	response.JsonExit(r,response.SUCCESS,"上传成功",url)

}
