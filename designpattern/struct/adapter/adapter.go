// Package adapter
package adapter

import (
	"fmt"
	"strings"
)

type UgreenAdapter struct {
	source SourceInterface
}

func (u *UgreenAdapter) Output() string {
	return u.CvtSignal(u.source.Output())
}

func (u *UgreenAdapter) CvtSignal(msg string) string {
	return fmt.Sprintf("typec:%s", strings.Split(msg, ":")[1])
}
