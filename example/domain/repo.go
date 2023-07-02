package domain

import (
	"context"
)

// +repo.impl
type Repo interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*CreaetUserReply, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserReply, error)
}

type CreateUserRequest struct{}

type CreaetUserReply struct{}

type RemoveUserRequest struct{}

type RemoveUserReply struct{}
