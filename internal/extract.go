package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

func Serialize(path, field string) []byte {
	parsedData := GetParsedData(path, field)
	data, err := json.MarshalIndent(parsedData, "", " ")
	if err != nil {
		fmt.Printf("FATAL: %v", err)
		os.Exit(1)
	}

	return data
}

func Deserialize(path string) []map[string]any {
	data, err := ReadFile(path)
	if err != nil {
		fmt.Printf("FATAL: %v", err)
	}
	var raw []map[string]any

	if err := json.Unmarshal(data, &raw); err != nil {
		fmt.Printf("FATAL: %v", err)
	}

	return raw
}
