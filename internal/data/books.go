package data

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mozillazg/go-slugify"
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

func (b *Book) Insert(book Book) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `insert into books (title, author_id, publication_year, slug, description, created_at, updated_at)
  values ($1, $2, $3, $4, $5, $6, $7) returning id`

	var id int
	err := db.QueryRowContext(ctx, stmt,
		book.Title, book.AuthorID, book.PublicationYear, slugify.Slugify(book.Title), book.Description, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (b *Book) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update books set title = $1, author_id = $2, publication_year = $3, slug = $4, description = $5, updated_at = $6, where id = $7`

	_, err := db.ExecContext(ctx, stmt,
		b.Title, b.AuthorID, b.PublicationYear, slugify.Slugify(b.Title), b.Description, time.Now(), b.ID)
	if err != nil {
		return err
	}

	if len(b.Genres) > 0 {
		stmt := `delete from book_genres where book_id = $1`
		_, err := db.ExecContext(ctx, stmt, b.ID)
		if err != nil {
			return fmt.Errorf("book updated, but genres not updated: %s", err.Error())
		}

		for _, x := range b.Genres {
			stmt := `insert into book_genres (book_id, genre_id, created_at, updated_at) values ($1, $2, $3, $4)`
			_, err := db.ExecContext(ctx, stmt, b.ID, x.ID, time.Now(), time.Now())
			if err != nil {
				return fmt.Errorf("book updated, but genres not updated: %s", err.Error())
			}
		}
	}

	return nil
}

func (b *Book) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from books where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
