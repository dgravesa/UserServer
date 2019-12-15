package controller

import "net/http"

// RegisterRoutes registers all routes for the user server.
func RegisterRoutes() {
	http.HandleFunc("/users", userHandler)
}
