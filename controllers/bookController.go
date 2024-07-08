package controllers

import (
	"gin-framework-test/basic-api/controllers/requests"
	"gin-framework-test/basic-api/domain/entities"
	"gin-framework-test/basic-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(bookService services.BookService) BookController {
	return BookController{
		bookService: bookService,
	}
}

func (c *BookController) HandlePostBook(ctx *gin.Context) {

	request := requests.BookRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	book := entities.Book{
		Name:   request.Name,
		Author: request.Author,
		Price:  request.Price,
	}

	err := c.bookService.Save(book)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.Status(http.StatusCreated)
}

func (c *BookController) HandleGetBooks(ctx *gin.Context) {
	books, err := c.bookService.GetBooks()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, books)
}

func (c *BookController) HandleDeleteBook(ctx *gin.Context) {
	id, success := ctx.Params.Get("id")
	if !success {
		ctx.Status(http.StatusNotFound)
	}

	bookId, err := strconv.Atoi(id)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	err = c.bookService.DeleteBook(bookId)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.Status(http.StatusOK)
}
