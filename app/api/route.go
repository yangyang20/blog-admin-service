package api

import (
	"blog-admin/app/api/admin"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

// 后台系统初始化
func Init() {
	s := g.Server()

	// 前台系统自定义错误页面
	s.BindStatusHandler(401, func(r *ghttp.Request) {
		if !gstr.HasPrefix(r.URL.Path, "/admin") {
			//admin.View.Render401(r)
		}
	})
	s.BindStatusHandler(403, func(r *ghttp.Request) {
		if !gstr.HasPrefix(r.URL.Path, "/admin") {
			//service.View.Render403(r)
		}
	})
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		if !gstr.HasPrefix(r.URL.Path, "/admin") {
			//service.View.Render404(r)
		}
	})
	s.BindStatusHandler(500, func(r *ghttp.Request) {
		if !gstr.HasPrefix(r.URL.Path, "/admin") {
			//service.View.Render500(r)
		}
	})

	// 路由注册
	s.Group("/", func(group *ghttp.RouterGroup) {
		// 权限控制路由
		group.Middleware(MiddlewareCORS)
		group.ALLMap(g.Map{
			"/upload":  admin.Upload,          // 上传
			"/timeline": admin.Timeline,
		})

	})
}



func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}