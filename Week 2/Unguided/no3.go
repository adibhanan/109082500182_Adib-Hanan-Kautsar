package main

import "fmt"

func main() {
	var berat int64
	var kg, sisaGram int64
	var biayaKg, biayaSisa, totalBiaya int64
	fmt.Print("Masukkan total berat (gram) : ")

	fmt.Scan(&berat)

	kg = berat / 1000
	biayaKg = kg * 10000
	sisaGram = berat - (kg * 1000)

	if sisaGram >= 500 {
		biayaSisa = sisaGram * 5
	} else {
		biayaSisa = sisaGram * 15
	}

	if kg > 10 {
		totalBiaya = biayaKg
	} else {
		totalBiaya = biayaKg + biayaSisa
	}
	fmt.Println("==== Detail Perhitungan ====")
	fmt.Printf("Detail berat : %d kg + %d gr \n", kg, sisaGram)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d \n", biayaKg, biayaSisa)
	fmt.Print("Total biaya : Rp.", totalBiaya)
}
