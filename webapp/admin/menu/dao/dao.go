package dao

import (
	"github.com/jinzhu/gorm"
	"tdpos/util"
)

type Dao struct {
	db *gorm.DB
}

func New() *Dao {
	return &Dao{util.DefaultDB()}
}
