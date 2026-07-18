package parser

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ParseFile[T any](path string) (*T, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var obj T

	if err := yaml.Unmarshal(data, &obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
