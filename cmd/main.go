package main

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	// Define your YAML structure here
	Key1 string `yaml:"key1"`
	Key2 int    `yaml:"key2"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <yaml-file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		os.Exit(1)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Printf("Failed to unmarshal YAML: %v\n", err)
		os.Exit(1)
	}

	output, err := yaml.Marshal(&config)
	if err != nil {
		fmt.Printf("Failed to marshal YAML: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
