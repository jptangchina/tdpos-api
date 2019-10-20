package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tdpos/util"
)

func NoMethodHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, util.NotFound())
	return
}
