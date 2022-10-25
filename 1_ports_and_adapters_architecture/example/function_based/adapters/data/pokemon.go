package data

import (
	"errors"

	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/function_based/core/domain"
)

func GetPokemon(id int) (*domain.Pokemon, error) {
	p, ok := data[id]
	if !ok {
		return nil, errors.New("data: invalid id")
	}

	return &p, nil
}

func ListPokemon() ([]*domain.Pokemon, error) {
	res := make([]*domain.Pokemon, 0)
	for _, v := range data {
		v := v
		res = append(res, &v)
	}

	return res, nil
}

var data = map[int]domain.Pokemon{
	1: {
		ID:   1,
		Name: "Bulbasaur",
		Types: []string{
			"Grass",
		},
	},
	2: {
		ID:   2,
		Name: "Squirtle",
		Types: []string{
			"Water",
		},
	},
	3: {
		ID:   3,
		Name: "Charmander",
		Types: []string{
			"Fire",
		},
	},
}
