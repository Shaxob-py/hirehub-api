package user

import "time"

type UserModel struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Role      string    `db:"role"`
	Skill     string    `db:"skill"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Skill string `json:"skill"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Skill    string `json:"skill"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Name  *string `json:"name"`
	Skill *string `json:"skill"`
}
