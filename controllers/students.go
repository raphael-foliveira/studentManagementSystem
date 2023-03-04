package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/models"
)

func ListStudents(c *gin.Context) {
	manager := models.Student{}
	var students = manager.All()
	c.JSON(http.StatusOK, students)
}

func RetrieveStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	student := models.Student{}
	student.Find(uint(studentId))
	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var newStudent models.Student
	fmt.Println("creating student")
	if err := (c.BindJSON(&newStudent)); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	newStudent.Create()
	c.JSON(http.StatusCreated, newStudent)
}

func UpdateStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	student := models.Student{}
	student.Find(uint(studentId))
	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println("about to run Update()")
	student.Update(uint(studentId))
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	student := models.Student{}
	student.Delete(uint(studentId))
	c.JSON(http.StatusNoContent, nil)
}
