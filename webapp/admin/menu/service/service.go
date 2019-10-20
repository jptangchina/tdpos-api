package service

import "tdpos/webapp/admin/menu/dao"

const (
	DatabaseError = 10000
)

type Service struct {
	dao *dao.Dao
}

func New() *Service {
	return &Service{dao.New()}
}
