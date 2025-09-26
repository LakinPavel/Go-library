package usecases_mock

import (
	"context"
	"testing"

	"github.com/project/library/internal/entity"
	"github.com/project/library/internal/usecase/library"
	"go.uber.org/mock/gomock"
)

func AuthorMock(t *testing.T, err error) library.AuthorUseCase {
	t.Helper()

	ctrl := gomock.NewController(t)
	MyMockAuthorUseCase := NewMockAuthorUseCase(ctrl)

	MyMockAuthorUseCase.EXPECT().RegisterAuthor(gomock.Any(), gomock.Any()).Return(entity.Author{Name: "testName"}, err).AnyTimes()
	MyMockAuthorUseCase.EXPECT().GetAuthor(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, authorID string) (entity.Author, error) {
		return entity.Author{
			ID:   authorID,
			Name: "testName",
		}, err
	}).AnyTimes()
	MyMockAuthorUseCase.EXPECT().ChangeAuthorInfo(gomock.Any(), gomock.Any(), gomock.Any()).Return(err).AnyTimes()
	MyMockAuthorUseCase.EXPECT().GetAuthorBooks(gomock.Any(), gomock.Any()).Return([]entity.Book{}, err).AnyTimes()

	return MyMockAuthorUseCase
}

func IncorrectAuthorMock(t *testing.T) library.AuthorUseCase {
	t.Helper()
	return AuthorMock(t, entity.ErrAuthorNotFound)
}

func CorrectAuthorMock(t *testing.T) library.AuthorUseCase {
	t.Helper()
	return AuthorMock(t, nil)
}

func BookMock(t *testing.T, err error) library.BooksUseCase {
	t.Helper()

	ctrl := gomock.NewController(t)
	MyMockBookUseCase := NewMockBooksUseCase(ctrl)

	MyMockBookUseCase.EXPECT().RegisterBook(gomock.Any(), gomock.Any(), gomock.Any()).Return(entity.Book{}, err).AnyTimes()
	MyMockBookUseCase.EXPECT().GetBook(gomock.Any(), gomock.Any()).Return(entity.Book{}, err).AnyTimes()
	MyMockBookUseCase.EXPECT().UpdateBook(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(err).AnyTimes()

	return MyMockBookUseCase
}

func IncorrectBookMock(t *testing.T) library.BooksUseCase {
	t.Helper()
	return BookMock(t, entity.ErrBookNotFound)
}

func CorrectBookMock(t *testing.T) library.BooksUseCase {
	t.Helper()
	return BookMock(t, nil)
}
