package aesfactory

import "fmt"

//
// error
//
var (
	ErrNotCbcStr error = fmt.Errorf("invalid cbc string")
	ErrNotEcbStr error = fmt.Errorf("invalid ecb string")
)
