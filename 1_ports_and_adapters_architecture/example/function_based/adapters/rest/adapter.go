package rest

import (
	"net/http"
)

const REST_API_PREFIX = "/api"

func Serve() {
	mux := http.NewServeMux()

	registerApi(mux, pokemonApi)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic("")
	}
}

func registerApi(mux *http.ServeMux, api map[string]func(w http.ResponseWriter, r *http.Request)) {
	for route, handler := range api {
		mux.Handle(route, http.HandlerFunc(handler))
	}
}
