package repository

import (
	"context"

	"github.com/project/library/internal/entity"
)

type (
	AuthorRepository interface {
		CreateAuthor(ctx context.Context, author entity.Author) (entity.Author, error)
		ChangeAuthorInfo(ctx context.Context, authorID string, authorName string) error
		GetAuthor(ctx context.Context, id string) (entity.Author, error)
		GetAuthorBooks(ctx context.Context, id string) ([]entity.Book, error)
	}

	BooksRepository interface {
		CreateBook(ctx context.Context, book entity.Book) (entity.Book, error)
		GetBook(ctx context.Context, bookID string) (entity.Book, error)
		UpdateBook(ctx context.Context, bookID string, name string, authorIDs []string) error
	}
)
