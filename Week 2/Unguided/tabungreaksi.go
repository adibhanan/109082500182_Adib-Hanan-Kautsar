package main

import "fmt"

func main() {
	var urutan, input1, input2, input3, input4 string
	var hasil bool
	hasil = true

	urutan = "merah kuning hijau ungu"

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scanln(&input1, &input2, &input3, &input4)
		input := input1 + " " + input2 + " " + input3 + " " + input4
		if urutan != input {
			hasil = false
		}
	}
	fmt.Print("Berhasil : ", hasil)
}
