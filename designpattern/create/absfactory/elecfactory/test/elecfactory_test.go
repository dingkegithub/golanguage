package absfactory

import (
	"testing"

	"github.com/dingkegithub/golanguage/designpattern/create/absfactory"
	"github.com/dingkegithub/golanguage/designpattern/create/absfactory/haier"
)

func TestElecFactory(t *testing.T) {
	var elecFactory absfactory.ElecFactory

	elecFactory = &haier.HaierFactory{}
	elecFactory.CreateTv()
	elecFactory.CreateFridge()
}
