package api

import (
	"Go/my-go-app/internal/api/handler"
	"Go/my-go-app/internal/api/middleware"
	"Go/my-go-app/pkg/logger"
	"net/http"
)

func NewRouter(userH *handler.UserHandler) http.Handler {
	mux := http.NewServeMux()

	// 1. Updated to include /api/
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK")) // Adding a body makes it easier to see in curl
	})

	// 2. This one was already correct for /api/v1/profile
	userRoutes := http.HandlerFunc(userH.GetProfile)
	mux.Handle("GET /api/v1/profile", middleware.Auth(userRoutes))

	// Wrap everything in the logger
	return logger.Middleware(mux)
}

// package api

// import (
// 	"Go/my-go-app/internal/api/handler"
// 	"Go/my-go-app/internal/api/middleware"
// 	"net/http"
// )

// func NewRouter(userH *handler.UserHandler) http.Handler {
// 	mux := http.NewServeMux()

// 	// Public Routes
// 	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	})

// 	// Protected Routes (Factor VI & VII)
// 	userRoutes := http.HandlerFunc(userH.GetProfile)
// 	mux.Handle("GET /api/v1/profile", middleware.Auth(userRoutes))

// 	return mux
// }
