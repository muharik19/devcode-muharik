package controllers

import (
	"net/http"
	"strconv"

	"github.com/devcode-muharik/models"
	repo "github.com/devcode-muharik/repositories"
	"github.com/gin-gonic/gin"
)

func CreatedActivity(c *gin.Context) {
	var request *models.RequestCreateActivity
	var response *models.ResponseActivity
	data := models.Activity{}
	var err error

	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": err.Error(), "data": data})
		return
	}

	response = repo.CreatedActivity(request)
	c.JSON(http.StatusCreated, response)
	return
}

func UpdateActivity(c *gin.Context) {
	var request *models.RequestUpdateActivity
	var response *models.ResponseActivity
	data := models.Activity{}
	var err error

	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": err.Error(), "data": data})
		return
	}

	response = repo.UpdateActivity(request, intVar)
	if response.Status == "Not Found" {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DeleteActivity(c *gin.Context) {
	var response *models.ResponseActivity
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	response = repo.DeleteActivity(intVar)
	if response.Status == "Not Found" {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func ListActivityAll(c *gin.Context) {
	var response *models.ResponseActivity

	response = repo.ListActivityAll()
	c.JSON(http.StatusOK, response)
	return
}

func ListActivityDetail(c *gin.Context) {
	var response *models.ResponseActivity
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	response = repo.ListActivityDetail(intVar)
	if response.Status == "Not Found" {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}
