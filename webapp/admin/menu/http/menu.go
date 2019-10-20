package http

import (
	"github.com/gin-gonic/gin"
	log "github.com/jptangchina/log4g"
	"net/http"
	"tdpos/util"
	"tdpos/webapp/admin/menu/model"
)

func AddNewMenuItem(c *gin.Context) {
	menuitem := model.Menuitem{
		Status: 1,
	}
	if err := c.ShouldBindJSON(&menuitem); err != nil {
		log.Error("请求参数异常：", err)
		c.JSON(http.StatusBadRequest, util.BadRequest())
		return
	}
	svc.AddNewMenuItem(&menuitem)
	c.JSON(http.StatusOK, util.Success(nil))
}

func UpdateMenuItemInfo(c *gin.Context) {
	var menuitem model.Menuitem
	if err := c.ShouldBindJSON(&menuitem); err != nil {
		log.Error("请求参数异常：", err)
		c.JSON(http.StatusBadRequest, util.BadRequest())
		return
	}
	svc.UpdateMenuItemInfo(&menuitem)
	c.JSON(http.StatusOK, util.Success(nil))
}
