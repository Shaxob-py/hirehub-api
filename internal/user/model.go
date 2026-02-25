package user

import "time"

type ModelUser struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Role      string    `db:"role"`
	about     string    `db:"about"`
	Skill     string    `db:"skill"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ModelUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	About string `json:"about"`
	Role  string `json:"role"`
	Email string `json:"email"`
	Skill string `json:"skill"`
}

type ModelCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	About    string `json:"about"`
	Skill    string `json:"skill"`
	Password string `json:"password"`
}

type ModelUpdateUser struct {
	ID    string  `db:"id"`
	Name  *string `json:"name"`
	Skill *string `json:"skill"`
	About *string `json:"about"`
}
