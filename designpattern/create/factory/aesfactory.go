package factory

import "github.com/dingkegithub/golanguage/designpattern/create/factory/internal/aesfactory"

type Aes interface {
	Encode(s string) string

	Decode(s string) (string, error)
}

//
// types
//
type AesType int

const (
	AesTypeCBC AesType = 1 << iota
	AesTypeECB
)

type AesFactory struct {
}

func (AesFactory) GetAes(t AesType) Aes {
	switch t {
	case AesTypeCBC:
		return &aesfactory.AesCbc{}

	case AesTypeECB:
		return &aesfactory.AesEcb{}

	default:
		return nil
	}
}
