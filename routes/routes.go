package routes

import (
	"go-blockchain/controllers"
	"net/http"
)

// InitRoutes sets up the routes and returns the router.
func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Define routes and their handlers
	mux.HandleFunc("/", controllers.HelloHandler)

	return mux
}
