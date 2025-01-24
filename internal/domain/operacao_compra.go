package domain

type Compra struct{}

func (b *Compra) Processar(c *CalculadoraImpostos, op Operacao) float64 {
	totalCompraAtual := float64(c.QuantidadeTotal) * c.PrecoMedioPonderado
	totalNovaCompra := float64(op.Quantity) * op.UnitCost
	c.QuantidadeTotal += op.Quantity
	c.PrecoMedioPonderado = (totalCompraAtual + totalNovaCompra) / float64(c.QuantidadeTotal)
	return 0.0
}
