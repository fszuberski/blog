package service

import (
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/core/domain"
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/ports"
)

type pokemonService struct {
	pokemonDataPort ports.PokemonDataPort
}

func NewPokemonService(pda ports.PokemonDataPort) *pokemonService {
	return &pokemonService{pda}
}

func (s pokemonService) GetPokemon(id int) (*domain.Pokemon, error) {
	return s.pokemonDataPort.GetPokemon(id)
}

func (s pokemonService) ListPokemon() ([]*domain.Pokemon, error) {
	return s.pokemonDataPort.ListPokemon()
}
