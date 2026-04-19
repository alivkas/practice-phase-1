package internal

import (
	"flag"
	"fmt"
)

const (
	FLAGS_SIZE     = 0
	FLAGS_CAPACITY = 2
	INFO           = `
	-json-path "your-path": путь до JSON файла 
	-parse "field to parse from JSON": спарсить указанное поле из JSON
	-version "true": получить информацию о версии утилиты
`
	VERSION = "v1.0.0"
)

type CLIFlag struct {
	Path, Field string
	Version     bool
}

func RegisterFlags() CLIFlag {
	filePath := flag.String("json-path", "", "Path to json to parse")
	keyWord := flag.String("parse", "", "Value to parse from json")
	version := flag.Bool("version", false, "utility version")
	flag.Usage = helpCommand
	flag.Parse()

	return CLIFlag{
		Path:    *filePath,
		Field:   *keyWord,
		Version: *version,
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}

func Visual(cfg CLIFlag) {
	if isFlagPassed("version") {
		fmt.Printf("JSON Parser %v", VERSION)
	} else {
		parsedData := GetParsedData(cfg.Path, cfg.Field)
		for _, data := range parsedData {
			outputTypes(data)
		}
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

func helpCommand() {
	fmt.Print(INFO)
}
