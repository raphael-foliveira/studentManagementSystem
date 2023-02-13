package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/db"
	"github.com/raphael-foliveira/studentManagementSystem/models"
)

func ListClasses(c *gin.Context) {
	var classes []models.Class
	db.Db.Model(&models.Class{}).Preload("Students").Find(&classes)
	c.JSON(http.StatusOK, classes)
}

func RetrieveClass(c *gin.Context) {
	var class models.Class
	db.Db.Model(&models.Class{}).Preload("Students").First(&class, c.Param("id"))
	if class.ID == 0 {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, class)
}

func CreateClass(c *gin.Context) {
	var newClass models.Class
	if c.BindJSON(&newClass) == nil {
		db.Db.Create(&newClass)
	}
	c.JSON(http.StatusCreated, newClass)
}

func DeleteClass(c *gin.Context) {
	classId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
	}
	deletedClass := models.Class{ID: uint(classId)}
	db.Db.Delete(&deletedClass)
	c.JSON(http.StatusOK, deletedClass)
}
