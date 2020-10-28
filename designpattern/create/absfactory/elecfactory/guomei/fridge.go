package guomei

type GmFridge struct {
	ModelName   string
	Temperature int
}

func NewGmBridge(name string, initTemp int) *GmFridge {
	return &GmFridge{
		ModelName:   name,
		Temperature: initTemp,
	}
}

func (h *GmFridge) TempUpper() int {
	h.Temperature += 1
	return h.Temperature
}

func (h *GmFridge) TempDower() int {
	h.Temperature -= 1
	return h.Temperature
}

func (h *GmFridge) Model() string {
	return h.ModelName
}
