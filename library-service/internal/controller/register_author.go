package controller

import (
	"context"

	generated "github.com/project/library/generated/api/library"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *implementation) RegisterAuthor(ctx context.Context, req *generated.RegisterAuthorRequest) (*generated.RegisterAuthorResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	author, err := i.authorUseCase.RegisterAuthor(ctx, req.GetName())
	if err != nil {
		return nil, i.convertErr(err)
	}

	return &generated.RegisterAuthorResponse{Id: author.ID}, nil
}
