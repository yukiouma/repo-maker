# Repository Maker

Helping you to generate the implements for the repository interfaces you defined

# Installation
```bash
$ go install github.com/yukiouma/repo-maker:v0.1.1
```

# Usage
For example, we start with the project structure as follow:
```bash
./example/
├── adapter
│   └── datasource
└── domain
    └── repo.go
```
We define the repository interface in `example/domain/repo.go`
```go
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

```
Please make sure write comment `// +repo.impl` for Repository interface definition

Then generate implements with repomaker
```bash
$ repomaker --in ./example/domain --out ./example/adapter/datasource
```
Implement file `repo.impl.go` will be generated in `./example/adapter/datasource`
```bash
./example/
├── adapter
│   └── datasource
│       └── repo.impl.go
└── domain
    └── repo.go
```

Content of `repo.impl.go`
```go
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

```

After that, you need to complete the logic for the implement