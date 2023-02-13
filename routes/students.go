package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/controllers"
	"github.com/raphael-foliveira/studentManagementSystem/middleware"
)

func addStudentsRoutes(router *gin.Engine) {
	students := router.Group("/students")
	students.Use(middleware.Auth())

	students.GET("/", controllers.ListStudents)
	students.GET("/:id", controllers.RetrieveStudent)

	students.POST("/", controllers.CreateStudent)
	students.PUT("/:id", controllers.UpdateStudent)

	students.DELETE("/:id", controllers.DeleteStudent)
}
