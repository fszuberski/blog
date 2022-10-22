package ports

import "github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/core/domain"

type PokemonDataPort interface {
	GetPokemonPort
	ListPokemonPort
}

type GetPokemonPort interface {
	GetPokemon(id int) (*domain.Pokemon, error)
}

type ListPokemonPort interface {
	ListPokemon() ([]*domain.Pokemon, error)
}
