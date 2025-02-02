package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type config struct {
	Login struct {
		User     string
		password string
	}
}


func main() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatalf("Error opening file: -%s\n", err)
	}

	ctg := &config{}
	dec := toml.NewDecoder(file)
	if err := dec.Decode(ctg); err != nil {
		log.Fatalf("Error can't decoder file: -%s\n", err)
	}

	fmt.Printf("User: %+v\n", ctg)
}