package domain

type OperacaoInterface interface {
	Processar(c *CalculadoraImpostos, op Operacao) float64
}
