package datasource

import (
	"context"

	domain "github.com/yukiouma/repo-maker/example/domain"
)

type Repo struct {
}

func NewRepo() domain.Repo {
	return &Repo{}
}

func (r *Repo) CreateUser(ctx context.Context, req *domain.CreateUserRequest) (*domain.CreaetUserReply, error) {
	panic("Unimplemenet")
}

func (r *Repo) RemoveUser(context.Context, *domain.RemoveUserRequest) (*domain.RemoveUserReply, error) {
	panic("Unimplemenet")
}
