package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"tdpos/common"
	"tdpos/handler"
)

func InitApp() {
	gin.SetMode(viper.GetString("profile"))
	engine := common.DefaultEngine()
	engine.NoMethod(handler.NoMethodHandler)
	engine.NoRoute(handler.NoRouteHandler)

	engine.Use(handler.GlobalErrorHandler())
}
