package storage

import (
	"errors"
	"fmt"
	"json-cloud-cli/bins"
)

type ApiClient interface {
	FetchData() (string, error)
}

type StorageCloud struct {
	client ApiClient
	items  []bins.Bin
}

var _ BinDb = (*StorageCloud)(nil)

func NewStorageCloud(client ApiClient) *StorageCloud {
	return &StorageCloud{
		client: client,
		items:  []bins.Bin{},
	}
}

func (storage *StorageCloud) GetBins() (*bins.BinList, error) {
	if _, err := storage.client.FetchData(); err != nil {
		return nil, err
	}
	return &bins.BinList{
		Items: storage.items,
	}, nil
}

func (storage *StorageCloud) SaveBin(bin *bins.Bin) error {
	if bin == nil {
		return errors.New("бин не может быть nil")
	}
	if _, err := storage.client.FetchData(); err != nil {
		return err
	}
	fmt.Printf("[Cloud] Бин сохранён: %s\n", bin.Name)
	storage.items = append(storage.items, *bin)
	return nil
}
