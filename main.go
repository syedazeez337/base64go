package main

import "fmt"

// consts required for the encoder and decoder
const (
	upper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower  = "abcdefghijklmnopqrstuvwxyz"
	symbol = "0123456789+/"
)

// building a base64 structure
type Base64 struct {
	table string
}

func baseInit() Base64 {
	loopup := upper + lower + symbol
	return Base64{
		table: loopup,
	}
}

/*
// get the character at index `x`
func getChar(t Base64, i int) byte {
	return t.table[i]
}
*/
// function to a method form
func (t Base64) getChar(i int) byte {
	return t.table[i]
}

func main() {
	s := "Hi"
	bit := s[0] >> 2
	fmt.Printf("H-> %v, %[1]b\n", s[0])
	fmt.Printf("    %v, %[1]b\n", bit)
	fmt.Printf("    %v, %[1]b\n", 0b10010111 & 0b00110000)
}

// calculate the encode length
func calcEncodeLength(input []byte) int {
	var nOutput int
	if len(input) < 3 {
		nOutput = 4
		return nOutput
	}

	nOutput = divCeil(len(input), 3)
	return nOutput * 4
}

// division ceiling
func divCeil(x, y int) int {
	return (x + (y - 1)) / y
}

// Calculate decode length function
func calcDecodeLength(input []byte) int {
	var nOutput int
	if len(input) < 4 {
		nOutput = 3
		return nOutput
	}

	nOutput = divFloor(len(input), 4)
	return nOutput * 3
}

func divFloor(x, y int) int {
	if y == 0 {
		panic("division by zero")
	}

	return x / y
}