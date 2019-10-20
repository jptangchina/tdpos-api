package service

import (
	log "github.com/jptangchina/log4g"
	"tdpos/common"
	"tdpos/webapp/admin/menu/model"
)

func (s *Service) AddNewMenuItem(menuitem *model.Menuitem) {
	menuitem.Id = 0
	if err := s.dao.Save(menuitem); err != nil {
		log.Error("新增菜单失败：", err)
		panic(common.NewAppError(DatabaseError, "新增菜单操作失败"))
	}
}

func (s *Service) UpdateMenuItemInfo(menuitem *model.Menuitem) {
	dbMenuItem, err := s.dao.FindById(menuitem.Id)
	if err != nil {
		log.Error("查询数据失败：", err)
		panic(common.NewAppError(DatabaseError, "查询数据库失败"))
	}
	if dbMenuItem == nil {
		log.Errorf("更新数据失败：数据%+v不存在\n", *menuitem)
		return
	}
	if err := s.dao.Save(menuitem); err != nil {
		log.Error("更新菜单失败：", err)
		panic(common.NewAppError(DatabaseError, "更新菜单操作失败"))
	}
}
