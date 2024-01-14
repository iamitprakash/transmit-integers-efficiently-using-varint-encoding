package main

import "fmt"

func main() {
//	fmt.Println("Hello")
}

var buf [11]byte

var byteShifts []uint8 = []uint8{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 1}
var BITMASK = []byte{
	0b00000001,
	0b00000011,
	0b00000111,
	0b00001111,
	0b00011111,
	0b00111111,
	0b01111111,
	0b11111111,
}

func EncodingUInt64(x uint64) []byte {
	var i int = 0

	for i = 0; i < len(byteShifts); i++ {
		buf[i] = getLSB(byte(x), byteShifts[i]) | 0b10000000

		x = x >> byteShifts[i]

		if x == 0 {
			break
		}

	}
	buf[i] = buf[i] & 0b01111111

	return append(make([]byte, 0, i+1), (buf[:i+1])...)
}

func getLSB(x byte, n uint8) byte {
	if n > 8 {
		panic(" can extract at max 8 bits from the number")
	}

	return byte(x) & BITMASK[n-1]
}

func DecodingUInt64(vint []byte) uint64 {
	var i int = 0
	var v uint64 = 0

	for i = 0; i < len(vint); i++ {
		b := getLSB(vint[i], 7)
		v = v | uint64(b)<<(7*i)

	}

	return v
}
