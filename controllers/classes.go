package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/models"
)

func ListClasses(c *gin.Context) {
	class := models.Class{}
	classes := class.All()
	c.JSON(http.StatusOK, classes)
}

func RetrieveClass(c *gin.Context) {
	classId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	class := models.Class{}
	class.Find(uint(classId))
	c.JSON(http.StatusOK, class)
}

func CreateClass(c *gin.Context) {
	var newClass models.Class
	if err := c.BindJSON(&newClass); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	newClass.Create()
	c.JSON(http.StatusCreated, newClass)
}

func DeleteClass(c *gin.Context) {
	classId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	class := models.Class{}
	class.Delete(uint(classId))
	c.JSON(http.StatusNoContent, nil)
}
