package library

import (
	"context"

	"github.com/google/uuid"
	"github.com/project/library/internal/entity"
)

func (l *libraryImpl) RegisterAuthor(ctx context.Context, authorName string) (entity.Author, error) {
	author, err := l.authorRepository.CreateAuthor(ctx, entity.Author{
		ID:   uuid.New().String(),
		Name: authorName,
	})

	if err != nil {
		return entity.Author{}, err
	}

	return author, nil
}

func (l *libraryImpl) GetAuthor(ctx context.Context, authorID string) (entity.Author, error) {
	return l.authorRepository.GetAuthor(ctx, authorID)
}

func (l *libraryImpl) GetAuthorBooks(ctx context.Context, authorID string) ([]entity.Book, error) {
	return l.authorRepository.GetAuthorBooks(ctx, authorID)
}

func (l *libraryImpl) ChangeAuthorInfo(ctx context.Context, authorID string, authorName string) error {
	return l.authorRepository.ChangeAuthorInfo(ctx, authorID, authorName)
}
