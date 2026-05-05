package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   int
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputData(T *arrayMahasiswa, N int) {
	for i := 0; i < N; i++ {
		fmt.Print("Masukkan data ke-", i+1, ": ")
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}
}

func cariNilaiPertama(T arrayMahasiswa, N int, nim int) int {
	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			return T[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(T arrayMahasiswa, N int, nim int) int {
	max := -1

	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			if T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}

	return max
}

func main() {
	var data arrayMahasiswa
	var N, nim int

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&N)

	inputData(&data, N)

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nim)

	nilaiPertama := cariNilaiPertama(data, N, nim)
	nilaiTerbesar := cariNilaiTerbesar(data, N, nim)

	if nilaiPertama == -1 {
		fmt.Println("Data mahasiswa tidak ditemukan")
	} else {
		fmt.Println("Nilai pertama dari NIM", nim, "adalah", nilaiPertama)
		fmt.Println("Nilai terbesar dari NIM", nim, "adalah", nilaiTerbesar)
	}
}
