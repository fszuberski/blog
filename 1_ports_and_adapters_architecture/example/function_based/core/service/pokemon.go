package service

import (
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/function_based/adapters"
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/function_based/core/domain"
)

func GetPokemon(id int) (*domain.Pokemon, error) {
	return adapters.GetPokemonPort().GetPokemon(id)
}

func ListPokemon() ([]*domain.Pokemon, error) {
	return adapters.ListPokemonPort().ListPokemon()
}
