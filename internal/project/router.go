package project

import "github.com/go-chi/chi/v5"

func ProjectRouter(r chi.Router, handler *ProjectHandle) {
	r.Route("/gigs", func(r chi.Router) {
		r.Get("/", handler.GetProject)
		r.Post("/", handler.CreateProject)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetProject)
			r.Put("/", handler.UpdateProject)
			r.Delete("/", handler.DeleteProject)
		})
	})

}
