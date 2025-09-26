package controller

import (
	"context"

	generated "github.com/project/library/generated/api/library"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *implementation) GetAuthorInfo(ctx context.Context, req *generated.GetAuthorInfoRequest) (*generated.GetAuthorInfoResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	author, err := i.authorUseCase.GetAuthor(ctx, req.GetId())
	if err != nil {
		return nil, i.convertErr(err)
	}

	return &generated.GetAuthorInfoResponse{Id: author.ID, Name: author.Name}, nil
}
