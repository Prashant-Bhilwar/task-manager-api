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

func UpdateTask(c *gin.Context) {
	username, _ := c.Get("user")

	// Find thr user
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	//get task ID from URL param
	taskID := c.Param("id")

	//Find the task by ID and user
	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", taskID, user.ID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Inavlid input",
		})
		return
	}

	//Update fields
	task.Title = input.Title
	task.Completed = input.Completed

	config.DB.Save(&task)

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"task":    task,
	})
}
