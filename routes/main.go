package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/db"
	"github.com/raphael-foliveira/studentManagementSystem/models"
)

func Start() {
	db.GetDb()
	models.MigrateAll()
	router := gin.Default()
	addStudentsRoutes(router)
	addClassesRoutes(router)
	addTeachersRoutes(router)

	router.Run(":8000")
}
