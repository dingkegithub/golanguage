package war3

import "fmt"

type War3 struct {
	race WarRace

	factory War3Factory
}

func NewWar3(race WarRace) *War3 {
	war := &War3{}
	war.race = race

	if race == ORC {
		war.factory = &OrcFactory{}
	} else {
		war.factory = &HumFactory{}
	}

	return war
}

func (w *War3) Start() {
	fmt.Println("Race:", w.race.String(), "starting...")
	w.factory.CreateHall()
	w.factory.CreateStore()
	fmt.Println("Race:", w.race.String(), "started...")
}
