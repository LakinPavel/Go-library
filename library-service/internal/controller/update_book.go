package controller

import (
	"context"

	generated "github.com/project/library/generated/api/library"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *implementation) UpdateBook(ctx context.Context, req *generated.UpdateBookRequest) (*generated.UpdateBookResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := i.booksUseCase.UpdateBook(ctx, req.GetId(), req.GetName(), req.GetAuthorIds())
	if err != nil {
		return nil, i.convertErr(err)
	}

	return &generated.UpdateBookResponse{}, nil
}
