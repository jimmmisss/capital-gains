package domain

import (
	"fmt"
)

// Operacao representa uma operação de compra ou venda.
type Operacao struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

// Resultado encapsula o valor do imposto calculado.
type Resultado struct {
	Imposto Imposto `json:"tax"`
}

// Imposto representa o imposto com formatação para JSON.
type Imposto float64

func (t Imposto) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", t)), nil
}
