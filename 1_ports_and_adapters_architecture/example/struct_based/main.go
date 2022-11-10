package main

import (
	// data "github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/adapters/data/mock"
	data "github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/adapters/data/pokeapi"
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/adapters/rest"
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/core/service"
)

func main() {

	pokemonDataAdapter := data.NewPokemonDataAdapter()
	pokemonService := service.NewPokemonService(pokemonDataAdapter)
	restAdapter := rest.NewRestAdapter(pokemonService)

	restAdapter.Serve()
}
