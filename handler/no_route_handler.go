package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tdpos/util"
)

func NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, util.NotFound())
	return
}
