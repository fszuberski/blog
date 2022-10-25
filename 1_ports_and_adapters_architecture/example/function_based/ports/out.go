package ports

import "github.com/fszuberski/blog/1_ports_and_adapters_architecture/function_based/core/domain"

type GetPokemonPort interface {
	GetPokemon(id int) (*domain.Pokemon, error)
}

type GetPokemonPortFunc func(id int) (*domain.Pokemon, error)

func (f GetPokemonPortFunc) GetPokemon(id int) (*domain.Pokemon, error) {
	return f(id)
}

type ListPokemonPort interface {
	ListPokemon() ([]*domain.Pokemon, error)
}

type ListPokemonPortFunc func() ([]*domain.Pokemon, error)

func (f ListPokemonPortFunc) ListPokemon() ([]*domain.Pokemon, error) {
	return f()
}
