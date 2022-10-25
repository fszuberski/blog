package ports

import "github.com/fszuberski/blog/1_ports_and_adapters_architecture/function_based/core/domain"

type GetPokemonUseCase interface {
	GetPokemon(id int) (*domain.Pokemon, error)
}

type GetPokemonUseCaseFunc func(id int) (*domain.Pokemon, error)

func (f GetPokemonUseCaseFunc) GetPokemon(id int) (*domain.Pokemon, error) {
	return f(id)
}

type ListPokemonUseCase interface {
	ListPokemon() ([]*domain.Pokemon, error)
}

type ListPokemonUseCaseFunc func() ([]*domain.Pokemon, error)

func (f ListPokemonUseCaseFunc) ListPokemon() ([]*domain.Pokemon, error) {
	return f()
}
