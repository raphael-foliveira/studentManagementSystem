package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/models"
)

func CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.BindJSON(&teacher); err != nil {
		c.JSON(400, err)
		return
	}
	teacher.Create()
	c.JSON(201, teacher)
}

func ListTeachers(c *gin.Context) {
	teacher := models.Teacher{}
	teachers := teacher.All()
	c.JSON(200, teachers)
}

func RetrieveTeacher(c *gin.Context) {
	teacherId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, "invalid id")
		return
	}
	teacher := models.Teacher{}
	teacher.Find(uint(teacherId))
	c.JSON(200, teacher)
}

func DeleteTeacher(c *gin.Context) {
	teacherId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, "invalid id")
	}
	teacher := models.Teacher{}
	teacher.Delete(uint(teacherId))
	c.JSON(204, nil)
}
