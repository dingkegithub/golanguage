package aesfactory

import (
	"fmt"
	"strings"
)

type AesCbc struct {
}

func (aes AesCbc) Encode(s string) string {
	return fmt.Sprintf("CBC:%s", s)
}

func (aes AesCbc) Decode(s string) (string, error) {
	if s == "" || len(s) == 0 {
		return "", ErrNotCbcStr
	}

	strs := strings.Split(s, ":")
	if len(strs) != 2 {
		return "", ErrNotCbcStr
	}

	if strs[0] != "CBC" {
		return "", ErrNotCbcStr
	}

	return strs[1], nil
}
