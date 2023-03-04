package models

import "gorm.io/gorm"

type StudentClass struct {
	ID        uint `gorm:"primaryKey"`
	StudentID uint
	EmailID   uint
	DeletedAt gorm.DeletedAt
}
