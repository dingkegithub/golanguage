// Package adapter
package adapter

import (
	"fmt"
)

type UsbSignal struct {
	msg string
}

func (us *UsbSignal) Input(msg string) {
	us.msg = fmt.Sprintf("usb:%s", msg)
}

func (us *UsbSignal) Output() string {
	defer func() {
		us.msg = ""
	}()
	return us.msg
}
