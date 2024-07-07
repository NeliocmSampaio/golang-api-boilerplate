package services

import (
	"fmt"
	"gin-framework-test/basic-api/domain/entities"
	"gin-framework-test/basic-api/domain/repositories"
)

type BookService interface {
	Save(book entities.Book) error
	GetBooks() (books []entities.Book, err error)
}

type bookService struct {
	db repositories.BookRepository
}

func NewBookService(db repositories.BookRepository) BookService {
	return &bookService{
		db: db,
	}
}

func (s *bookService) Save(book entities.Book) (err error) {

	err = s.db.Add(book)
	if err != nil {
		return err
	}

	fmt.Println("Book Created!")
	return nil
}

func (s *bookService) GetBooks() (books []entities.Book, err error) {

	books, err = s.db.GetBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}
