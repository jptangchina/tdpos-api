package http

import (
	"github.com/gin-gonic/gin"
	"tdpos/common"
	"tdpos/webapp/admin/menu/service"
)

var svc *service.Service

func Init() {
	InitService()
	initRouter(common.DefaultEngine())
}

func InitService() {
	svc = service.New()
}

func initRouter(engine *gin.Engine) {
	group := engine.Group("/menu")

	group.POST("/save", AddNewMenuItem)
	group.PUT("/save", UpdateMenuItemInfo)
}
