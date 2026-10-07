package bin

import "time"

type Bin struct {
	Id string
	Private bool
	CreatedAt time.Time
	Name string
}

type BinList struct {
	Items []Bin
}

func NewBin(id string, name string, isPrivate bool) Bin {
	return Bin{
		Id:        id,
		Name:      name,
		Private:   isPrivate,
		CreatedAt: time.Now().UTC(),
	}
}

func NewBinList() BinList {
	return BinList{
		Items: []Bin{},
	}
}
