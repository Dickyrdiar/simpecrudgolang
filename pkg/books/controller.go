package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/books")
	routes.POST("/", h.addBook)
	routes.GET("/", h.GetBooks)
	routes.GET("/books/:id", h.GetBook)
	routes.PUT("/books/:id", h.UpdateBook)
	routes.DELETE("/books/:id", h.DeleteBook)
}
