package adapters

import (
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/function_based/adapters/data"
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/function_based/ports"
)

func GetPokemonPort() ports.GetPokemonPort {
	return ports.GetPokemonPortFunc(data.GetPokemon)
}

func ListPokemonPort() ports.ListPokemonPort {
	return ports.ListPokemonPortFunc(data.ListPokemon)
}
