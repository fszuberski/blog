package service

import (
	"strings"

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
	p, err := s.pokemonDataPort.GetPokemon(id)
	if err != nil {
		return nil, err
	}

	p.Name = strings.ToUpper(p.Name)

	return p, nil
}

func (s pokemonService) ListPokemon() ([]*domain.Pokemon, error) {
	pl, err := s.pokemonDataPort.ListPokemon()
	if err != nil {
		return nil, err
	}

	for _, p := range pl {
		p.Name = strings.ToUpper(p.Name)
	}

	return pl, nil
}
