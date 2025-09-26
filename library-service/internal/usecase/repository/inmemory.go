package repository

import (
	"context"
	"sync"

	"github.com/project/library/internal/entity"
)

var _ AuthorRepository = (*inMemoryImpl)(nil)
var _ BooksRepository = (*inMemoryImpl)(nil)

type inMemoryImpl struct {
	authorsMx *sync.RWMutex
	authors   map[string]*entity.Author

	booksMx *sync.RWMutex
	books   map[string]*entity.Book

	authorBooksMx *sync.RWMutex
	authorBooks   map[string]map[string]struct{}
}

func NewInMemoryRepository() *inMemoryImpl {
	return &inMemoryImpl{
		authorsMx: new(sync.RWMutex),
		authors:   make(map[string]*entity.Author),

		books:   map[string]*entity.Book{},
		booksMx: new(sync.RWMutex),

		authorBooksMx: new(sync.RWMutex),
		authorBooks:   make(map[string]map[string]struct{}),
	}
}

func (i *inMemoryImpl) CreateAuthor(_ context.Context, author entity.Author) (entity.Author, error) {
	i.authorsMx.Lock()
	defer i.authorsMx.Unlock()
	i.authorBooksMx.Lock()
	defer i.authorBooksMx.Unlock()

	if _, ok := i.authors[author.ID]; ok {
		return entity.Author{}, entity.ErrAuthorAlreadyExists
	}

	i.authorBooks[author.ID] = make(map[string]struct{})
	i.authors[author.ID] = &author
	return author, nil
}

func (i *inMemoryImpl) GetAuthor(_ context.Context, authorID string) (entity.Author, error) {
	i.authorsMx.RLock()
	defer i.authorsMx.RUnlock()

	if v, ok := i.authors[authorID]; !ok {
		return entity.Author{}, entity.ErrAuthorNotFound
	} else {
		return *v, nil
	}
}

func (i *inMemoryImpl) CreateBook(_ context.Context, book entity.Book) (entity.Book, error) {
	i.booksMx.Lock()
	defer i.booksMx.Unlock()
	i.authorBooksMx.Lock()
	defer i.authorBooksMx.Unlock()
	i.authorsMx.Lock()
	defer i.authorsMx.Unlock()

	if _, ok := i.books[book.ID]; ok {
		return entity.Book{}, entity.ErrBookAlreadyExists
	}

	i.books[book.ID] = &book
	for _, authorID := range book.AuthorIDs {
		if _, ok := i.authors[authorID]; !ok {
			return entity.Book{}, entity.ErrAuthorNotFound
		}
		i.authorBooks[authorID][book.ID] = struct{}{}
	}
	return book, nil
}

func (i *inMemoryImpl) GetBook(_ context.Context, bookID string) (entity.Book, error) {
	i.booksMx.RLock()
	defer i.booksMx.RUnlock()

	if v, ok := i.books[bookID]; !ok {
		return entity.Book{}, entity.ErrBookNotFound
	} else {
		return *v, nil
	}
}

func (i *inMemoryImpl) UpdateBook(_ context.Context, bookID string, newName string, newAuthorIDs []string) error {
	i.booksMx.Lock()
	defer i.booksMx.Unlock()
	i.authorBooksMx.Lock()
	defer i.authorBooksMx.Unlock()
	i.authorsMx.Lock()
	defer i.authorsMx.Unlock()

	if previousBook, ok := i.books[bookID]; !ok {
		return entity.ErrBookNotFound
	} else {
		for _, authorID := range previousBook.AuthorIDs {
			if _, o := i.authors[authorID]; !o {
				return entity.ErrAuthorNotFound
			}
			delete(i.authorBooks[authorID], bookID)
		}

		for _, authorID := range newAuthorIDs {
			if _, o := i.authors[authorID]; !o {
				return entity.ErrAuthorNotFound
			}
			i.authorBooks[authorID][bookID] = struct{}{}
		}
		i.books[bookID] = &entity.Book{ID: bookID, Name: newName, AuthorIDs: newAuthorIDs}
		return nil
	}
}

func (i *inMemoryImpl) ChangeAuthorInfo(_ context.Context, authorID string, authorName string) error {
	i.authorsMx.Lock()
	defer i.authorsMx.Unlock()

	if _, ok := i.authors[authorID]; !ok {
		return entity.ErrAuthorNotFound
	} else {
		i.authors[authorID] = &entity.Author{Name: authorName, ID: authorID}
		return nil
	}
}

func (i *inMemoryImpl) GetAuthorBooks(ctx context.Context, authorID string) ([]entity.Book, error) {
	i.authorsMx.RLock()
	defer i.authorsMx.RUnlock()
	i.authorBooksMx.RLock()
	defer i.authorBooksMx.RUnlock()
	if _, ok := i.authors[authorID]; !ok {
		return []entity.Book{}, entity.ErrAuthorNotFound
	}
	if booksID, ok := i.authorBooks[authorID]; !ok {
		return []entity.Book{}, entity.ErrBookNotFound
	} else {
		authorBooks := make([]entity.Book, 0, len(booksID))
		for bookID := range booksID {
			book, _ := i.GetBook(ctx, bookID)
			authorBooks = append(authorBooks, book)
		}
		return authorBooks, nil
	}
}
