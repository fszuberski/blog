package cli

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"strconv"

	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/ports"
)

type CliRunner interface {
	Run() error
}

type cliAdapter struct {
	PokemonService ports.PokemonService
}

func NewCliAdapter(ps ports.PokemonService) *cliAdapter {
	return &cliAdapter{ps}
}

func (ca cliAdapter) Run() error {
	args := flag.Args()
	if len(args) == 0 {
		return errors.New("no arguments passed")
	}

	cmd := args[0]
	switch cmd {
	case "list":
		return ca.list()
	case "get":
		return ca.get()
	default:
		fmt.Println(fmt.Errorf("invalid cmd: %s", cmd))
	}

	return nil
}

func (ca cliAdapter) list() error {
	pl, err := ca.PokemonService.ListPokemon()
	if err != nil {
		return err
	}

	return prettyPrint(pl)
}

func (ca cliAdapter) get() error {
	args := flag.Args()

	if len(args) < 2 {
		return errors.New("invalid amount of arguments for command")
	}

	idStr := args[1]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}

	p, err := ca.PokemonService.GetPokemon(int(id))
	if err != nil {
		return err
	}

	return prettyPrint(&p)
}

func prettyPrint(input any) error {
	j, err := json.MarshalIndent(&input, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(j))
	return nil
}
