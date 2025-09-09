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

	fmt.Println("Kamipa DB connected:", kamipaDB != nil)
	fmt.Println("Simipa DB connected:", simipaDB != nil)

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
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
