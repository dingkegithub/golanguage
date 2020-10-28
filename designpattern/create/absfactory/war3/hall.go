package war3

type Hall interface {
	UpgradeHall()

	CreateFarmer()
}

type OrcHall struct {
	Level int

	FarmerNum int
}

func (orc *OrcHall) UpgradeHall() {
	orc.Level += 1
}

func (orc *OrcHall) CreateFarmer() {
	orc.FarmerNum += 1
}

type HumHall struct {
	Level int

	FarmerNum int
}

func (hum *HumHall) UpgradeHall() {
	hum.Level += 1
}

func (hum *HumHall) CreateFarmer() {
	hum.FarmerNum += 1
}
