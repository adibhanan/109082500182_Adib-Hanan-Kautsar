package main

import "fmt"

const nProv int = 10

type namaProv [nProv]string
type popProv [nProv]int
type tumbuhProv [nProv]float64

func inputData(prov *namaProv, pop *popProv, tumbuh *tumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func provinsiTercepat(tumbuh tumbuhProv) int {
	idx := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idx] {
			idx = i
		}
	}
	return idx
}

func indeksProvinsi(prov namaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func prediksi(prov namaProv, pop popProv, tumbuh tumbuhProv) {
	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := int(float64(pop[i]) * (1 + tumbuh[i]))
			fmt.Println(prov[i], prediksi)
		}
	}
}

func main() {
	var prov namaProv
	var pop popProv
	var tumbuh tumbuhProv
	var cari string

	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	inputData(&prov, &pop, &tumbuh)

	fmt.Print("\nMasukkan nama provinsi yang dicari : ")
	fmt.Scan(&cari)

	idx := provinsiTercepat(tumbuh)
	fmt.Println("\nProvinsi dengan angka pertumbuhan tercepat :", prov[idx])

	idxCari := indeksProvinsi(prov, cari)
	fmt.Println("\nData provinsi yang dicari :", cari)

	if idxCari != -1 {
		fmt.Println()
		prediksi(prov, pop, tumbuh)
	}
}
