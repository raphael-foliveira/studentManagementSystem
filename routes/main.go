package routes

import "github.com/gin-gonic/gin"

func Start() {
	router := gin.Default()
	addStudentsRoutes(router)
	addClassesRoutes(router)

	router.Run(":8000")
}
