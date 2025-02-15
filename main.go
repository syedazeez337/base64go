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
	table := baseInit()
	res := table.getChar(28)
	fmt.Println(string(res))
}
