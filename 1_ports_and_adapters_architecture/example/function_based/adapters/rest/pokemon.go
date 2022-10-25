package rest

import (
	"encoding/json"
	"net/http"

	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/function_based/core"
)

const POKEMON_API_PREFIX = REST_API_PREFIX + "/pokemon"

var pokemonApi = map[string]func(w http.ResponseWriter, r *http.Request){
	POKEMON_API_PREFIX: getPokemonList(),
}

func getPokemonList() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		p, err := core.ListPokemonUseCase().ListPokemon()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(&p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(response)
	}
}
