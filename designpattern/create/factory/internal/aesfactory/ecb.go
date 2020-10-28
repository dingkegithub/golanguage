package aesfactory

import (
	"fmt"
	"strings"
)

type AesEcb struct {
}

func (aes AesEcb) Encode(s string) string {
	return fmt.Sprintf("ECB:%s", s)
}

func (aes AesEcb) Decode(s string) (string, error) {
	if s == "" || len(s) == 0 {
		fmt.Println(1, s)
		return "", ErrNotEcbStr
	}

	strs := strings.Split(s, ":")
	if len(strs) != 2 {
		return "", ErrNotEcbStr
	}

	if strs[0] != "ECB" {
		return "", ErrNotEcbStr
	}

	return strs[1], nil
}
