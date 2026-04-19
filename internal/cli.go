package internal

import (
	"flag"
	"fmt"
)

const (
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
	flag.Usage = flag.PrintDefaults
	flag.Parse()

	return CLIFlag{
		Path:    *filePath,
		Field:   *keyWord,
		Version: *version,
	}
}

func Visual(cfg CLIFlag) {
	if cfg.Version == true {
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
	default:
		fmt.Println(v)
	}
}
