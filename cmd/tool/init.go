package main

import "github.com/alivkas/phase-1/internal"

func Init() {
	cfg := internal.RegisterFlags()
	internal.Visual(cfg)
}
