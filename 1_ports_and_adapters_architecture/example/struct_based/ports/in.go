package ports

import "github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/core/domain"

type PokemonService interface {
	GetPokemonUseCase
	ListPokemonUseCase
}

type GetPokemonUseCase interface {
	GetPokemon(id int) (*domain.Pokemon, error)
}

type ListPokemonUseCase interface {
	ListPokemon() ([]*domain.Pokemon, error)
}
