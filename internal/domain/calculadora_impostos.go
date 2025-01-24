package domain

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type CalculadoraImpostos struct {
	QuantidadeTotal         int
	PrecoMedioPonderado     float64
	PrejuizoAcumulado       float64
	ValorVendaIsentoImposto float64
	AliquotaImposto         float64
	Strategies              map[string]OperacaoInterface
}

func NewCalculadoraImpostos() *CalculadoraImpostos {
	return &CalculadoraImpostos{
		ValorVendaIsentoImposto: 20000.0,
		AliquotaImposto:         0.2,
		Strategies: map[string]OperacaoInterface{
			"buy":  &Compra{},
			"sell": &Venda{},
		},
	}
}

func (c *CalculadoraImpostos) ProcessarOperacoes(operacoes []Operacao) []Resultado {
	var resultados []Resultado

	for _, op := range operacoes {
		if err := ValidarOperacao(op); err != nil {
			log.Printf("Erro na operação: %v. Operação ignorada.\n", err)
			resultados = append(resultados, Resultado{Imposto: 0.00})
			continue
		}

		strategy, ok := c.Strategies[op.Operation]
		if !ok {
			log.Printf("Operação desconhecida: %s. Ignorada.\n", op.Operation)
			resultados = append(resultados, Resultado{Imposto: 0.00})
			continue
		}

		tax := strategy.Processar(c, op)
		resultados = append(resultados, Resultado{Imposto: Imposto(arredondar(tax))})
	}

	return resultados
}

func ValidarOperacao(op Operacao) error {
	if op.Quantity <= 0 {
		return errors.New(fmt.Sprintf("Quantidade inválida: %d", op.Quantity))
	}
	if op.UnitCost <= 0 {
		return errors.New(fmt.Sprintf("Custo unitário inválido: %.2f", op.UnitCost))
	}
	return nil
}

func arredondar(valor float64) float64 {
	return math.Round(valor*100) / 100
}
