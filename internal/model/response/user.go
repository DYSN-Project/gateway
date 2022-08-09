package response

import (
	"time"
)

type User struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	Status    int        `json:"status"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func NewUserIngot() *User {
	return &User{}
}
