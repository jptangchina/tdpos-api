package util

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/jptangchina/log4g"
	"github.com/spf13/viper"
)

func DefaultDB() *gorm.DB {
	db, err := gorm.Open(viper.GetString("mysql.dialect"), viper.GetString("mysql.url"))
	if err != nil {
		log.Panicf("Can not get connection with database, cause by: %v\n", err)
		return nil
	}
	db.SingularTable(true)
	db.LogMode(false)
	return db
}
