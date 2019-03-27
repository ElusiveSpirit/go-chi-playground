package store

import (
	"awesomeProject/auth"
	"github.com/go-chi/chi"
)

var controller = &Controller{Repository: Repository{}}

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Get("/", controller.Index)
		r.With(auth.AuthenticationMiddleware).Post("/", controller.AddProduct)

		r.Route("/{productID}", func(r chi.Router) {
			r.Use(controller.ProductCtx)
			r.Get("/", controller.GetProduct)
			r.With(auth.AuthenticationMiddleware).Put("/", controller.UpdateProduct)
			r.With(auth.AuthenticationMiddleware).Delete("/", controller.AddProduct)
		})
	})

	r.Get("/search/{query}", controller.SearchProduct)

	return r
}
