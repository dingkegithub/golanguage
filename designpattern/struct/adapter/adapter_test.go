// Package adapter
package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	// 假设外设信号里包含大秦帝国这部影片的数据
	msg := "BT-Seed 大秦帝国 裂变"
	usbSignal := &UsbSignal{}
	usbSignal.Input(msg)

	// macpro 从 USB 直接不能读取信号播放
	// 因此先将usb 插入到绿联usb-typec 信号
	// 转换器上
	adapter := NewUgreeUsbAdapter(usbSignal)

	// 转换器插入到电脑
	myMacPro := NewMacPro(adapter)

	// 影片播放
	myMacPro.Display()
}
