package domain

import (
	"errors"
	"fmt"
	"log"
	"math"
)

// CalculadoraImpostos gerencia o estado necessário para cálculo de impostos.
type CalculadoraImpostos struct {
	quantidadeTotal         int
	precoMedioPonderado     float64
	prejuizoAcumulado       float64
	valorVendaIsentoImposto float64
	aliquotaImposto         float64
}

// NewCalculadoraImpostos cria uma nova instância de CalculadoraImpostos.
func NewCalculadoraImpostos() *CalculadoraImpostos {
	return &CalculadoraImpostos{
		valorVendaIsentoImposto: 20000.0,
		aliquotaImposto:         0.2,
	}
}

// ProcessarOperacoes calcula os impostos para uma lista de operações.
func (c *CalculadoraImpostos) ProcessarOperacoes(operacoes []Operacao) []Resultado {
	var resultados []Resultado

	for _, op := range operacoes {
		// Validação de entrada
		if err := ValidarOperacao(op); err != nil {
			log.Printf("Erro na operação: %v. Operação ignorada.\n", err)
			resultados = append(resultados, Resultado{Tax: 0.00})
			continue
		}

		switch op.Operation {
		case "buy":
			c.processarCompra(op)
			resultados = append(resultados, Resultado{Tax: 0.00})
		case "sell":
			tax := c.processarVenda(op)
			resultados = append(resultados, Resultado{Tax: Tax(arredondar(tax))})
		default:
			log.Printf("Operação desconhecida: %s. Ignorada.\n", op.Operation)
			resultados = append(resultados, Resultado{Tax: 0.00})
		}
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

func (c *CalculadoraImpostos) processarCompra(op Operacao) {
	totalCompraAtual := float64(c.quantidadeTotal) * c.precoMedioPonderado
	totalNovaCompra := float64(op.Quantity) * op.UnitCost
	c.quantidadeTotal += op.Quantity
	c.precoMedioPonderado = (totalCompraAtual + totalNovaCompra) / float64(c.quantidadeTotal)
}

func (c *CalculadoraImpostos) processarVenda(op Operacao) float64 {
	valorVenda := float64(op.Quantity) * op.UnitCost
	valorCompra := float64(op.Quantity) * c.precoMedioPonderado
	lucroOuPrejuizo := valorVenda - valorCompra

	if lucroOuPrejuizo < 0 {
		c.prejuizoAcumulado += math.Abs(lucroOuPrejuizo)
		c.quantidadeTotal -= op.Quantity
		return 0.0
	}

	if c.prejuizoAcumulado > 0 {
		if lucroOuPrejuizo <= c.prejuizoAcumulado {
			c.prejuizoAcumulado -= lucroOuPrejuizo
			lucroOuPrejuizo = 0
		} else {
			lucroOuPrejuizo -= c.prejuizoAcumulado
			c.prejuizoAcumulado = 0
		}
	}

	c.quantidadeTotal -= op.Quantity
	if valorVenda > c.valorVendaIsentoImposto {
		return lucroOuPrejuizo * c.aliquotaImposto
	}

	return 0.0
}

func arredondar(valor float64) float64 {
	return math.Round(valor*100) / 100
}
