package adapter

import (
	"fmt"
	"strings"
)

type MacPro struct {
	externalDev MacPro2019
}

func NewMacPro(dev MacPro2019) *MacPro {
	return &MacPro{
		externalDev: dev,
	}
}

func (mp *MacPro) Display() {
	msg := mp.externalDev.ReadSignal()

	if strings.Contains(msg, "typec:") {
		fmt.Println("MacPro-UHD: ", msg)
	} else {
		fmt.Println("Unsupport Signal")
	}
}
