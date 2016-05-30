package main

import (
	"fmt"
	"os"
	"sync"
)

const N = 2
const D = 2

var f *os.File
var wg sync.WaitGroup

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
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func checkZeroAllNumber(mas []uint32, res []uint32) {
	count := 0
	for i := range mas {
		if mas[i] == 0 {
			count += 1
		}
	}
	if count == len(mas) {
		str := fmt.Sprint(res) + "\n"
		b, _ := f.WriteString(str)
		fmt.Println(res, b)

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
func isEndMas(mas []uint32, end uint32) bool {
	count := 0
	for i := range mas {
		if mas[i] == end-1 {
			count += 1
		}
	}
	if count == len(mas) {
		return true
	}
	return false
}
func gen(mas []uint32, pos uint32) {
	if pos == size() {
		//		if isEndMas(mas, K()/3) {
		//			defer wg.Done()
		//		}
		IA(mas, 0)
		//fmt.Println(mas)
		return
	}
	for i := 0; i < int(K()/3); i += 1 {
		mas[pos] = uint32(i)
		gen(mas, pos+1)
	}
}
func gen2(mas2 []uint32, pos uint32) {
	if pos == size() {
		if isEndMas(mas2, K()) {
			defer wg.Done()
		}
		IA(mas2, 0)
		return
	}
	for i := 0; i < int(K()); i += 1 {
		mas2[pos] = uint32(i)
		gen2(mas2, pos+1)
	}
}
func main() {
	fmt.Println("start")
	//r := make([]uint32, size())
	r2 := make([]uint32, size())
	var nameFile string
	nameFile = fmt.Sprintf("key for N=%d D=%d.txt", N, D)
	f, _ = os.Create(nameFile)
	defer f.Close()
	wg.Add(1)
	fmt.Println(int(K()) / 3)
	go gen2(r2, 0)

	wg.Wait()

}
