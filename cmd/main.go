package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jimmmisss/nubank/capital-gains/internal/infra"
	"github.com/jimmmisss/nubank/capital-gains/internal/service"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	processor := service.NewProcessor()

	for {
		array, err := infra.ReadNextArray(reader)
		if err != nil {
			break
		}

		resultados, err := processor.ProcessInput(array)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
			continue
		}

		fmt.Println(resultados)
	}
}
