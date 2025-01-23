package domain

import "fmt"

// Operacao representa uma operação de compra ou venda.
type Operacao struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

// Resultado encapsula o valor do imposto calculado.
type Resultado struct {
	Tax Tax `json:"tax"`
}

// Tax representa o imposto com formatação para JSON.
type Tax float64

func (t Tax) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", t)), nil
}
