package user

type UserHandler struct {
	store *User
}

func NewUserHandler(store *User) *UserHandler {
	return &UserHandler{
		store: store,
	}
}
