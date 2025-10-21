package c

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/goccy/go-yaml"
)

func loadTestData[T any](t *testing.T) []T {
	testName := t.Name()
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		t.Fatalf("Failed to get caller information")
	}
	testDir := filepath.Dir(filename)
	yamlPath := filepath.Join(testDir, "testdata", fmt.Sprintf("%s.yaml", testName))

	// Open the YAML file
	file, err := os.Open(yamlPath)
	if err != nil {
		t.Fatalf("Failed to open test data file %s: %v", yamlPath, err)
	}
	defer file.Close()

	// Parse the YAML file
	var documents []T
	decoder := yaml.NewDecoder(file, yaml.Strict())
	for {
		var document T
		err := decoder.Decode(&document)
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatalf("Failed to parse test data file %s: %v", yamlPath, err)
		}
		documents = append(documents, document)
	}

	return documents
}
