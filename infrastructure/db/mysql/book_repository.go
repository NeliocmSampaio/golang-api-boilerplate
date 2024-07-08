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
		SELECT id, title, author, price, deleted
		FROM tab_books
		WHERE tab_books.deleted = 0;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Price, &book.Deleted); err != nil {
			return books, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return books, err
	}

	return books, nil
}

func (r *BookRepository) DeleteBook(id int) (err error) {
	query := `
		UPDATE tab_books
		SET tab_books.deleted = 1
		WHERE tab_books.id = ?
		;
	`

	_, err = r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
