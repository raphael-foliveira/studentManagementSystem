package models

import "github.com/raphael-foliveira/studentManagementSystem/db"

func MigrateAll() {
	db.Db.AutoMigrate(
		&Student{},
		&StudentClass{},
		&Teacher{},
		&Class{},
	)
}
