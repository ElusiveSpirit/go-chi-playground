package auth

import (
	"awesomeProject/types"
	"github.com/go-chi/chi"
)

var controller = &Controller{Repository: Repository{}}

var routes = types.Routes{
	types.Route{
		Name:        "Authentication",
		Method:      "POST",
		Pattern:     "/get-token",
		HandlerFunc: controller.GetToken,
	},
}

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()

	for _, route := range routes {
		router.Method(route.Method, route.Pattern, route.HandlerFunc)
	}

	return router
}
