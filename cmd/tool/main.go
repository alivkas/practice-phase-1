package main

import (
	"fmt"
	"log"

	"github.com/alivkas/phase-1/internal"
)

func main() {
	log.Println("Hello world")
	internal.ParseValues()
	fmt.Println(internal.ParseJson())
}
