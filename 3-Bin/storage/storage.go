package storage

import (
	"encoding/json"
	"json-cloud-cli/bins"
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
	data, err := os.ReadFile(storage.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &bins.BinList{}, err
		}
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
	_, err = file.WriteString(string(data))
	if err != nil {
		return err
	}
	return nil
}
