package route

import (
	"github.com/gorilla/mux"
	"github.com/midoon/kamipa_backend/internal/controller"
)

type RouteConfig struct {
	Router         *mux.Router
	UserController *controller.UserController
}

func (rc *RouteConfig) Setup() {
	rc.setupPublicRoute()
	rc.setupPrivateRoute()
}

// without middleware
func (rc *RouteConfig) setupPublicRoute() {
	rc.Router.HandleFunc("/api/auth/register", rc.UserController.Register).Methods("POST")
	rc.Router.HandleFunc("/api/auth/login", rc.UserController.Login).Methods("POST")
}

// with middleware
func (rc *RouteConfig) setupPrivateRoute() {
	// auth := rc.router.NewRoute().Subrouter()
	// auth.Use(AuthMiddleware)
	// auth.HandleFunc("/users", UsersHandler).Methods("GET")
	// auth.HandleFunc("/contacts", ContactsHandler).Methods("GET")

}
