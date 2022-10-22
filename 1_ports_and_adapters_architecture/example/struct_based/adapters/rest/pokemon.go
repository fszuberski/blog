package rest

import (
	"encoding/json"
	"net/http"
)

const POKEMON_API_PREFIX = REST_API_PREFIX + "/pokemon"

func (a restAdapter) pokemonApi() []func() (string, func(w http.ResponseWriter, r *http.Request)) {
	return []func() (string, func(w http.ResponseWriter, r *http.Request)){
		a.getPokemonList,
	}
}

func (a restAdapter) getPokemonList() (string, func(w http.ResponseWriter, r *http.Request)) {
	return POKEMON_API_PREFIX, func(w http.ResponseWriter, _ *http.Request) {
		p, err := a.PokemonService.ListPokemon()

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
