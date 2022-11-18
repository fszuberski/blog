package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/adapters/cli"
	data_mock "github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/adapters/data/mock"
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/ports"

	data_pokeapi "github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/adapters/data/pokeapi"
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/adapters/rest"
	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/core/service"
)

func main() {
	// Parsing flags.
	env := flag.String("env", "prod", "Environment. Default: prod")
	invokeCli := flag.Bool("cli", false, "Flag indicating if application is ran in CLI mode. Default: false")
	flag.Parse()

	// Dependency injection.
	// The exact implementation of the pokemonDataAdapter depends on the env flag.
	var pokemonDataAdapter ports.PokemonDataPort
	if *env == "prod" {
		pokemonDataAdapter = data_pokeapi.NewPokemonDataAdapter()
	} else {
		pokemonDataAdapter = data_mock.NewPokemonDataAdapter()
	}

	pokemonService := service.NewPokemonService(pokemonDataAdapter)
	restAdapter := rest.NewRestAdapter(pokemonService)
	cliAdapter := cli.NewCliAdapter(pokemonService)

	// If cli invocation, run the command and exit
	if *invokeCli {
		err := cliAdapter.Run()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

		return
	}

	// Serve REST API
	restAdapter.Serve()
}
