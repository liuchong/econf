package main

import (
	"fmt"

	"github.com/liuchong/econf"
)

type Config struct {
	// Public fields
	Host    string
	Port    int
	Debug   bool
	Tags    []string
	Weights []float32

	// Private fields
	password string
}

func main() {
	cfg := &Config{}

	// Use default separator (comma)
	econf.SetFields(cfg)

	// Or use custom separator
	// econf.SetFieldsWithSep(cfg, "#")

	fmt.Printf("Config loaded: %+v\n", cfg)
}
