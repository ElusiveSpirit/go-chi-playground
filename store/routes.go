package store

import (
	"awesomeProject/auth"
	"github.com/go-chi/chi"
	"net/http"
)

var controller = &Controller{Repository: Repository{}}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/products",
		controller.Index,
	},
	Route{
		"AddProduct",
		"POST",
		"/products",
		auth.AuthenticationMiddleware(controller.AddProduct),
	},
	Route{
		"UpdateProduct",
		"PUT",
		"/products/{id}",
		auth.AuthenticationMiddleware(controller.UpdateProduct),
	},
	// Get Product by {id}
	Route{
		"GetProduct",
		"GET",
		"/products/{id}",
		controller.GetProduct,
	},
	// Delete Product by {id}
	Route{
		"DeleteProduct",
		"DELETE",
		"/products/{id}",
		auth.AuthenticationMiddleware(controller.DeleteProduct),
	},
	// Search product with string
	Route{
		"SearchProduct",
		"GET",
		"/search/{query}",
		controller.SearchProduct,
	}}

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()

	for _, route := range routes {
		router.Method(route.Method, route.Pattern, route.HandlerFunc)
	}

	return router
}
