// Package adapter
package adapter

import (
	"fmt"
	"strings"
)

type UgreenUsbAdapter struct {
	dev *UsbSignal
}

func NewUgreeUsbAdapter(dev *UsbSignal) MacPro2019 {
	return &UgreenUsbAdapter{
		dev: dev,
	}
}

func (u *UgreenUsbAdapter) ReadSignal() string {
	s := u.dev.Output()
	tp, signal := u.cvtSignal(s)
	fmt.Printf("signal cvt from %s to typec\n", tp)

	return signal
}

func (u *UgreenUsbAdapter) cvtSignal(s string) (string, string) {
	tp := strings.Split(s, ":")
	out := strings.Builder{}
	out.WriteString("typec:")
	out.WriteString(tp[1])
	return tp[0], out.String()
}
