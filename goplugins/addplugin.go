package main

import "fmt"

var A, B int

type Mm struct {
	C int
	D int
}

func NewMm(*Mm) *Mm {
	return &Mm{
		C: 1,
		D: 2,
	}
}

func Add() {
	fmt.Println(A, "+", B, "=", A+B)
}
