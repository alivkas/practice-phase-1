package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	CHMOD = 0644
)

func WriteFile(pathInput, pathOutput, field string) {
	fileName := generateName(pathInput)
	path := pathOutput + "\\" + fileName
	data := Serialize(pathInput, field)
	if err := os.WriteFile(path, data, CHMOD); err != nil {
		fmt.Printf("FATAL: %v \n", err)
		os.Exit(1)
	}
}

func generateName(pathInput string) string {
	return strings.TrimSuffix(filepath.Base(pathInput), ".json") + "_parsed.json"
}
