package models

import (
	"github.com/raphael-foliveira/studentManagementSystem/db"
	"gorm.io/gorm"
)

type Teacher struct {
	ID        uint           `gorm:"primaryKey"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Classes   []Class        `json:"classes,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (t *Teacher) All() []Teacher {
	var teachers []Teacher
	db.Db.Model(t).Preload("Classes").Find(&teachers)
	return teachers
}

func (t *Teacher) Create() {
	db.Db.Model(t).Preload("Classes").Create(t)
}

func (t *Teacher) Find(id uint) {
	db.Db.First(t, id)
}
func (t *Teacher) Delete(id uint) {
	db.Db.Delete(t, id)
}
