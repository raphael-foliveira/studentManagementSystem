package main

import (
	"github.com/raphael-foliveira/studentManagementSystem/db"
	"github.com/raphael-foliveira/studentManagementSystem/routes"
)

func main() {
	db.GetDb()
	routes.Start()
}
