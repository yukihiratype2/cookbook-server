// Package model provides ...
package model

// User api type
type User struct {
	Email string `json:"email" validate:"required,email"`
}

// AuthParams struct
type AuthParams struct {
	User
	Password string `json:"password,omitempty" validate:"required"`
}

// CreateUserParams struct
type CreateUserParams struct {
	User
	UserName string `json:"username" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type JWTUserParams struct {
	User
	UserName string `json:"username" validate:"required"`
}
