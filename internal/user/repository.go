package user

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (u *User) GetAllUsers(role string) ([]ModelUserResponse, error) {

	var users []ModelUserResponse

	query := "SELECT name, email, skill, about FROM users WHERE role = $1"

	if err := u.db.Select(&users, query, role); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) GetUserById(userId string) (*ModelUserResponse, error) {
	var user ModelUserResponse
	query := "SELECT name, email, skill, about FROM users WHERE id = $1"
	if err := u.db.Get(&user, query, userId); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) CreateUser(createUser ModelCreateUser) (*ModelCreateUser, error) {
	var user ModelCreateUser

	query := `
		INSERT INTO users (name, email, skill, about)
		VALUES ($1, $2, $3, $4)
		RETURNING  name, email, skill, about`

	err := u.db.QueryRow(query,
		createUser.Name,
		createUser.Email,
		createUser.Skill,
		createUser.About,
	).Scan(&user.Name, &user.Email, &user.Skill, &user.About)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) UpdateUser(userUpdate ModelUpdateUser) (*ModelUpdateUser, error) {
	var user ModelUpdateUser

	query := `
UPDATE users SET name = $1, skill = $2, about = $3
WHERE id = $4`
	if _, err := u.db.Exec(query, userUpdate.Name, userUpdate.Skill, userUpdate.Skill, userUpdate.ID); err != nil {
		return nil, err
	}

	return &user, nil
}
