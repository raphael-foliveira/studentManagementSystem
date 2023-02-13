package db

import (
	"fmt"

	"github.com/raphael-foliveira/studentManagementSystem/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetDb() (err error) {
	dsn := "host=localhost user=postgres password=123 dbname=smsys port=5432 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	fmt.Println("running migrations...")
	err = Db.AutoMigrate(&models.Student{}, &models.Class{})
	if err != nil {
		return err
	}
	return nil
}
