package main

import (
	"fmt"
	"os"
	"github.com/pkg/errors"
)


type Config struct {
	//pass
}

func readconfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config file")
	}

	defer file.Close()

	return &Config{}, nil
}

func main(){
	config, err := readconfig("/path/to/fake/file.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(config)
}