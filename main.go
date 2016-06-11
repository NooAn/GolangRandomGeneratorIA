package main
	// q  должны повторяться
	//k and D вводим числа, выбираем с0 которые удачные
	//выходная последовательность в которую вводим s0, k, D, 

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var n, d,count uint
var f *os.File
var wg sync.WaitGroup

func size() int {
	return 1 << n
}
func modK(x int) int {
	return x & (K() - 1)
}
func K( ) int {
	return 1 << (2*n+ d)
}

func mod(x int) int { // operation (mod 2^k), where 2^k = size()
	return x & (size() - 1)
}

func checkStateOne(mas []int, res []int) bool {
	for i := range mas {
		if mas[i] != res[i] {
			return false
		}
	}
	return true
}
func check(mas []int) bool {
	for i := range mas {
		if mas[i] != mas[mod(mas[i])] {
			return false
		}
	}
	return true
}
func checkTwo(mas []int, k []int) bool {
	for i := range k {
		if k[i] != mas[i] {
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
	if count == len(mas) {
		str := fmt.Sprint(res) + "\n"
		f.WriteString(str)
		//	fmt.Println(res, b)
	}
}

/**
Where r[i] - выходная последовательность
*/
func IA(mas []int, q int) {
	r := make([]int, 2*size())
	var x int
	masRes := make([]int, len(mas))
	copy(masRes, mas)
	b := q
	var i int
	for i = 0; i < int(count); i += 1 {
		x = mas[mod(i)]
		masRes[mod(i)] = modK(masRes[mod(x)] + b)
		r[i] = modK(masRes[mod(masRes[mod(i)]>>n)] + x)
		b = r[i]
	}

	// Проверка на неизменения состояния
	masRes2 := make([]int, len(mas))
	copy(masRes2, mas)
	masRes3 := make([]int, len(mas))
	copy(masRes3, masRes)


	// проверка на сплошные нули и равенство уравнений
	if checkTwo(masRes2, masRes3) {
		str := "S: " + fmt.Sprint(mas) + "  	 q:" + fmt.Sprint(r)+"\n"
		f.WriteString(str)
		fmt.Println(str)
		//checkZeroAllNumber(r, mas)
	}
}

func genAll(mas2 []int, pos int) {
	//fmt.Println(K())
    //fmt.Println("N and D",n,d)
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

	var nameFile string

	fmt.Println("Введите N")
	fmt.Scan(&n)
    fmt.Println("Введите D")
    fmt.Scan(&d)
    r := make([]int, 1<<n)
    fmt.Println(" Введите сколько элементов выходной последовательности выводить")
    fmt.Scan(&count)
	nameFile = fmt.Sprintf("Ключи N=%d D=%d.txt", n, d)
	f, _ = os.Create(nameFile)
	defer f.Close()
	genAll(r, 0)
	fmt.Println("Ключи сохранились в файле: ",nameFile)
	fmt.Println("Нажмите клавишу для выхода")
	fmt.Scan(&count)
	//genRandomNumber(r2)
}
