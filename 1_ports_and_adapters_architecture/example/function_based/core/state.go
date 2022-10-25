package core

import (
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/function_based/core/service"
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/function_based/ports"
)

func GetPokemonUseCase() ports.GetPokemonUseCase {
	return ports.GetPokemonUseCaseFunc(service.GetPokemon)
}
func ListPokemonUseCase() ports.ListPokemonUseCase {
	return ports.ListPokemonUseCaseFunc(service.ListPokemon)
}
