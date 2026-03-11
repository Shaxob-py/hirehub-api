package user

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	store *User
}

func NewUserHandler(store *User) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckUser(inp *ModelCreateUser) error {

	if inp.Skill == "" || inp.About == "" {
		return errors.New("invalid user data")
	}

	if !strings.HasSuffix(inp.Email, "@gmail.com") {
		return errors.New("invalid email")
	}

	hash, err := HashPassword(inp.Password)
	if err != nil {
		return err
	}

	inp.Password = hash

	return nil
}
