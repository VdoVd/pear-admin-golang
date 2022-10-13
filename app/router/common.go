package router

import (
	"github.com/gin-gonic/gin"
	controller "pear-admin-golang/app/controller"
	"pear-admin-golang/app/middleware"
)

func CommonRouter(common *gin.Engine) {
	common.GET("/", middleware.CheckDefaultPage)
	common.GET("login", middleware.CheckLoginPage, controller.Login)
	common.POST("login", controller.LoginHandler)
	common.GET("logout", controller.Logout)
	common.GET("not_found", controller.NotFound)
	common.GET("captcha", controller.GetCaptcha)
	common.POST("captcha_verify", controller.CaptchaVerify)
}
