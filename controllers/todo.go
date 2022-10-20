package controllers

import (
	"net/http"
	"strconv"

	"github.com/devcode-muharik/models"
	repo "github.com/devcode-muharik/repositories"
	"github.com/gin-gonic/gin"
)

func CreatedTodo(c *gin.Context) {
	var request *models.RequestCreateTodo
	var response *models.Todo
	data := models.Activity{}
	var err error

	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": err.Error(), "data": data})
		return
	}

	response = repo.CreatedTodo(request)
	c.JSON(http.StatusCreated, response)
	return
}

func UpdateTodo(c *gin.Context) {
	var request *models.RequestUpdateTodo
	var response *models.Todo
	data := models.Activity{}
	var err error

	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": err.Error(), "data": data})
		return
	}

	response = repo.UpdateTodo(request, intVar)
	if response.ActivityGroupID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": "Activity with ID " + id + " Not Found", "data": data})
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DeleteTodo(c *gin.Context) {
	var response *models.Todo
	data := models.Activity{}
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	response = repo.DeleteTodo(intVar)
	if response.ActivityGroupID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": "Activity with ID " + id + " Not Found", "data": data})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": data})
	return
}

func ListTodoAll(c *gin.Context) {
	var response *models.ResponseActivity
	activity_group_id := c.Query("activity_group_id")
	intID, _ := strconv.Atoi(activity_group_id)
	response = repo.ListTodoAll(intID)
	c.JSON(http.StatusOK, response)
	return
}

func ListTodoDetail(c *gin.Context) {
	var response *models.ResponseActivity
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	response = repo.ListTodoDetail(intVar)
	if response.Status == "Not Found" {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}
