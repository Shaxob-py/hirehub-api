package user

import (
	"errors"

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

	query := `
	INSERT INTO users (password, role, name, email, skill, about)
	VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING name, role, email, skill, about
	`

	var user ModelCreateUser

	err := u.db.QueryRow(
		query,
		createUser.Password,
		createUser.Role,
		createUser.Name,
		createUser.Email,
		createUser.Skill,
		createUser.About,
	).Scan(
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Skill,
		&user.About,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) UpdateUser(userUpdate ModelUpdateUser, id string) (*ModelUpdateUser, error) {
	var user ModelUpdateUser

	query := `
UPDATE users SET name = $1, skill = $2, about = $3
WHERE id = $4;`
	if _, err := u.db.Exec(query, userUpdate.Name, userUpdate.Skill, userUpdate.Skill, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) DeleteUser(userId string) error {
	query := `DELETE FROM users WHERE id = $1`
	if _, err := u.db.Exec(query, userId); err != nil {
		return err
	}
	return nil
}

func (u *User) CheckUserEmail(req ModelUserLogin) error {
	var user ModelUserLogin

	query := "SELECT  email, password FROM users WHERE email = $1"
	if err := u.db.Get(&user, query, req.Email); err != nil {
		return err
	}

	if !CheckPasswordHash(req.Password, user.Password) {
		return errors.New("invalid password or email")
	}
	return nil
}
