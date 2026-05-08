package bin

import (
	"bins/config"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
	env 	config.Env
}

type List struct {
	List []Bin
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func NewBinList(list []Bin) *List {
	return &List{
		List: list,
	}
}
