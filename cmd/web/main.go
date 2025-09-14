package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/midoon/kamipa_backend/internal/configs"
)

func main() {
	cnf := configs.GetConfig()

	kamipaDB := configs.KamipaNewDatabase(cnf)
	simipaDB := configs.SimipaNewDatabase(cnf)
	validate := configs.NewValidator()

	r := mux.NewRouter()

	configs.BootStrap(&configs.BootstrapConfig{
		KamipaDB: kamipaDB,
		SimipaDB: simipaDB,
		Router:   r,
		Validate: validate,
	})

	addr := fmt.Sprintf("%s:%s", cnf.Server.Host, cnf.Server.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	fmt.Println("Server is running on port", addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
