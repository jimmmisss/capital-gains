package service

import (
	"encoding/json"
	"errors"

	"github.com/jimmmisss/nubank/capital-gains/internal/domain"
)

// Processor encapsula a lógica de processamento de entrada.
type Processor struct {
	calculadora *domain.CalculadoraImpostos
}

// NewProcessor cria uma nova instância de Processor.
func NewProcessor() *Processor {
	return &Processor{
		calculadora: domain.NewCalculadoraImpostos(),
	}
}

// ProcessInput processa a entrada JSON e retorna os resultados.
func (p *Processor) ProcessInput(input []byte) (string, error) {
	var operacoes []domain.Operacao
	if err := json.Unmarshal(input, &operacoes); err != nil {
		return "", errors.New("erro ao decodificar JSON")
	}

	resultados := p.calculadora.ProcessarOperacoes(operacoes)

	output, err := json.Marshal(resultados)
	if err != nil {
		return "", errors.New("erro ao gerar saída JSON")
	}

	return string(output), nil
}
