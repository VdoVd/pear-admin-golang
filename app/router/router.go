package router

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"net/http"
	"pear-admin-golang/app/core/config"
	"pear-admin-golang/app/middleware"
	"pear-admin-golang/app/util/session"
)

func InitRouter(staticFs, templateFs embed.FS) *gin.Engine {
	gin.SetMode(config.Instance().App.RunMode)
	r := gin.New()

	t, _ := template.ParseFS(templateFs, "template/**/**/*.html")
	r.SetHTMLTemplate(t)

	r.Static(config.Instance().App.ImgUrlPath, config.Instance().App.ImgSavePath)
	r.Static("/runtime/file", "runtime/file")

	fads, _ := fs.Sub(staticFs, "static")
	r.StaticFS("/static", http.FS(fads))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(session.EnableCookieSession(config.Instance().App.JwtSecret))
	CommonRouter(r)
	SystemRouter(r)
	TaskRouter(r)
	return r
}
