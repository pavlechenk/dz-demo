package storage

import (
	"encoding/json"
	"errors"
	"json-cloud-cli/bins"
	"json-cloud-cli/file"
	"os"
)

type BinDb interface {
	GetBins() (*bins.BinList, error)
	SaveBin(bin *bins.Bin) error
}

type StorageJson struct {
	filePath string
}

func NewStorage(filePath string) *StorageJson {
	return &StorageJson{
		filePath: filePath,
	}
}

func (storage StorageJson) GetBins() (*bins.BinList, error) {
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

func (storage StorageJson) SaveBin(bin *bins.Bin) error {
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
