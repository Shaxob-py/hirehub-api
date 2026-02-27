package user

import "github.com/go-chi/chi/v5"

func UserRouter(r chi.Router, handler *UserHandler) {
	r.Route("/users", func(r chi.Router) {

		r.Get("/", handler.GetAllUsers)
		r.Post("/", handler.CreateUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetUser)
			r.Put("/", handler.UpdateUser)
			r.Delete("/", handler.DeleteUser)
		})
	})
}
