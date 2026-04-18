package internal

import "strings"

func ParseJson() []string {
	field := Flags[1]
	splittedJson := strings.Split(ReadFile(), ",")
	parsedField := make([]string, 0)

	for _, data := range splittedJson {
		if strings.Contains(data, field) {
			parsedField = append(parsedField, data)
		}
	}

	return parsedField
}
