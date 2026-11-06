package bins

import "time"

type Bin struct {
	Id 			string 		`json:"id"`
	Private 	bool		`json:"private"`
	CreatedAt 	time.Time	`json:"createdAt"`
	Name 		string		`json:"name"`
}

type BinList struct {
	Items []Bin `json:"items"`
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
