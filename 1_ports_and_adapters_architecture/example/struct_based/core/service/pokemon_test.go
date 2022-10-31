package service

import (
	"errors"
	"testing"

	"github.com/fszuberski/blog/1_ports_and_adapters_architecture/struct_based/core/domain"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetPokemon(t *testing.T) {
	c := &domain.Pokemon{
		ID:   4,
		Name: "Charmander",
		Types: []string{
			"Fire",
		},
	}

	testErr := errors.New("")

	testCases := []struct {
		name          string
		dataPortCalls func(*PokemonDataPortMock) *PokemonDataPortMock
		input         int
		expected      *domain.Pokemon
		expectedErr   error
	}{
		{
			name: "return Pokemon with capitalised name given valid id",
			dataPortCalls: func(p *PokemonDataPortMock) *PokemonDataPortMock {
				p.On("GetPokemon", 4).Return(c, nil)
				return p
			},
			input: 4,
			expected: &domain.Pokemon{
				ID:   4,
				Name: "CHARMANDER",
				Types: []string{
					"Fire",
				},
			},
		},
		{
			name: "return zero value Pokemon given invalid id",
			dataPortCalls: func(p *PokemonDataPortMock) *PokemonDataPortMock {
				p.On("GetPokemon", 4).Return(&domain.Pokemon{}, nil)
				return p
			},
			input:    4,
			expected: &domain.Pokemon{},
		},
		{
			name: "return error given pokemonDataPort error",
			dataPortCalls: func(p *PokemonDataPortMock) *PokemonDataPortMock {
				p.On("GetPokemon", 4).Return(&domain.Pokemon{}, testErr)
				return p
			},
			input:       4,
			expectedErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dataPort := tc.dataPortCalls(new(PokemonDataPortMock))
			service := NewPokemonService(dataPort)

			p, err := service.GetPokemon(tc.input)

			if tc.expectedErr != nil {
				require.Nil(t, p)
				require.NotNil(t, err)
				return
			}

			require.Equal(t, tc.expected, p)
			require.Nil(t, err)
		})
	}
}

type PokemonDataPortMock struct {
	mock.Mock
}

func (p PokemonDataPortMock) GetPokemon(id int) (*domain.Pokemon, error) {
	args := p.Called(id)
	return args.Get(0).(*domain.Pokemon), args.Error(1)
}

func (p PokemonDataPortMock) ListPokemon() ([]*domain.Pokemon, error) {
	return nil, nil
}
