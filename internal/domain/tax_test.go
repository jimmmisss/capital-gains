package domain_test

import (
	"errors"
	"testing"

	"github.com/jimmmisss/nubank/capital-gains/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestProcessarOperacoes(t *testing.T) {
	tests := []struct {
		name      string
		operacoes []domain.Operacao
		expected  []domain.Resultado
	}{
		{
			name: "Caso 1 - Operações com valores abaixo do limite isento",
			operacoes: []domain.Operacao{
				{Operation: "buy", UnitCost: 10.00, Quantity: 100},
				{Operation: "sell", UnitCost: 15.00, Quantity: 50},
				{Operation: "sell", UnitCost: 15.00, Quantity: 50},
			},
			expected: []domain.Resultado{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00},
			},
		},
		{
			name: "Caso 2 - Lucro seguido de prejuízo",
			operacoes: []domain.Operacao{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 5.00, Quantity: 5000},
			},
			expected: []domain.Resultado{
				{Tax: 0.00}, {Tax: 10000.00}, {Tax: 0.00},
			},
		},
		{
			name: "Caso 3 - Prejuízo parcialmente deduzido",
			operacoes: []domain.Operacao{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 5.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 3000},
			},
			expected: []domain.Resultado{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 1000.00},
			},
		},
		{
			name: "Caso 4 - Venda sem lucro nem prejuízo",
			operacoes: []domain.Operacao{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 10000},
			},
			expected: []domain.Resultado{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00},
			},
		},
		{
			name: "Caso 5 - Lucro após vendas sem lucro",
			operacoes: []domain.Operacao{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 5000},
			},
			expected: []domain.Resultado{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 10000.00},
			},
		},
		{
			name: "Caso 6 - Dedução completa de prejuízos",
			operacoes: []domain.Operacao{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 2.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
			},
			expected: []domain.Resultado{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 3000.00},
			},
		},
		{
			name: "Caso 7 - Prejuízo acumulado com novas compras e vendas",
			operacoes: []domain.Operacao{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 2.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
				{Operation: "buy", UnitCost: 20.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 30.00, Quantity: 4350},
				{Operation: "sell", UnitCost: 30.00, Quantity: 650},
			},
			expected: []domain.Resultado{
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 0.00}, {Tax: 3000.00},
				{Tax: 0.00}, {Tax: 0.00}, {Tax: 3700.00}, {Tax: 0.00},
			},
		},
		{
			name: "Caso 8 - Lucro sem prejuízo",
			operacoes: []domain.Operacao{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 50.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 20.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 50.00, Quantity: 10000},
			},
			expected: []domain.Resultado{
				{Tax: 0.00}, {Tax: 80000.00}, {Tax: 0.00}, {Tax: 60000.00},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calculadora := domain.NewCalculadoraImpostos()
			resultados := calculadora.ProcessarOperacoes(tt.operacoes)
			assert.Equal(t, tt.expected, resultados)
		})
	}
}

func TestValidarOperacao(t *testing.T) {
	tests := []struct {
		name     string
		operacao domain.Operacao
		expected error
	}{
		{
			name: "Operação válida",
			operacao: domain.Operacao{
				Operation: "buy",
				UnitCost:  10.00,
				Quantity:  100,
			},
			expected: nil,
		},
		{
			name: "Quantidade inválida",
			operacao: domain.Operacao{
				Operation: "buy",
				UnitCost:  10.00,
				Quantity:  0,
			},
			expected: errors.New("Quantidade inválida: 0"),
		},
		{
			name: "Custo unitário inválido",
			operacao: domain.Operacao{
				Operation: "buy",
				UnitCost:  0.00,
				Quantity:  100,
			},
			expected: errors.New("Custo unitário inválido: 0.00"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := domain.ValidarOperacao(tt.operacao)
			if tt.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expected.Error())
			}
		})
	}
}
