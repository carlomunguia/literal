package data

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	AuthorID        int       `json:"author_id"`
	PublicationYear int       `json:"publication_year"`
	Slug            string    `json:"slug"`
	Author          Author    `json:"author"`
	Description     string    `json:"description"`
	Genres          []Genre   `json:"genres"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Author struct {
	ID         int       `json:"id"`
	AuthorName string    `json:"author_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Book) GetAll(genreIDs ...int) ([]*Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	where := ""
	if len(genreIDs) > 0 {
		var IDs []string
		for _, x := range genreIDs {
			IDs = append(IDs, strconv.Itoa(x))
		}
		where = fmt.Sprintf("where b.id in (%s)", strings.Join(IDs, ","))
	}

	query := fmt.Sprintf(`
    select b.id, b.title, b.author_id, b.publication_year, b.slug, b.description, b.created_at, b.updated_at,
    a.id, a.author_name, a.created_at, a.updated_at,
    from books b
    left join authors a on (b.author_id = a.id)
    %s
    order by b.title`, where)

	var books []*Book

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID, &book.Title, &book.AuthorID, &book.PublicationYear, &book.Slug, &book.Description, &book.CreatedAt, &book.UpdatedAt,
			&book.Author.ID, &book.Author.AuthorName, &book.Author.CreatedAt, &book.Author.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		genres, err := b.genresForBook(book.ID)
		if err != nil {
			return nil, err
		}

		book.Genres = genres

		books = append(books, &book)
	}

	return books, nil
}

func (b *Book) GetBookById(id int) (*Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select b.id, b.title, b.author_id, b.publication_year, b.slug, b.description, b.created_at, b.updated_at,
  a.id, a.author_name, a.created_at, a.updated_at,
  from books b
  left join authors a on (b.author_id = a.id)
  where b.id = $1`

	row := db.QueryRowContext(ctx, query, id)

	var book Book

	err := row.Scan(
		&book.ID, &book.Title, &book.AuthorID, &book.PublicationYear, &book.Slug, &book.Description, &book.CreatedAt, &book.UpdatedAt,
		&book.Author.ID, &book.Author.AuthorName, &book.Author.CreatedAt, &book.Author.UpdatedAt)
	if err != nil {
		return nil, err
	}

	genres, err := b.genresForBook(id)
	if err != nil {
		return nil, err
	}

	book.Genres = genres

	return &book, nil
}

func (b *Book) GetBookBySlug(slug string) (*Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select b.id, b.title, b.author_id, b.publication_year, b.slug, b.description, b.created_at, b.updated_at,
  a.id, a.author_name, a.created_at, a.updated_at,
  from books b
  left join authors a on (b.author_id = a.id)
  where b.slug = $1`

	row := db.QueryRowContext(ctx, query, slug)

	var book Book

	err := row.Scan(
		&book.ID, &book.Title, &book.AuthorID, &book.PublicationYear, &book.Slug, &book.Description, &book.CreatedAt, &book.UpdatedAt,
		&book.Author.ID, &book.Author.AuthorName, &book.Author.CreatedAt, &book.Author.UpdatedAt)
	if err != nil {
		return nil, err
	}

	genres, err := b.genresForBook(book.ID)
	if err != nil {
		return nil, err
	}

	book.Genres = genres

	return &book, nil
}

func (b *Book) genresForBook(id int) ([]Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var genres []Genre

	query := `select id, genre_name, created_at, updated_at from genres where id in
  (select genre_id from book_genres where book_id = $1)`

	genreRows, err := db.QueryContext(ctx, query, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer genreRows.Close()

	var genre Genre
	for genreRows.Next() {
		err := genreRows.Scan(&genre.ID, &genre.GenreName, &genre.CreatedAt, &genre.UpdatedAt)
		if err != nil {
			return nil, err
		}

		genres = append(genres, genre)
	}

	return genres, nil
}
