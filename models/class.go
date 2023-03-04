package models

import (
	"github.com/raphael-foliveira/studentManagementSystem/db"
	"gorm.io/gorm"
)

type Class struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `json:"name"`
	Students  []Student      `json:"students" gorm:"many2many:student_classes"`
	TeacherID uint           `json:"teacherId"`
	Teacher   Teacher        `json:"teacher,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (c *Class) All() []Class {
	var classes []Class
	db.Db.Model(c).Preload("Students").Preload("Teacher").Find(&classes)
	return classes
}

func (c *Class) Find(id uint) {
	db.Db.Model(c).Preload("Students").Preload("Teacher").First(c, id)
}

func (c *Class) Create() {
	db.Db.Create(c)
}

func (c *Class) Delete(id uint) {
	db.Db.Delete(c, id)
}
