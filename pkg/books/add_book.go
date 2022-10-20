package books

import (
	"net/http"

	"github.com/dickyrdiar/go-gin-api-tasks/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddBooksReqBody struct {
	Title       string `json: "title" binding:"required"`
	Author      string `json: "Author" binding:"required"`
	Description string `json: "Description" binding:"required"`
}

func (h handler) addBook(c *gin.Context) {
	body := AddBooksReqBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &book)
}
