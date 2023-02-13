package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/db"
	"github.com/raphael-foliveira/studentManagementSystem/models"
)

func ListStudents(c *gin.Context) {
	var students []models.Student
	db.Db.Model(&models.Student{}).Preload("CurrentClasses").Find(&students)
	c.JSON(http.StatusOK, students)
}

func RetrieveStudent(c *gin.Context) {
	var student models.Student
	db.Db.Model(&models.Student{}).Preload("CurrentClasses").First(&student, c.Param("id"))
	if student.ID == 0 {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var newStudent models.Student
	if (c.BindJSON(&newStudent)) == nil {
		db.Db.Create(&newStudent)
	}
	c.JSON(http.StatusCreated, newStudent)
}

func UpdateStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
	}
	student := models.Student{ID: uint(studentId)}
	db.Db.First(&models.Student{}).Preload("CurrentClasses").First(&student)
	if (c.BindJSON(&student)) == nil {
		db.Db.Save(&student)
	}
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	db.Db.Delete(&models.Student{}, c.Param("id"))
}
