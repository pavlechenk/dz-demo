package file

import (
	"os"
	"strings"
)

func ReadFile(filePath string) ([]byte, error) {
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