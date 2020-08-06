package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	http.Handle("/", NewRouter())

	log.Println("Listening on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
}

type AppRouter struct {
	*mux.Router
}

func NewRouter() *AppRouter {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./web"))
	r.PathPrefix("/").Handler(fs)

	return &AppRouter{
		Router: r,
	}
}
