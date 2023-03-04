package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/controllers"
)

func addTeachersRoutes(router *gin.Engine) {
	teachers := router.Group("/teachers")
	teachers.GET("/", controllers.ListTeachers)
	teachers.GET("/:id", controllers.RetrieveTeacher)

	teachers.POST("/", controllers.CreateTeacher)
	// teachers.PUT("/:id", controllers.UpdateTeacher)

	teachers.DELETE("/:id", controllers.DeleteTeacher)
}
