package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const N = 2
const D = 2

var f *os.File
var wg sync.WaitGroup

func size() int {
	return 1 << N
}
func modK(x int) int {
	return x & (K() - 1)
}
func K() int {
	return 1 << (2*N + D)
}

func mod(x int) int { // operation (mod 2^k), where 2^k = size()
	return x & (size() - 1)
}
func check2(e error) {
	if e != nil {
		panic(e)
	}
}
func checkStateOne(mas []int, res []int) bool{
	for i:= range mas {
		if (mas[i] != res[i]) {
			return false 
		}
	}
	return true
}
func check(mas []int ) bool{
	for i:=range mas {
		if mas[i] != mas[mod(mas[i])] {
			return false
		}
	}
	return true
}
func checkZeroAllNumber(mas []int, res []int) {
	
	count := 0
	for i := range mas {
		if mas[i] == 0 {
			count += 1
		}
	}
	if count == len(mas)  {
			str := fmt.Sprint(res) + "\n"
			f.WriteString(str)
			//	fmt.Println(res, b)
	}
}
/**
 Where r[i] - выходная последовательность
*/
func IA(mas []int, q int) {
	r := make([]int, size())
	var x int
	masRes := make([]int, len(mas))
	copy(masRes, mas)
	b := q
	var i int
	for i = 0; i < size(); i += 1 {
		x = mas[i]
		masRes[i] = modK(masRes[mod(x)] + b)
		r[i] = modK(masRes[mod(masRes[i]>>N)] + x)
		b = r[i]
	}

	// Проверка на неизменения состояния
	masRes2 := make([]int, len(mas))
	copy(masRes2, mas)

	// if checkStateOne(r, masRes2) {
	// 	str := fmt.Sprint(r) + "\n"
	// 	f.WriteString(str)
	// 	fmt.Println(str, b)
		
	// }
	// проверка на сплошные нули и равенство уравнений
	if check(masRes2) {
		str := fmt.Sprint(res) + "\n"
		f.WriteString(str)
		//checkZeroAllNumber(r, mas)
	}
}
func isEndMas(mas []int, end int) bool {
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
func genAll(mas2 []int, pos int) {
	if pos == size() {
		IA(mas2, 0)
		return
	}
	for i := 0; i < int(K()); i += 1 {
		mas2[pos] = int(i)
		genAll(mas2, pos+1)
	}
}
func genRandomNumber(mas []int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	var i int64
	fmt.Println("Kol:", int64(K()*K()*K()*K()*size()))
	for i = 0; i < int64(K()*K()*K()*K()*size()); i += 1 {
		for j := 0; j < size(); j += 1 {
			mas[j] = rand.Intn(int(K()))
		}
		//		fmt.Println(mas)
		IA(mas, 0)
	}
}
func main() {
	fmt.Println("start")
	r := make([]int, size())
	var nameFile string
	nameFile = fmt.Sprintf("key3 for N=%d D=%d.txt", N, D)
	f, _ = os.Create(nameFile)
	defer f.Close()
	genAll(r, 0)

	//genRandomNumber(r2)
}
