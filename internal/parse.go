package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func GetParsedData(path, field string) []any {
	raw, err := deserialize(path)
	parsedData, err := parse(raw, field)
	if err != nil {
		fmt.Printf("FATAL: %v \n", err)
		os.Exit(1)
	}

	return parsedData
}

func parse(raw []map[string]any, field string) ([]any, error) {
	parsedData := make([]any, len(raw))

	for i, m := range raw {
		val, ok := m[field]
		if ok {
			parsedData[i] = val
		} else {
			return nil, errors.New("Field not found")
		}
	}

	return parsedData, nil
}

func deserialize(path string) ([]map[string]any, error) {
	if path == "" {
		return nil, errors.New("Incorrect path")
	}
	data := ReadFile(path)
	var raw []map[string]any

	if err := json.Unmarshal(data, &raw); err != nil {
		fmt.Printf("FATAL: %v", err)
	}

	return raw, nil
}
