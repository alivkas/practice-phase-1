package internal

import (
	"encoding/json"
	"errors"
)

func ParseJson() ([]any, error) {
	field := Flags[1]
	data := ReadFile()
	var raw []map[string]any

	if err := json.Unmarshal(data, &raw); err != nil {
		panic(err)
	}

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
