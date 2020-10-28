package war3

import "fmt"

type WarRace int

const (
	ORC = 1 << iota
	HUM
)

func (r WarRace) String() string {
	switch r {
	case ORC:
		return "ORC"
	case HUM:
		return "HUM"
	default:
		return "Unknown"
	}
}

type War3Factory interface {
	CreateStore() Store

	CreateHall() Hall
}

type OrcFactory struct {
}

func (orc *OrcFactory) CreateStore() Store {
	fmt.Println("Orc Store Create")
	return &OrcStore{}
}

func (orc *OrcFactory) CreateHall() Hall {
	fmt.Println("Orc Hall Create")
	return &OrcHall{}
}

type HumFactory struct {
}

func (hum *HumFactory) CreateStore() Store {
	fmt.Println("Hum Store Create")
	return &HumStore{}
}

func (hum *HumFactory) CreateHall() Hall {
	fmt.Println("Hum Hall Create")
	return &HumHall{}
}
