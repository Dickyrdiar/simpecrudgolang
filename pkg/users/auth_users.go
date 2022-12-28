package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) session_users(c *gin.Context) {
	var logindata struct {
		Email string `json:"Eamil" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}


	if err := c.ShouldBindJSON(&logindata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := h.DB.Where("username = ?", logindata.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{ "error": "Invalid Username or Password" })
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), 
	[]byte(logindata.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{ "error": "Invalid Username or Password" })
		return
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims {
		"username": user.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}).SignedString([]byte("secret"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}