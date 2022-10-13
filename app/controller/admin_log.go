package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pear-admin-golang/app/global/request"
	"pear-admin-golang/app/global/response"
	"pear-admin-golang/app/service"
)

func LogList(c *gin.Context) {
	c.HTML(http.StatusOK, "log_list.html", gin.H{})
}

func LogLogin(c *gin.Context) {
	var f request.LayerListForm
	if err := c.ShouldBind(&f); err != nil {
		response.SuccessResp(c)
	}
}

func LogOperate(c *gin.Context) {
	var f request.LayerListForm
	if err := c.ShouldBind(&f); err != nil {
		response.SuccessResp(c).SetCode(0).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	count, data, err := service.OperLogListJsonService(f)
	if err != nil {
		response.SuccessResp(c).SetCode(0).SetMsg(err.Error()).SetCount(count).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetCode(0).SetCount(count).SetData(data).WriteJsonExit()
}
