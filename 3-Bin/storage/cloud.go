package storage

import "json-cloud-cli/bins"

type StorageCloud struct {
	URL string
}

func NewStorageCloud(url string) *StorageCloud {
	return &StorageCloud{URL: url}
}

func (storage StorageCloud) GetBins() (*bins.BinList, error) {
	return &bins.BinList{
		Items: []bins.Bin{},
	}, nil
}

func (storage StorageCloud) SaveBin(bin *bins.Bin) error {
	return nil
}
