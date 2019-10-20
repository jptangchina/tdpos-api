package model

type Menuitem struct {
	Id     int     `json:"id" binding:"omitempty,min=1" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name   string  `json:"name" binding:"required"`
	Price  float32 `json:"price" binding:"required"`
	Type   int     `json:"type" binding:"min=0,max=1"`
	Status int     `json:"status" binding:"omitempty,min=-1,max=1"`
}

func (m *Menuitem) TableName() string {
	return "menuitem"
}
