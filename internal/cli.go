package internal

import (
	"flag"
)

const (
	FLAGS_SIZE     = 0
	FLAGS_CAPACITY = 2
)

var Flags = make([]string, FLAGS_SIZE, FLAGS_CAPACITY)

func ParseValues() {
	filePath := flag.String("json-path", "NULL", "Path to json to parse")
	keyWord := flag.String("parse", "NULL", "Value to parse from json")

	flag.Parse()

	Flags = append(Flags, *filePath, *keyWord)
}
