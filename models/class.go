package models

type Class struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Students []Student `json:"students" gorm:"many2many:student_classes"`
}
