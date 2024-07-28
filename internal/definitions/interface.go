package definitions

import "context"

type UserService interface {
	CreateUser(ctx context.Context, user *CreateUserRequest) (id int64, err error)
	UpdateUser(ctx context.Context, user *UpdateUserRequest) error
	DeleteUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (*User, error)
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *CreateUserRequest) (id int64, err error)
	UpdateUser(ctx context.Context, user *UpdateUserRequest) error
	DeleteUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (*User, error)
}
