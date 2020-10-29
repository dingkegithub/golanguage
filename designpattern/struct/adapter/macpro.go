package adapter

import (
	"fmt"
	"strings"
)

type MacPro struct {
}

func (mp *MacPro) Display(s SourceInterface) {
	msg := s.Output()
	if strings.Contains(msg, "typec:") {
		fmt.Println("MacPro-UHD: ", msg)
	} else {
		fmt.Println("Unsupport Signal")
	}
}
