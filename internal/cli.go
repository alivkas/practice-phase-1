package internal

import (
	"flag"
	"fmt"
)

const (
	FLAGS_SIZE     = 0
	FLAGS_CAPACITY = 2
)

type CLIFlag struct {
	Path  string
	Field string
}

func RegisterFlags() CLIFlag {
	filePath := flag.String("json-path", "NULL", "Path to json to parse")
	keyWord := flag.String("parse", "NULL", "Value to parse from json")
	flag.Parse()
	return CLIFlag{
		Path:  *filePath,
		Field: *keyWord,
	}
}

func Visual(cfg CLIFlag) {
	parsedData := GetParsedData(cfg.Path, cfg.Field)
	for _, data := range parsedData {
		outputTypes(data)
	}
}

func outputTypes(data any) {
	switch v := data.(type) {
	case []any:
		for _, val := range v {
			fmt.Printf("%v\n", val)
		}
	case any:
		fmt.Println(v)
	default:
		fmt.Println("Unexpected type")
	}
}
