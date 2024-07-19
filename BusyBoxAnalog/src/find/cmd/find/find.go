package main

import (
	"goday02/src/find/internal/find"
	"goday02/src/find/internal/input"
)

func main() {
	cfg := input.ParseFile()
	find.Find(cfg)
}