// handlers/user.go
package handlers

import (
	"go-orm-template/auth"
	"go-orm-template/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var user *models.UserWithPassword
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInDB, err := models.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username"})
		return
	}

	if !auth.VerifyPassword(user.Password, userInDB.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateJWT(user.Username, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    userInDB,
	})
}
