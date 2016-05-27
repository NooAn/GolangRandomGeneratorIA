package main

import (
	"fmt"
)

const N = 2
const D = 3

func size() uint32 {
	return 1 << N
}
func modK(x uint32) uint32 {
	return x & (K() - 1)
}
func K() uint32 {
	return 1 << (2*N + D)
}

func mod(x uint32) uint32 { // operation (mod 2^k), where 2^k = size()
	return x & (size() - 1)
}

func checkZeroAllNumber(mas []uint32, res []uint32) {
	count := 0
	for i := range mas {
		if mas[i] == 0 {
			count += 1
		}
	}
	if count == len(mas) {
		fmt.Println(res)
	}
}

func IA(mas []uint32, q uint32) {
	r := make([]uint32, size())
	var x uint32
	masRes := make([]uint32, len(mas))
	copy(masRes, mas)
	b := q
	var i uint32
	for i = 0; i < size(); i += 1 {
		x = mas[i]
		masRes[i] = modK(masRes[mod(x)] + b)
		r[i] = modK(masRes[mod(masRes[i]>>N)] + x)
		b = r[i]
	}
	checkZeroAllNumber(r, mas)
}
func gen(mas []uint32, pos uint32) {
	if pos == size() {
		IA(mas, 0)
		return
	}
	for i := 0; i < int(K()); i += 1 {
		mas[pos] = uint32(i)
		gen(mas, pos+1)
	}
}
func gen2(mas2 []uint32, pos uint32) {
	if pos == size() {
		IA(mas2, 0)
		return
	}
	for i := int(K()) / 2; i < int(K()); i += 1 {
		mas2[pos] = uint32(i)
		gen2(mas2, pos+1)
	}
}
func main() {
	fmt.Println("start")
	r := make([]uint32, size())

	go gen(r, 0)
	//go gen2(r2, 0)
}
