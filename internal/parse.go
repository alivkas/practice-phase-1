package internal

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrFieldNotFound = errors.New("Field not found")
)

func GetParsedData(path, field string) []map[string]any {
	parsedData, err := parse(Deserialize(path), field)
	if err != nil {
		if errors.Is(err, ErrFieldNotFound) {
			fmt.Printf("FATAL: %v \n", err)
		}

		os.Exit(1)
	}

	return parsedData
}

func parse(raw []map[string]any, field string) ([]map[string]any, error) {
	parsedData := make([]map[string]any, len(raw))

	for i, m := range raw {
		parsedJson := make(map[string]any, 1)
		val, ok := m[field]
		if ok {
			parsedJson[field] = val
			parsedData[i] = parsedJson
		} else {
			return nil, ErrFieldNotFound
		}
	}

	return parsedData, nil
}
