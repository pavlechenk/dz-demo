package storage

import (
	"encoding/json"
	"errors"
	"json-cloud-cli/bins"
	"json-cloud-cli/file"
	"os"
)

type Storage struct {
	filePath string
}

func NewStorage(dir string) *Storage {
	return &Storage{
		filePath: dir,
	}
}

func (storage Storage) GetBins() (*bins.BinList, error) {
	if !file.IsJson(storage.filePath) {
		return nil, errors.New("Файл не является JSON")
	}
	data, err := file.ReadFile(storage.filePath)
	if err != nil {
		return nil, err
	}
	var binList bins.BinList
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return nil, err
	}
	return &binList, nil
}

func (storage Storage) SaveBin(bin *bins.BinList) error {
	data, err := json.Marshal(bin)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(storage.filePath, os.O_APPEND, os.FileMode(os.O_CREATE))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(string(data))
	if err != nil {
		return err
	}
	return nil
}
