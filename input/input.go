package input

import (
	"io"
	"os"
	"path/filepath"
)

func Read(filename string) (string, error) {
	path := filepath.Join("input", filename)
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
