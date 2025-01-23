package infra

import (
	"bufio"
	"bytes"
	"io"
)

// ReadNextArray lê o próximo array JSON do leitor.
func ReadNextArray(reader *bufio.Reader) ([]byte, error) {
	var buffer bytes.Buffer
	inArray := false

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		buffer.WriteString(line)

		if bytes.Contains(buffer.Bytes(), []byte("[")) {
			inArray = true
		}
		if bytes.Contains(buffer.Bytes(), []byte("]")) && inArray {
			return buffer.Bytes(), nil
		}
		if err == io.EOF {
			break
		}
	}

	return nil, io.EOF
}
