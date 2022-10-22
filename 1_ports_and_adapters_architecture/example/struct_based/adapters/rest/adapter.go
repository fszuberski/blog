package rest

import (
	"net/http"

	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/ports"
)

const REST_API_PREFIX = "/api"

type restAdapter struct {
	PokemonService ports.PokemonService
}

func NewRestAdapter(ps ports.PokemonService) *restAdapter {
	return &restAdapter{ps}
}

func (a restAdapter) Serve() {
	mux := http.NewServeMux()

	registerApi(mux, a.pokemonApi())

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic("")
	}
}

func registerApi(mux *http.ServeMux, api []func() (string, func(w http.ResponseWriter, r *http.Request))) {
	for _, function := range api {
		route, handler := function()
		mux.Handle(route, http.HandlerFunc(handler))
	}
}
