package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var (
	ErrFieldNotFound = errors.New("Field not found")
)

func GetParsedData(path, field string) []any {
	parsedData, err := parse(deserialize(path), field)
	if err != nil {
		if errors.Is(err, ErrFieldNotFound) {
			fmt.Printf("FATAL: %v \n", err)
		}

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
			return nil, ErrFieldNotFound
		}
	}

	return parsedData, nil
}

func deserialize(path string) []map[string]any {
	data := ReadFile(path)
	var raw []map[string]any

	if err := json.Unmarshal(data, &raw); err != nil {
		fmt.Printf("FATAL: %v", err)
	}

	return raw
}
