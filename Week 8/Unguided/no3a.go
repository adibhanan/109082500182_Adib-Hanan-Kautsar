package main

import "fmt"

const MAX = 100

type ArrString [MAX]string

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang ArrString
	var n int = 0
	var i int = 1

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[n] = klubA
		} else if skorB > skorA {
			pemenang[n] = klubB
		} else {
			pemenang[n] = "Draw"
		}

		n++
		i++
	}

	for j := 0; j < n; j++ {
		fmt.Printf("Hasil %d: %s\n", j+1, pemenang[j])
	}

	fmt.Println("Pertandingan selesai")
}
