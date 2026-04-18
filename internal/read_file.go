package internal

import (
	"io"
	"os"
	"strings"
)

func ReadFile() []byte {
	file, err := os.Open(Flags[0])
	if err != nil {
		panic(err)
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
