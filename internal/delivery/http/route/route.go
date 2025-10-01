package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/midoon/kamipa_backend/internal/controller"
	"github.com/midoon/kamipa_backend/internal/delivery/http/middleware"
	"github.com/midoon/kamipa_backend/internal/util"
)

type RouteConfig struct {
	Router         *mux.Router
	UserController *controller.UserController
	TokenUtil      *util.TokenUtil
}

func (rc *RouteConfig) Setup() {
	rc.setupPublicRoute()
	rc.setupPrivateRoute()
}

// without middleware
func (rc *RouteConfig) setupPublicRoute() {
	rc.Router.HandleFunc("/api/auth/register", rc.UserController.Register).Methods("POST")
	rc.Router.HandleFunc("/api/auth/login", rc.UserController.Login).Methods("POST")
	rc.Router.HandleFunc("/api/auth/refresh", rc.UserController.RefreshToken).Methods("POST")
}

// with middleware
func (rc *RouteConfig) setupPrivateRoute() {
	api := rc.Router.PathPrefix("/api").Subrouter()

	// inject middleware
	api.Use(func(next http.Handler) http.Handler {
		return middleware.AuthMiddleware(rc.TokenUtil, next)
	})

	api.HandleFunc("/auth/logout", rc.UserController.Logout).Methods("DELETE")

}
