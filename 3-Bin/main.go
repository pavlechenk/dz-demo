package main

import (
	"json-cloud-cli/storage"
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

func main() {
	storageCloud := storage.NewStorageCloud("https://google.com")
	binService := NewBinService(storageCloud)
	binList, err := binService.repo.GetBins()
	if err != nil {
		color.Red(err.Error())
	}
	binService.repo.SaveBin(&binList.Items[0])
}
