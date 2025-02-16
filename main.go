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
	b64 := baseInit()

	input := []byte("Hello, Go!")
	encoded, err := b64.encode(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Encoded:", string(encoded))
}

// Encode function
func (b Base64) encode(input []byte) ([]byte, error) {
	if len(input) == 0 {
		return []byte(""), nil
	}

	nOut := calcEncodeLength(input)
	out := make([]byte, nOut)
	var buf [3]byte
	count := 0
	iout := 0

	for i := 0; i < len(input); i++ {
		buf[count] = input[i]
		count++
		if count == 3 {
			out[iout] = b.getChar(int(buf[0] >> 2))
			out[iout+1] = b.getChar(int(((buf[0] & 0x03) << 4) | (buf[1] >> 4)))
			out[iout+2] = b.getChar(int(((buf[1] & 0x0f) << 2) | (buf[2] >> 6)))
			out[iout+3] = b.getChar(int(buf[2] & 0x3f))
			iout += 4
			count = 0
		}
	}

	// Handle any remaining bytes (with appropriate padding).
	if count == 1 {
		// Only one byte left.
		out[iout] = b.getChar(int(buf[0] >> 2))
		out[iout+1] = b.getChar(int((buf[0] & 0x03) << 4))
		out[iout+2] = '='
		out[iout+3] = '='
	} else if count == 2 {
		// Two bytes remain.
		out[iout] = b.getChar(int(buf[0] >> 2))
		out[iout+1] = b.getChar(int(((buf[0] & 0x03) << 4) | (buf[1] >> 4)))
		out[iout+2] = b.getChar(int((buf[1] & 0x0f) << 2))
		out[iout+3] = '='
	}
	
	return out, nil
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