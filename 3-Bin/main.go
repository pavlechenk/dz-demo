package main

import (
	"flag"
	"fmt"
	"json-cloud-cli/api"
	"json-cloud-cli/bins"
	"json-cloud-cli/storage"
)

type BinService struct {
	repo storage.BinDb
}

func NewBinService(repo storage.BinDb) *BinService {
	return &BinService{repo: repo}
}

func (s *BinService) GetBins() (*bins.BinList, error) {
	return s.repo.GetBins()
}

func (s *BinService) SaveBin(bin *bins.Bin) error {
	return s.repo.SaveBin(bin)
}

func main() {
	useLocal := flag.Bool("local", false, "Использовать локальный JSON вместо облака")
	flag.Parse()

	var db storage.BinDb
	if *useLocal {
		db = storage.NewStorage("data.json")
	} else {
		client := api.NewClient("https://jsonbin.io")
		db = storage.NewStorageCloud(client)
	}

	binService := NewBinService(db)

	newBin := bins.NewBin("1", "Test Bin", false)
	if err := binService.SaveBin(&newBin); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	binList, err := binService.GetBins()
	if err != nil {
		fmt.Println("Ошибка получения:", err)
		return
	}

	if binList != nil && len(binList.Items) > 0 {
		fmt.Println("Первый бин:", binList.Items)
	} else {
		fmt.Println("Список бинов пуст")
	}
}
