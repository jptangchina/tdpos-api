package initialize

import (
	"github.com/gin-gonic/gin"
	log "github.com/jptangchina/log4g"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigFile("conf/app.toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("Can not load config file, cause by: %v\n", err)
	}

	viper.SetDefault("port", 8081)
	viper.SetDefault("profile", gin.TestMode)
}
