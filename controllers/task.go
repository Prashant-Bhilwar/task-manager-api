package controllers

import (
	"net/http"
	"task-manager/config"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	username, _ := c.Get("user") //from JWT middleware

	//Get the user
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found",
		})
		return
	}
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	//Create the task
	task := models.Task{
		Title:     input.Title,
		Completed: false,
		UserID:    user.ID,
	}
	config.DB.Create(&task)
	c.JSON(http.StatusCreated, gin.H{
		"messsage": "Task created successfully",
		"task":     task,
	})
}

func GetTask(c *gin.Context) {
	username, _ := c.Get("user")

	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found",
		})
		return
	}
	var tasks []models.Task
	config.DB.Where("user_id = ?", user.ID).Find(&tasks)

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}
