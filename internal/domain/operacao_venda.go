package domain

import (
	"math"
)

type Venda struct{}

func (s *Venda) Processar(c *CalculadoraImpostos, op Operacao) float64 {
	valorVenda := float64(op.Quantity) * op.UnitCost
	valorCompra := float64(op.Quantity) * c.PrecoMedioPonderado
	lucroOuPrejuizo := valorVenda - valorCompra

	if lucroOuPrejuizo < 0 {
		c.PrejuizoAcumulado += math.Abs(lucroOuPrejuizo)
		c.QuantidadeTotal -= op.Quantity
		return 0.0
	}

	if c.PrejuizoAcumulado > 0 {
		if lucroOuPrejuizo <= c.PrejuizoAcumulado {
			c.PrejuizoAcumulado -= lucroOuPrejuizo
			lucroOuPrejuizo = 0
		} else {
			lucroOuPrejuizo -= c.PrejuizoAcumulado
			c.PrejuizoAcumulado = 0
		}
	}

	c.QuantidadeTotal -= op.Quantity
	if valorVenda > c.ValorVendaIsentoImposto {
		return lucroOuPrejuizo * c.AliquotaImposto
	}

	return 0.0
}
