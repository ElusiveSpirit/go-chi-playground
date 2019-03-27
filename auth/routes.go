package auth

import (
	"github.com/go-chi/chi"
)

var controller = &Controller{Repository: Repository{}}

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/get-token", controller.GetToken)

	return r
}
