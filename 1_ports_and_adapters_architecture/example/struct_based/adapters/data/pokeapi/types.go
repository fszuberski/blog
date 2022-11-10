package pokeapi

import "github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/core/domain"

type Pokemon struct {
	ID    int
	Name  string
	Types []PokemonType
}

type PokemonType struct {
	Slot int
	Type Type
}

type Type struct {
	Name string
}

func (p Pokemon) toDomain() *domain.Pokemon {
	types := []string{}
	for _, t := range p.Types {
		types = append(types, t.Type.Name)
	}

	return &domain.Pokemon{
		ID:    p.ID,
		Name:  p.Name,
		Types: types,
	}
}

type PokemonListResponse struct {
	Results []PokemonListResponseItem
}

type PokemonListResponseItem struct {
	Name string
	Url  string
}
