package app

import (
	"blog-admin/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmode"
)

// 应用启动
func Run() {
	// 绑定Swagger Plugin
	s := g.Server()

	// 静态目录设置
	uploadPath := g.Cfg().GetString("upload.path")
	visitpath :=g.Cfg().GetString("upload.visitpath")
	//log := g.Log()
	//log.Print("路径",uploadPath)
	if uploadPath == "" {
		g.Log().Fatal("文件上传配置路径不能为空")
	}
	s.AddStaticPath(visitpath, uploadPath)

	// HOOK, 开发阶段禁止浏览器缓存,方便调试
	if gmode.IsDevelop() {

		s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			r.Response.Header().Set("Cache-Control", "no-store")
		})
	}

	// 业务系统初始化
	api.Init()

	// 启动Http Server
	s.Run()
}

