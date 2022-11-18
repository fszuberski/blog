package data

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/core/domain"
)

const POKEAPI_BASE_URL = "https://pokeapi.co/api/v2"

type pokemonDataAdapter struct {
}

func NewPokemonDataAdapter() *pokemonDataAdapter {
	return &pokemonDataAdapter{}
}

func (a pokemonDataAdapter) GetPokemon(id int) (*domain.Pokemon, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s/%d", POKEAPI_BASE_URL, "pokemon", id), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	p := Pokemon{}
	err = json.NewDecoder(res.Body).Decode(&p)
	if err != nil {
		return nil, err
	}

	return p.toDomain(), nil
}

func (a pokemonDataAdapter) ListPokemon() ([]*domain.Pokemon, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", POKEAPI_BASE_URL, "pokemon"), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	plr := PokemonListResponse{}
	err = json.NewDecoder(res.Body).Decode(&plr)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	pl := []*domain.Pokemon{}
	for _, item := range plr.Results {
		req, err := http.NewRequest(http.MethodGet, item.Url, nil)
		if err != nil {
			return nil, err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		p := Pokemon{}
		err = json.NewDecoder(res.Body).Decode(&p)
		if err != nil {
			return nil, err
		}
		res.Body.Close()

		pl = append(pl, p.toDomain())
	}

	return pl, nil
}
