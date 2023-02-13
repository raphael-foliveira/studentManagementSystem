package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/controllers"
)

func addClassesRoutes(router *gin.Engine) {
	classes := router.Group("/classes")
	classes.GET("/", controllers.ListClasses)
	classes.GET("/:id", controllers.RetrieveClass)

	classes.POST("/", controllers.CreateClass)

	classes.DELETE("/:id", controllers.DeleteClass)
}
