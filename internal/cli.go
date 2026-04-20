package internal

import (
	"flag"
	"fmt"
)

const (
	VERSION = "v1.0.0"
)

type CLIFlag struct {
	Path, Field, Output string
	Version             bool
}

func RegisterFlags() CLIFlag {
	filePath := flag.String("json-path", "", "Path to json to parse")
	keyWord := flag.String("parse", "", "Value to parse from json")
	version := flag.Bool("version", false, "utility version")
	output := flag.String("output", "", "output path")
	flag.Usage = flag.PrintDefaults
	flag.Parse()

	return CLIFlag{
		Path:    *filePath,
		Field:   *keyWord,
		Version: *version,
		Output:  *output,
	}
}

func Visual(cfg CLIFlag) {
	if cfg.Version == true {
		fmt.Printf("JSON Parser %v", VERSION)
	} else if cfg.Output != "" && cfg.Path != "" {
		parsedData := GetParsedData(cfg.Path, cfg.Field)
		for _, data := range parsedData {
			outputTypes(data)
		}
		WriteFile(cfg.Path, cfg.Output, cfg.Field)
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
