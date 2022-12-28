package users

import (
	"net/http"

	"github.com/dickyrdiar/go-gin-api-tasks/pkg/common/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json: "email`
	Password string `json: "password"`
}

func (h handler) Register(c *gin.Context) {
	var user models.User
	// reqBody := user{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
			return
		}


	user.Password = string(hashedPassword)
	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Error creating success"})
		return
	}

	c.JSON(http.StatusOK, gin.H{ "message": "User already created" })
}
