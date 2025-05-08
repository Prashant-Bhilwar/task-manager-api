package controllers

import (
	"net/http"
	"task-manager/config"
	"task-manager/models"
	"task-manager/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}
	var existingUser models.User
	if err := config.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Username already taken",
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered succcessfully",
	})
}

func Login(c *gin.Context) {
	var input models.User
	var dbUser models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}
	if err := config.DB.Where("username = ?", input.Username).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Password",
		})
		return
	}
	token, err := utils.GenerateToken(dbUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Token generation failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
