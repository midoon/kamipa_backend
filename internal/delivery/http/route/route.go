package route

import (
	"github.com/gorilla/mux"
	"github.com/midoon/kamipa_backend/internal/controller"
)

type RouteConfig struct {
	router         *mux.Router
	userController *controller.UserController
}

func (rc *RouteConfig) Setup() {
	rc.SetupPublicRoute()
	rc.SetupPrivateRoute()
}

// without middleware
func (rc *RouteConfig) SetupPublicRoute() {
	rc.router.HandleFunc("/api/auth/register", rc.userController.Register).Methods("POST")
}

// with middleware
func (rc *RouteConfig) SetupPrivateRoute() {
	// auth := rc.router.NewRoute().Subrouter()
	// auth.Use(AuthMiddleware)
	// auth.HandleFunc("/users", UsersHandler).Methods("GET")
	// auth.HandleFunc("/contacts", ContactsHandler).Methods("GET")

}
