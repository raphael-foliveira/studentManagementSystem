package models

import "time"

type Student struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email"`
	Course         string    `json:"course"`
	CurrentClasses []Class   `json:"currentClasses" gorm:"many2many:student_classes"`
	Semester       int       `json:"semester"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
