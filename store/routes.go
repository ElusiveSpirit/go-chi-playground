package store

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
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
		"Authentication",
		"POST",
		"/get-token",
		controller.GetToken,
	},
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
		AuthenticationMiddleware(controller.AddProduct),
	},
	Route{
		"UpdateProduct",
		"PUT",
		"/products/{id}",
		AuthenticationMiddleware(controller.UpdateProduct),
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
		AuthenticationMiddleware(controller.DeleteProduct),
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
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,    // Log API request calls
		middleware.Recoverer, // Recover from panics without crashing server
	)

	for _, route := range routes {
		router.Method(route.Method, route.Pattern, route.HandlerFunc)
	}

	return router
}
