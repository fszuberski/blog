package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const POKEMON_API_PREFIX = REST_API_PREFIX + "/pokemon"

func (a restAdapter) pokemonApi() []func() (string, func(w http.ResponseWriter, r *http.Request)) {
	return []func() (string, func(w http.ResponseWriter, r *http.Request)){
		a.getPokemonList,
		a.getPokemon,
	}
}

func (a restAdapter) getPokemonList() (string, func(w http.ResponseWriter, r *http.Request)) {
	return POKEMON_API_PREFIX, func(w http.ResponseWriter, _ *http.Request) {
		p, err := a.PokemonService.ListPokemon()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(&p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}

func (a restAdapter) getPokemon() (string, func(w http.ResponseWriter, r *http.Request)) {
	return fmt.Sprintf("%s/", POKEMON_API_PREFIX), func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/", POKEMON_API_PREFIX))
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid :id"))
			return
		}

		p, err := a.PokemonService.GetPokemon(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(&p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}
