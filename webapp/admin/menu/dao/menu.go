package dao

import (
	"tdpos/webapp/admin/menu/model"
)

func (d *Dao) Save(menuitem *model.Menuitem) error {
	return d.db.Save(menuitem).Error
}

func (d *Dao) FindById(id int) (*model.Menuitem, error) {
	var menuitem model.Menuitem
	if err := d.db.Where("id = ?", id).Find(&menuitem).Error; err != nil {
		return nil, err
	}
	return &menuitem, nil
}
