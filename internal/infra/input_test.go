package infra

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadNextArray(t *testing.T) {
	input := `[{"operation":"buy","unit-cost":10.0,"quantity":100}]
	`
	reader := bufio.NewReader(strings.NewReader(input))

	result, err := ReadNextArray(reader)
	assert.NoError(t, err)
	assert.Contains(t, string(result), `"operation":"buy"`)
}

func TestReadNextArray_EOF(t *testing.T) {
	input := ``
	reader := bufio.NewReader(strings.NewReader(input))

	result, err := ReadNextArray(reader)
	assert.Error(t, err) // Espera erro de EOF
	assert.Nil(t, result)
}
