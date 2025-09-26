package controller

import (
	"context"

	generated "github.com/project/library/generated/api/library"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *implementation) ChangeAuthorInfo(ctx context.Context, req *generated.ChangeAuthorInfoRequest) (*generated.ChangeAuthorInfoResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := i.authorUseCase.ChangeAuthorInfo(ctx, req.GetId(), req.GetName())
	if err != nil {
		return nil, i.convertErr(err)
	}

	return &generated.ChangeAuthorInfoResponse{}, nil
}
