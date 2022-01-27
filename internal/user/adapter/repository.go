package adapter

import (
	"bitsports/ent"
	entUser "bitsports/ent/user"
	"bitsports/internal/user"
	"context"
	"github.com/pkg/errors"
	"gopkg.in/jeevatkm/go-model.v1"
)

// repository is the repository of user.IRepository
type repository struct {
	client *ent.Client
}

// NewRepository will create an object that represent the user.IRepository interface
func NewRepository(client *ent.Client) *repository {
	return &repository{client}
}

func (r repository) QueryByEmail(ctx context.Context, email string) (*user.User, error) {
	q := r.client.User.
		Query().
		Where(entUser.EmailEQ(email))

	u, err := q.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, user.ErrAuthenticationFailure
		}

		return nil, errors.WithStack(err)
	}

	return toUser(u), nil
}

func (r repository) Create(ctx context.Context, usr user.NewUser) (*user.User, error) {
	u, err := r.client.
		User.
		Create().
		SetName(usr.Name).
		SetEmail(usr.Email).
		SetPasswordHash(usr.PasswordHash).
		Save(ctx)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	return toUser(u), nil
}

func toUser(dbUsr *ent.User) *user.User {
	var u user.User
	_ = model.Copy(&u, dbUsr)
	return &u
}
