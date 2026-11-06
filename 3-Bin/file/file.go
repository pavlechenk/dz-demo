package file

import (
	"errors"
	"os"
	"strings"
)

func ReadFile(filePath string) ([]byte, error) {
	isJson := IsJson(filePath)
	if !isJson {
		return []byte{}, errors.New("Файл не является JSON")
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func IsJson(filePath string) bool {
	filePath = strings.ToLower(filePath)
	return strings.HasSuffix(filePath, ".json")
}