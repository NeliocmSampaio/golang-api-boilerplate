package repositories

import "gin-framework-test/basic-api/domain/entities"

type BookRepository interface {
	Add(entities.Book) (err error)
	GetBooks() (books []entities.Book, err error)
	DeleteBook(id int) (err error)
}
