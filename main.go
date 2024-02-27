package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func Customer() chi.Router {
	route := chi.NewRouter()

	return route
}

func Invoice() chi.Router {
	route := chi.NewRouter()

	return route
}

func main() {
	router := chi.NewRouter()

	router.Mount("/customer/", Customer())
	router.Mount("/invoice/", Invoice())

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
