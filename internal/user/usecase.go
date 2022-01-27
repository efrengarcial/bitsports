package user

import (
	"bitsports/pkg/auth"
	"context"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UseCase struct {
	repo IRepository
}

// NewUseCase will create new a UseCase object
func NewUseCase(repo IRepository) *UseCase {
	return &UseCase{
		repo,
	}
}

// CreateUser inserts a new user into the database.
func (uc UseCase) CreateUser(ctx context.Context, nu NewUser) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	nu.PasswordHash = hash

	return uc.repo.Create(ctx, nu)
}
// Auth finds a user by their email and verifies their password
func (uc UseCase) Auth(ctx context.Context, email, password string, tkn *Token) error {
	user, err := uc.repo.QueryByEmail(ctx, email)

	if err != nil {
		return err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		return ErrAuthenticationFailure
	}
	// If we are this far the request is valid. Create some claims for the user
	// and generate their token.
	claims := auth.NewClaims(user.Email,  time.Now(), time.Hour)

	tkn.Token, err = auth.GenerateToken(claims)
	if err != nil {
		return errors.Wrap(err, "generating token")
	}
	tkn.Valid = true
	return nil
}
