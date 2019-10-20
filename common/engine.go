package common

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var engine *gin.Engine
var once sync.Once

func DefaultEngine() *gin.Engine {
	once.Do(func() {
		engine = gin.New()
	})
	return engine
}
