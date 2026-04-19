package internal

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("FATAL: Incorrect path\n")
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	var stringData strings.Builder

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		stringData.WriteString(string(data[:n]))
	}

	return []byte(stringData.String())
}
