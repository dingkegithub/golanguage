package haier

type HaierFridge struct {
	ModelName   string
	Temperature int
}

func NewHaierBridge(name string, initTemp int) *HaierFridge {
	return &HaierFridge{
		ModelName:   name,
		Temperature: initTemp,
	}
}

func (h *HaierFridge) TempUpper() int {
	h.Temperature += 1
	return h.Temperature
}

func (h *HaierFridge) TempDower() int {
	h.Temperature -= 1
	return h.Temperature
}

func (h *HaierFridge) Model() string {
	return h.ModelName
}
