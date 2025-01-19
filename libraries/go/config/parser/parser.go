package parser

import (
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func readFileOrPanic(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return b
}

func Parse[T any](data []byte) *T {
	var config T
	if err := toml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	return &config
}

func ReadAndParse[T any](filepath string) *T {
	b := readFileOrPanic(filepath)

	return Parse[T](b)
}
