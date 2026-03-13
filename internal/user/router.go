package user

import "github.com/go-chi/chi/v5"

func UserRouter(r chi.Router, handler *UserHandler) {
	r.Route("/api/v1/users", func(r chi.Router) {

		r.Get("/", handler.GetAllUsersHandler)
		r.Post("/", handler.CreateUserHandler)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetUserHandler)
			r.Put("/", handler.UpdateUserHandler)
			r.Delete("/", handler.DeleteUserHandler)
		})
	})

	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", handler.CreateUserHandler)
		r.Post("/login", handler.LoginUserHandler)
	})
}
