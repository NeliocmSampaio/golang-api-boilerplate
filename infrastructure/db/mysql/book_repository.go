package mysql

import (
	"database/sql"
	"fmt"
	"gin-framework-test/basic-api/domain/entities"
	"gin-framework-test/basic-api/domain/repositories"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) repositories.BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) Add(book entities.Book) (err error) {
	query := `
		INSERT INTO tab_books (title, author, price)
		VALUES (?, ?, ?)
		;
	`

	result, err := r.db.Exec(query,
		book.Name,
		book.Author,
		book.Price)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	fmt.Println(id)

	fmt.Println("Book Created!")
	return nil
}

func (r *BookRepository) GetBooks() (books []entities.Book, err error) {
	query := `
		SELECT * FROM tab_books;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Price); err != nil {
			return books, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return books, err
	}

	return books, nil
}
