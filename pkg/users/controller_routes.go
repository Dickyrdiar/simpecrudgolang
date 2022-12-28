package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler {
		DB: db,
	}

	routes := r.Group("/users")
	routes.POST("/users", h.Register)
	routes.POST("/sessions", h.session_users)
}