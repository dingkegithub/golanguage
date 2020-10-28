package war3

import "fmt"

type Store interface {
	BuyShoe()

	UpgradeStore()

	BuySupply()
}

type OrcStore struct {
	Level int
}

func (orc *OrcStore) UpgradeStore() {
	orc.Level += 1
}

func (orc *OrcStore) BuySupply() {
	fmt.Println("buy supply success")
}

func (orc *OrcStore) BuyShoe() {
	fmt.Println("buy shoe success")
}

// hum store
type HumStore struct {
	Level int
}

func (hum *HumStore) UpgradeStore() {
	hum.Level += 1
}

func (hum *HumStore) BuySupply() {
	fmt.Println("buy supply success")
}

func (hum *HumStore) BuyShoe() {
	fmt.Println("buy shoe success")
}
