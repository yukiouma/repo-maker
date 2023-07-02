package domain

import (
	"context"
)

// +repo.impl
type Repo interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (CreaetUserReply, error)
	FindUser(ctx context.Context, id int) ([]*User, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserReply, error)
}

type User struct{}

type FindeUserRequest struct{}

type CreateUserRequest struct{}

type CreaetUserReply struct{}

type RemoveUserRequest struct{}

type RemoveUserReply struct{}
