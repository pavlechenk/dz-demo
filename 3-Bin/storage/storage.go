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
		if errors.Is(err, os.ErrNotExist) {
			emptyList := bins.NewBinList()
			return &emptyList, nil
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

func (storage Storage) SaveBin(bin *bins.Bin) error {
	binList, err := storage.GetBins()
	if err != nil {
		return err
	}
	binList.Items = append(binList.Items, *bin)
	file, err := os.Create(storage.filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(binList)
}
