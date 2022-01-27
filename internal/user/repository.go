package user

import "context"

type IRepository interface {
	QueryByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, usr NewUser) (*User, error)
}

