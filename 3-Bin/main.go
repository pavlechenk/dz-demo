package main

import (
	"json-cloud-cli/bins"
	"json-cloud-cli/storage"
	"time"

	"github.com/fatih/color"
)

type BinService struct {
	repo storage.BinDb
}

func NewBinService(repo storage.BinDb) *BinService {
	return &BinService{
		repo: repo,
	}
}

func (service BinService) GetBins() (*bins.BinList, error) {
	return service.repo.GetBins()
}

func (service BinService) SaveBin(bin *bins.Bin) error {
	return service.repo.SaveBin(bin)
}

func main() {
	storageCloud := storage.NewStorageCloud("https://google.com")
	binService := NewBinService(storageCloud)
	_, err := binService.GetBins()
	if err != nil {
		color.Red(err.Error())
	}
	bin := bins.Bin{
		Id: "1",
		Private: false,
		CreatedAt: time.Now(),
		Name: "MyBins",
	}
	err = binService.SaveBin(&bin)
	if err != nil {
		color.Red(err.Error())
	}

}
