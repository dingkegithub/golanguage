// Package adapter
package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	msg := "BT-Seed 大秦帝国 裂变"

	signal := &Usb30Source{}
	signal.Input(msg)

	macpro := &MacPro{}
	macpro.Display(signal)

	signal.Input(msg)
	signalAdapter := &UgreenAdapter{
		source: signal,
	}
	macpro.Display(signalAdapter)
}
