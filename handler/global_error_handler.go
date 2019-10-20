package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/jptangchina/log4g"
	"net/http"
	"tdpos/common"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if appErr, ok := err.(*common.AppError); ok {
					log.Error(appErr.Response.Msg)
					c.JSON(appErr.HttpStatus, appErr.Response)
					return
				}
				if err, ok := err.(error); ok {
					log.Errorf("Unhandled Error, caused by: %v\n", err)
					c.JSON(http.StatusInternalServerError, "Unhandled Error")
					return
				}
				log.Errorf("Unknown Error, caused by: %v\n", err)
				c.JSON(http.StatusInternalServerError, "Unknown Error")
				return
			}
		}()
		c.Next()
	}
}
