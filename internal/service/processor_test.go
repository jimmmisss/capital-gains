package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessInput_ValidInput(t *testing.T) {
	processor := NewProcessor()

	input := `[{"operation":"buy","unit-cost":10.0,"quantity":100}]`
	result, err := processor.ProcessInput([]byte(input))

	assert.NoError(t, err)
	assert.Contains(t, result, `"tax":0.00`)
}

func TestProcessInput_InvalidJSON(t *testing.T) {
	processor := NewProcessor()

	input := `{"operation":"buy","unit-cost":10.0,"quantity":100}`
	_, err := processor.ProcessInput([]byte(input))

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erro ao decodificar JSON")
}
