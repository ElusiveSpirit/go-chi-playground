package main

import (
	"awesomeProject/auth"
	"awesomeProject/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"time"
)

type RouterFunc func() *chi.Mux

type ModuleRoute struct {
	Pattern    string
	RouterFunc RouterFunc
}

type ModuleRoutes []ModuleRoute

var modules = ModuleRoutes{
	ModuleRoute{
		Pattern:    "/store",
		RouterFunc: store.CreateRouter,
	},
	ModuleRoute{
		Pattern:    "/auth",
		RouterFunc: auth.CreateRouter,
	},
}

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,    // Log API request calls
		middleware.Recoverer, // Recover from panics without crashing server
		middleware.Timeout(60*time.Second),
	)

	router.Route("/api/v1", func(r chi.Router) {
		for _, module := range modules {
			r.Mount(module.Pattern, module.RouterFunc())
		}
	})

	return router
}
