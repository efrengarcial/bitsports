package user

import (
	"errors"
	"time"
)

var (
	// ErrAuthenticationFailure occurs when a user attempts to authenticate but
	// anything goes wrong.
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// User represents an individual user.
type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"` //// CreatedAt holds the value of the "created_at" field.
	UpdatedAt    time.Time `json:"updatedAt"` // UpdatedAt holds the value of the "updated_at" field.
}

// NewUser contains information needed to create a new User.
type NewUser struct {
	Name            string `json:"name" validate:"required,lte=100"`
	Email           string `json:"email" validate:"required,email,lte=100"`
	Password        string `json:"password" validate:"required,lte=100"`
	PasswordConfirm string `json:"password_confirm" validate:"eqfield=Password"`
	PasswordHash    []byte `json:"-"`
}

// Token Entity
type Token struct {
	Token string `json:"token"`
	Valid bool   `json:"valid"`
}
