package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID 				uuid.UUID `json:"id" validate:"omitempty,uuidv4"`
	Name 			string `json:"name" validate:"required,min=2,max=100"`
	Email 			string `json:"email" validate:"required,email"`
	Password 		string `json:"password,omitempty" validate:"required,min=8,max=255"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"updated_at"`
}

func NewUser(name, email, password string) *User {
	now := time.Now()
	return &User{
		ID: 		uuid.New(),
		Name: 		name,
		Email: 		email,
		Password: 	password,
		CreatedAt: 	now,
		UpdatedAt: 	now,
	}
}

func (u *User) Touch() {
	u.UpdatedAt = time.Now()
}